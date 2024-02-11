package day21

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Tile struct {
	id   uint32
	y, x int
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := make([][]*Tile, len(lines))

	// Read grid.
	var start *Tile
	nextId := uint32(1)
	for y, line := range lines {
		grid[y] = make([]*Tile, len(line))
		for x, char := range line {
			if char == '#' {
				continue
			}

			grid[y][x] = &Tile{nextId, y, x}
			nextId++

			if char == 'S' {
				start = grid[y][x]
			}
		}
	}

	findNeighbors := func(y, x int, visited bitmap.Bitmap) []*Tile {
		neighbors := make([]*Tile, 0, 4)
		if y > 0 {
			n := grid[y-1][x]
			if n != nil && !visited.Contains(n.id) {
				neighbors = append(neighbors, n)
			}
		}
		if y < len(grid)-1 {
			n := grid[y+1][x]
			if n != nil && !visited.Contains(n.id) {
				neighbors = append(neighbors, n)
			}
		}
		if x > 0 {
			n := grid[y][x-1]
			if n != nil && !visited.Contains(n.id) {
				neighbors = append(neighbors, n)
			}
		}
		if x < len(grid[0])-1 {
			n := grid[y][x+1]
			if n != nil && !visited.Contains(n.id) {
				neighbors = append(neighbors, n)
			}
		}
		return neighbors
	}

	// Find visited tiles per even number of steps.
	queue := make([]*Tile, 0, 1000)
	nextQueue := make([]*Tile, 0, 1000)
	visited := make(bitmap.Bitmap, 0, nextId/64+1)
	numVisited := 0

	queue = append(queue, start)
	for steps := 0; steps <= 64; steps++ {
		for _, cur := range queue {
			if visited.Contains(cur.id) {
				continue
			}

			if steps%2 == 0 {
				visited.Set(cur.id)
				numVisited++
			}

			nextQueue = append(nextQueue, findNeighbors(cur.y, cur.x, visited)...)
		}

		queue, nextQueue = nextQueue, queue[:0:cap(queue)]
	}

	return numVisited
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 42, result)
	} else {
		assert.Equal(t, 3594, result)
	}
}

// Part 2.

func toID(y, x int) uint32 {
	return uint32(y+math.MaxInt16)<<16 | uint32(x+math.MaxInt16)
}

func fromID(id uint32) (int, int) {
	y := int(id>>16) - math.MaxInt16
	x := int(id&0xffff) - math.MaxInt16
	return y, x
}

func TestToID(t *testing.T) {
	t.Skip()
	t.Parallel()

	assert.Equal(t, uint32(0x7fff7fff), toID(0, 0))
	assert.Equal(t, uint32(0x7fff8000), toID(0, 1))
	assert.Equal(t, uint32(0x80007fff), toID(1, 0))
	assert.Equal(t, uint32(0x7fff7ffe), toID(0, -1))
	assert.Equal(t, uint32(0x7ffe7fff), toID(-1, 0))
	assert.Equal(t, uint32(0x7ffe7ffe), toID(-1, -1))
	y, x := fromID(0x7ffe7ffe)
	assert.Equal(t, -1, y)
	assert.Equal(t, -1, x)
	y, x = fromID(0x80008000)
	assert.Equal(t, 1, y)
	assert.Equal(t, 1, x)
}

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	height := len(lines)
	width := len(lines[0])
	start := width / 2

	isWalkable := func(y, x int) bool {
		return lines[common.EuclideanMod(y, height)][common.EuclideanMod(x, width)] != '#'
	}

	findNeighbors := func(ID uint32, visited bitmap.Bitmap) []uint32 {
		neighbors := make([]uint32, 0, 4)
		y, x := fromID(ID)

		if isWalkable(y-1, x) && !visited.Contains(toID(y-1, x)) {
			neighbors = append(neighbors, toID(y-1, x))
		}
		if isWalkable(y+1, x) && !visited.Contains(toID(y+1, x)) {
			neighbors = append(neighbors, toID(y+1, x))
		}
		if isWalkable(y, x-1) && !visited.Contains(toID(y, x-1)) {
			neighbors = append(neighbors, toID(y, x-1))
		}
		if isWalkable(y, x+1) && !visited.Contains(toID(y, x+1)) {
			neighbors = append(neighbors, toID(y, x+1))
		}

		return neighbors
	}

	queues := [2][]uint32{make([]uint32, 0, 5400), make([]uint32, 0, 5400)}
	counts := [2]int{}
	visited := make(bitmap.Bitmap, 0, 34000000)
	goalSteps := 26501365
	yForX := [3]int{}

	queues[0] = append(queues[0], toID(start, start))
	for steps := 0; steps <= 3*width; steps++ {
		parity := steps % 2
		nextParity := (steps + 1) % 2
		queues[nextParity] = queues[nextParity][:0:cap(queues[nextParity])]

		for _, curID := range queues[parity] {
			if !visited.Contains(curID) {
				visited.Set(curID)
				counts[parity]++
				queues[nextParity] = append(queues[nextParity], findNeighbors(curID, visited)...)
			}
		}

		isPointOnCurve := (steps-start)%width == 0
		if isPointOnCurve {
			yForX[steps/width] = counts[parity]
			if steps/width >= 2 {
				break
			}
		}
	}

	a, b, c := common.FindQuadraticCoefficients(yForX[0], yForX[1], yForX[2])
	x := goalSteps / width
	return common.SolveQuadratic(a, b, c, x)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 528192899606863, result)
	} else {
		assert.Equal(t, 605247138198755, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//part1()
		part2()
	}
}
