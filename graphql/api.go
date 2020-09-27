package graphql

import (
	"log"

	"github.com/graphql-go/graphql"
)

var playerObject *graphql.Object
var schema graphql.Schema

func Do(query string) *graphql.Result {
	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	return graphql.Do(params)
}

func init() {
	playerObjConfig := graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"x": &graphql.Field{
				Type: graphql.Float,
			},
			"y": &graphql.Field{
				Type: graphql.Float,
			},
			"distance": &graphql.Field{
				Type: graphql.Float,
			},
			"lastUpdate": &graphql.Field{
				Type: graphql.Int,
			},
		},
	}
	playerObject = graphql.NewObject(playerObjConfig)
	var err error
	schema, err = initSchema()
	if err != nil {
		log.Fatal("could not setup graphql schema: ", err)
	}
}

func initSchema() (graphql.Schema, error) {
	queryObjConf := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"kClosest":   kClosestQuery(),
			"allPlayers": allPlayersQuery(),
		},
	}
	query := graphql.NewObject(queryObjConf)

	mutationObjConf := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"spawn": spawnMutation(),
			"leave": leaveMutation(),
			"move":  moveMutation(),
		},
	}
	mutation := graphql.NewObject(mutationObjConf)

	schemaConf := graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	}
	return graphql.NewSchema(schemaConf)
}

func kClosestQuery() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["k"] = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}
	args["x"] = &graphql.ArgumentConfig{
		Type:         graphql.Float,
		DefaultValue: 0.0,
		Description:  "X-Coordinate",
	}
	args["y"] = &graphql.ArgumentConfig{
		Type:         graphql.Float,
		DefaultValue: 0.0,
		Description:  "Y-Coordinate",
	}

	return &graphql.Field{
		Type:        graphql.NewList(playerObject),
		Description: "Get k closest players from point (x, y)",
		Args:        args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			players := kClosest(p.Args["x"].(float64), p.Args["y"].(float64), p.Args["k"].(int))
			return players, nil
		},
	}
}

func spawnMutation() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["name"] = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	}
	args["x"] = &graphql.ArgumentConfig{
		Type:         graphql.Float,
		DefaultValue: 0,
		Description:  "X-Coordinate",
	}
	args["y"] = &graphql.ArgumentConfig{
		Type:         graphql.Float,
		DefaultValue: 0,
		Description:  "Y-Coordinate",
	}

	return &graphql.Field{
		Type: playerObject,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			player := spawn(p.Args["name"].(string), p.Args["x"].(float64), p.Args["y"].(float64))
			return player, nil
		},
	}
}

func leaveMutation() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["id"] = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.ID),
	}

	return &graphql.Field{
		Type: graphql.ID,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			leave(p.Args["id"].(string))
			return p.Args["id"], nil
		},
	}
}

func allPlayersQuery() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}

	return &graphql.Field{
		Type: graphql.NewList(playerObject),
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return all(), nil
		},
	}
}

func moveMutation() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["id"] = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.ID),
	}
	args["x"] = &graphql.ArgumentConfig{
		Type:        graphql.NewNonNull(graphql.Float),
		Description: "X-Coordinate",
	}
	args["y"] = &graphql.ArgumentConfig{
		Type:        graphql.NewNonNull(graphql.Float),
		Description: "Y-Coordinate",
	}

	return &graphql.Field{
		Type: playerObject,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			player := move(p.Args["id"].(string), p.Args["x"].(float64), p.Args["y"].(float64))
			return player, nil
		},
	}
}
