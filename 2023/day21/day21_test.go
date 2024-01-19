package day21

import (
	_ "embed"
	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
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
	visited := make(bitmap.Bitmap, 0, nextId/64+1)
	queue := make([]*Tile, 0, 2000)
	queue = append(queue, start)
	nextQueue := make([]*Tile, 0, 300) // For each step, tracks the new tiles found.

	for steps := 0; steps <= 64; steps++ {
		for len(queue) > 0 {
			cur := queue[len(queue)-1]
			queue = queue[: len(queue)-1 : cap(queue)]

			if visited.Contains(cur.id) {
				continue
			}

			if steps%2 == 0 {
				visited.Set(cur.id)
			}

			nextQueue = append(nextQueue, findNeighbors(cur.y, cur.x, visited)...)
		}

		queue, nextQueue = nextQueue, queue
	}

	sum := visited.Count()
	return sum
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

func part2() int {
	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 16733044, result)
	} else {
		assert.Equal(t, 357134560737, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		//part2()
	}
}
