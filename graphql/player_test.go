package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpawn(t *testing.T) {
	clear()
	p := spawn("a", 0, 0)

	assert.Len(t, all(), 1)
	assert.Equal(t, "a", p.Name)
	assert.Equal(t, p, all()[0])
}

func TestLeave(t *testing.T) {
	clear()
	leave("1234567")
	assert.Len(t, all(), 0)

	p := spawn("1", 0, 0)
	leave(p.ID)
	assert.Len(t, all(), 0)
}

func TestMove(t *testing.T) {
	clear()
	p := spawn("a", 0, 0)

	p2 := move(p.ID, 1, 1)
	assert.Len(t, all(), 1)
	assert.Equal(t, p2, all()[0])
}

func TestDistance(t *testing.T) {
	clear()
	assert.Equal(t, float64(0), distance(0, 0, 0, 0))
	assert.Equal(t, float64(2), distance(0, 0, 1, 1))
	assert.Equal(t, float64(2), distance(1, 1, 0, 0))
	assert.Equal(t, float64(692), distance(0, 100, 265, 527))
	assert.Equal(t, float64(1521), distance(-100, 310, 582, -529))
}

func TestKClosest(t *testing.T) {
	clear()
	p1 := spawn("a", 0, 0)
	p2 := spawn("a", 1, 1)
	p3 := spawn("a", 2, 2)
	p4 := spawn("a", 3, 3)
	p5 := spawn("a", 4, 4)
	p6 := spawn("a", 5, 5)
	p7 := spawn("a", 6, 6)
	p8 := spawn("a", 7, 7)
	p9 := spawn("a", 8, 8)
	p10 := spawn("a", 9, 9)

	players := kClosest(0, 0, 5)
	p1.Distance = 0
	p2.Distance = 2
	p3.Distance = 4
	p4.Distance = 6
	p5.Distance = 8
	assert.Equal(t, p1, players[0])
	assert.Equal(t, p2, players[1])
	assert.Equal(t, p3, players[2])
	assert.Equal(t, p4, players[3])
	assert.Equal(t, p5, players[4])
	assert.Len(t, players, 5)

	p10.Distance = 182
	p9.Distance = 184
	p8.Distance = 186
	p7.Distance = 188
	p6.Distance = 190
	p5.Distance = 192
	players = kClosest(100, 100, 6)
	assert.Equal(t, p10, players[0])
	assert.Equal(t, p9, players[1])
	assert.Equal(t, p8, players[2])
	assert.Equal(t, p7, players[3])
	assert.Equal(t, p6, players[4])
	assert.Equal(t, p5, players[5])
	assert.Len(t, players, 6)
}
