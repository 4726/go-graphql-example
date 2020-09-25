package main

import (
	"github.com/graphql-go/graphql"
)

var playerObject *graphql.Object

func init() {
	playerObjConfig := ObjectConfig{
		Name: "Player",
		Fields: Player,
	}
	playerObject = graphql.NewObject(playerObjConfig)
}

func schemas() graphql.Fields {
	return graphql.Fields{
		"kClosest": kClosestQuery(),
		"spawn": spawnMutation(),
		"leave": leaveMutation(),
		"allPlayers": allQuery(),
	}
}

func kClosestQuery() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["k"] = &graphql.ArgumentConfig{
		Type: graphql.NonNull(graphql.Int),
	}
	args["x"] = &graphql.ArgumentConfig{
		Type: graphql.Float,
		DefaultValue: 0,
		Description: "X-Coordinate",
	}
	args["y"] = &graphql.ArgumentConfig{
		Type: graphql.Float,
		DefaultValue: 0,
		Description: "Y-Coordinate",
	}
	
	return &graphql.Field{
		Type: playerObject,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return kClosest(p.Args["x"], p.Args["y"], p.Args["k"]), nil
		}
	}
}

func spawnMutation() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["name"] = &graphql.ArgumentConfig{
		Type: graphql.NonNull(graphql.String),
	}
	args["x"] = &graphql.ArgumentConfig{
		Type: graphql.Float,
		DefaultValue: 0,
		Description: "X-Coordinate",
	}
	args["y"] = &graphql.ArgumentConfig{
		Type: graphql.Float,
		DefaultValue: 0,
		Description: "Y-Coordinate",
	}
	
	return &graphql.Field{
		Type: graphql.ID,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := spawn(p.Args["name"], p.Args["x"], p.Args["y"])
			return graphql.ID.ParseLiteral(id), nil
		}
	}
}

func leaveMutation() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["id"] = &graphql.ArgumentConfig{
		Type: graphql.NonNull(graphql.ID),
	}
	
	return &graphql.Field{
		Type: graphql.ID,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := spawn(graphql.ID.ParseValue(p.Args["id"]))
			return p.Args["id"], nil
		}
	}
}

func allPlayersQuery() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	
	return &graphql.Field{
		Type: graphql.NewList(playerObject),
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return allPlayers(), nil
		}
	}
}

func moveMutation() *graphql.Field {
	args := map[string]*graphql.ArgumentConfig{}
	args["id"] = &graphql.ArgumentConfig{
		Type: graphql.NonNull(graphql.ID),
	}
	args["x"] = &graphql.ArgumentConfig{
		Type: graphql.NonNull(graphql.Float),
		Description: "X-Coordinate",
	}
	args["y"] = &graphql.ArgumentConfig{
		Type: graphql.NonNull(graphql.Float),
		Description: "Y-Coordinate",
	}
	
	return &graphql.Field{
		Type: graphql.ID,
		Args: args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			move(graphql.ID.ParseValue(p.Args["id"]), p.Args["x"], p.Args["y"])
			return p.Args["id"], nil
		}
	}
}