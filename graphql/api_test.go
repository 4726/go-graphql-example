package graphql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mapToPlayer(m map[string]interface{}) Player {
	var p Player
	if v, ok := m["name"]; ok {
		p.Name = v.(string)
	}
	if v, ok := m["id"]; ok {
		p.ID = v.(string)
	}
	if v, ok := m["x"]; ok {
		p.X = v.(float64)
	}
	if v, ok := m["y"]; ok {
		p.Y = v.(float64)
	}
	if v, ok := m["distance"]; ok {
		p.Distance = v.(float64)
	}
	if v, ok := m["lastUpdate"]; ok {
		p.LastUpdate = int64(v.(int))
	}
	return p
}

func TestQueryKClosestRequired(t *testing.T) {
	res := Do(`query {
		kClosest(x: 0, y: 0) {
			name
			distance
		}
	}`)

	assert.Len(t, res.Errors, 1)
}

func TestQueryKClosestNone(t *testing.T) {
	res := Do(`query {
		kClosest(k: 5) {
			name
			distance
		}
	}`)

	assert.Len(t, res.Data.(map[string]interface{})["kClosest"], 0)
	assert.Len(t, res.Errors, 0)
}

func TestQueryKClosest(t *testing.T) {
	clear()
	p1 := spawn("a", 0, 0)
	p2 := spawn("a", 1, 1)
	p3 := spawn("a", 2, 2)
	p4 := spawn("a", 3, 3)
	p5 := spawn("a", 4, 4)
	spawn("a", 5, 5)
	spawn("a", 6, 6)
	spawn("a", 7, 7)
	spawn("a", 8, 8)
	spawn("a", 9, 9)

	res := Do(`query {
		kClosest(k: 5) {
			name
			distance
			id
			lastUpdate
			x
			y
		}
	}`)
	players := res.Data.(map[string]interface{})["kClosest"].([]interface{})
	p1.Distance = 0
	p2.Distance = 2
	p3.Distance = 4
	p4.Distance = 6
	p5.Distance = 8
	assert.Equal(t, p1, mapToPlayer(players[0].(map[string]interface{})))
	assert.Equal(t, p2, mapToPlayer(players[1].(map[string]interface{})))
	assert.Equal(t, p3, mapToPlayer(players[2].(map[string]interface{})))
	assert.Equal(t, p4, mapToPlayer(players[3].(map[string]interface{})))
	assert.Equal(t, p5, mapToPlayer(players[4].(map[string]interface{})))
	assert.Len(t, players, 5)
	assert.Len(t, res.Errors, 0)
}

func TestQueryKClosest2(t *testing.T) {
	clear()
	spawn("a", 0, 0)
	spawn("a", 1, 1)
	spawn("a", 2, 2)
	spawn("a", 3, 3)
	p5 := spawn("a", 4, 4)
	p6 := spawn("a", 5, 5)
	p7 := spawn("a", 6, 6)
	p8 := spawn("a", 7, 7)
	p9 := spawn("a", 8, 8)
	p10 := spawn("a", 9, 9)

	res := Do(`query {
		kClosest(x: 100, y: 100, k: 6) {
			name
			distance
			id
			lastUpdate
			x
			y
		}
	}`)
	players := res.Data.(map[string]interface{})["kClosest"].([]interface{})
	p10.Distance = 182
	p9.Distance = 184
	p8.Distance = 186
	p7.Distance = 188
	p6.Distance = 190
	p5.Distance = 192
	assert.Equal(t, p10, mapToPlayer(players[0].(map[string]interface{})))
	assert.Equal(t, p9, mapToPlayer(players[1].(map[string]interface{})))
	assert.Equal(t, p8, mapToPlayer(players[2].(map[string]interface{})))
	assert.Equal(t, p7, mapToPlayer(players[3].(map[string]interface{})))
	assert.Equal(t, p6, mapToPlayer(players[4].(map[string]interface{})))
	assert.Equal(t, p5, mapToPlayer(players[5].(map[string]interface{})))
	assert.Len(t, players, 6)
	assert.Len(t, res.Errors, 0)
}

func TestMutationSpawnRequired(t *testing.T) {
	res := Do(`mutation {
		spawn(x: 100, y: 100) {
			id
		}
	}`)

	assert.Len(t, res.Errors, 1)
}

func TestMutationSpawn(t *testing.T) {
	clear()
	res := Do(`mutation {
		spawn(name: "a", x: 100, y: 101) {
			id
			name
			distance
			x
			y
			lastUpdate
		}
	}`)
	player := mapToPlayer(res.Data.(map[string]interface{})["spawn"].(map[string]interface{}))
	assert.Equal(t, 100.0, player.X)
	assert.Equal(t, 101.0, player.Y)
	assert.NotEmpty(t, player.ID)
	assert.Equal(t, "a", player.Name)
	assert.Equal(t, all()[0], player)
	assert.Len(t, all(), 1)
	assert.Len(t, res.Errors, 0)
}

func TestMutationLeaveRequired(t *testing.T) {
	clear()
	res := Do(`mutation {
		leave() {
			id
		}	
	}`)
	assert.Len(t, res.Errors, 1)
}

func TestMutationLeave(t *testing.T) {
	clear()

	p := spawn("a", 0, 0)
	res := Do(fmt.Sprintf(`mutation {
		leave(id: "%v")
	}`, p.ID))
	assert.Len(t, res.Errors, 0)
	assert.Len(t, all(), 0)
	assert.Equal(t, p.ID, res.Data.(map[string]interface{})["leave"])
}

func TestQueryAll(t *testing.T) {
	clear()

	res := Do(`query {
		allPlayers {
			name
			id
			lastUpdate
			x
			y
		}
	}`)
	players := res.Data.(map[string]interface{})["allPlayers"].([]interface{})

	assert.Len(t, res.Errors, 0)
	assert.Len(t, players, 0)

	p := spawn("a", 0, 0)

	res = Do(`query {
		allPlayers {
			name
			id
			lastUpdate
			x
			y
		}
	}`)

	players = res.Data.(map[string]interface{})["allPlayers"].([]interface{})

	assert.Len(t, res.Errors, 0)
	assert.Len(t, players, 1)
	assert.Equal(t, p, mapToPlayer(players[0].(map[string]interface{})))
}

func TestMutationmoveRequired(t *testing.T) {
	clear()

	res := Do(`mutation {
		move {
			id
		}
	}`)
	assert.Len(t, res.Errors, 3)
}

func TestMutationMove(t *testing.T) {
	clear()

	p := spawn("a", 0, 0)
	res := Do(fmt.Sprintf(`mutation {
		move(id: "%v", x: 100, y: 101) {
			id
			name
			x
			y
		}
	}`, p.ID))
	pRes := res.Data.(map[string]interface{})["move"].(map[string]interface{})
	assert.Len(t, res.Errors, 0)
	assert.Len(t, all(), 1)
	assert.Equal(t, p.ID, pRes["id"])
	assert.Equal(t, p.Name, pRes["name"])
	assert.Equal(t, 100.0, pRes["x"])
	assert.Equal(t, 101.0, pRes["y"])
}
