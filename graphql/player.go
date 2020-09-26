package graphql

import (
	"container/heap"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	Name, ID       string
	X, Y, Distance float64
	LastUpdate     int64
}

var allPlayers = make(map[string]Player)
var mux sync.RWMutex

//kClosest returns the k closest players from the point (x, y)
func kClosest(x, y float64, k int) []Player {
	mux.RLock()
	defer mux.RUnlock()

	maxHeap := &maxHeap{}
	for _, v := range allPlayers {
		heap.Push(maxHeap, idDist{v.ID, distance(x, y, v.X, v.Y)})
		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}

	var players []Player

	for maxHeap.Len() > 0 {
		pop := heap.Pop(maxHeap).(idDist)
		p := allPlayers[pop.ID]
		p.Distance = pop.Dist
		players = append(players, p)
	}

	for i := 0; i < len(players)/2; i++ {
		players[i], players[len(players)-i-1] = players[len(players)-i-1], players[i]
	}

	return players
}

func spawn(name string, x, y float64) Player {
	mux.Lock()
	defer mux.Unlock()

	var id string
	for {
		tryID := randID(16)
		if _, ok := allPlayers[tryID]; !ok {
			id = tryID
			break
		}
	}
	allPlayers[id] = Player{
		Name:       name,
		ID:         id,
		X:          x,
		Y:          y,
		LastUpdate: time.Now().Unix(),
	}

	return allPlayers[id]
}

func leave(id string) {
	mux.Lock()
	defer mux.Unlock()

	delete(allPlayers, id)
}

func all() []Player {
	mux.RLock()
	defer mux.RUnlock()

	var res []Player
	for _, v := range allPlayers {
		res = append(res, v)
	}
	return res
}

func move(id string, x, y float64) Player {
	mux.Lock()
	defer mux.Unlock()

	if p, ok := allPlayers[id]; ok {
		p.X = x
		p.Y = y
		p.LastUpdate = time.Now().Unix()
		allPlayers[id] = p
	}
	return allPlayers[id]
}

func distance(x, y, x2, y2 float64) float64 {
	xDist := math.Abs(x2 - x)
	yDist := math.Abs(y2 - y)
	return xDist + yDist
}

func clear() {
	mux.Lock()
	defer mux.Unlock()

	allPlayers = map[string]Player{}
}

func randID(size int) string {
	res := make([]rune, size)

	for i := range res {
		r := rand.Intn(26 + 26 + 10)
		if r < 26 {
			res[i] = rune('a' + r)
		} else if r >= 26 && r < 52 {
			res[i] = rune('A' - 26 + r)
		} else {
			res[i] = rune('0' - 52 + r)
		}
	}

	return string(res)
}
