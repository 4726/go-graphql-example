package main

import (
	"container/heap"
	"math"
	"math/rand"
	"time"
)

type Player struct {
	Name, ID       string
	X, Y, Distance float64
	LastUpdate     int64
}

var allPlayers map[string]Player

//kClosest returns the k closest players from the point (x, y)
func kClosest(x, y, k int) []Player {
	maxHeap := &MaxHeap{}
	for _, v := range allPlayers {
		heap.Push(maxHeap, IDDist{v.ID, distance(x, y, v.X, v.Y)})
		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}

	var players []Player

	for maxHeap.Len() > 0 {
		pop := heap.Pop(maxHeap).(IDDist)
		p := allPlayers[pop.ID]
		p.Distance = pop.Dist
		players = append(players, p)
	}

	for i := 0; i < len(players)/2; i++ {
		players[i], players[len(players)-i-1] = players[len(players)-i-1], players[i]
	}

	return players
}

func spawn(name string, x, y float64) string {
	id := randID(16)
	allPlayers[id] = Player{
		Name: name,
		ID:   id,
		X:    x,
		Y:    y,
	}

	return id
}

func leave(id string) {
	delete(allPlayers, id)
}

func allPlayers() []Player {
	var res []Player
	for _, v := range allPlayers {
		res = append(res, v)
	}
	return res
}

func move(id string, x, y float64) {
	if _, ok := allPlayers[id]; ok {
		allPlayers[id].X = x
		allPlayers[id].Y = y
		allPlayers[id].LastUpdate = time.Now().Unix()
	}
}

func distance(x, y, x2, y2 float64) float64 {
	xDist := math.Abs(x2 - x)
	yDist := math.Abs(y2 - y)
	return xDist + yDist
}

func randID(size int) string {
	res := make([]rune, size)

	for i := range res {
		r := rand.Intn(26 + 26 + 10)
		if r < 26 {
			res[i] = 'a' + r
		} else if r >= 26 && r < 52 {
			res[i] = 'A' - 26 + r
		} else {
			res[i] = '0' - 52 + r
		}
	}

	return string(res)
}
