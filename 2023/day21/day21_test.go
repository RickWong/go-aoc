package day21

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
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
	id        int
	y, x      int
	hash      bool
	neighbors []*Tile
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := make([][]*Tile, len(lines))

	// Read grid.
	var start *Tile
	for y, line := range lines {
		grid[y] = make([]*Tile, len(line))
		for x, char := range line {
			grid[y][x] = &Tile{y<<16 | x, y, x, char == '#', nil}
			if char == 'S' {
				start = grid[y][x]
			}
		}
	}

	// Connect neighbors.
	// TODO: Use bitset
	visited := make(map[int]bool, len(grid)*len(grid[0]))
	queue := make([]*Tile, 0, 64)
	queue = append(queue, start)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visited[cur.id] {
			continue
		}

		visited[cur.id] = true

		if cur.y > 0 && !grid[cur.y-1][cur.x].hash {
			cur.neighbors = append(cur.neighbors, grid[cur.y-1][cur.x])
		}

		if cur.y < len(grid)-1 && !grid[cur.y+1][cur.x].hash {
			cur.neighbors = append(cur.neighbors, grid[cur.y+1][cur.x])
		}

		if cur.x > 0 && !grid[cur.y][cur.x-1].hash {
			cur.neighbors = append(cur.neighbors, grid[cur.y][cur.x-1])
		}

		if cur.x < len(grid[cur.y])-1 && !grid[cur.y][cur.x+1].hash {
			cur.neighbors = append(cur.neighbors, grid[cur.y][cur.x+1])
		}

		for _, neighbor := range cur.neighbors {
			if !visited[neighbor.id] {
				queue = append(queue, neighbor)
			}
		}
	}

	// Find visisted tiles per even number of steps.
	maps.Clear(visited)
	queue = append(queue, start)
	nextQueue := make([]*Tile, 0, 64)

	for steps := 0; steps <= 64; steps++ {
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			if !visited[cur.id] {
				nextQueue = append(nextQueue, cur.neighbors...)
			}

			if steps%2 == 0 {
				visited[cur.id] = true
			}
		}

		queue, nextQueue = nextQueue, queue
	}

	//Print(grid, visited)

	sum := len(visited)
	return sum
}

func Print(grid [][]*Tile, visited map[int]bool) {
	for _, row := range grid {
		for _, tile := range row {
			if tile.hash {
				print("#")
			} else if visited[tile.id] {
				print("O")
			} else {
				print(".")
			}
		}
		println()
	}
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
		assert.Equal(t, 82000210, result)
	} else {
		assert.Equal(t, 357134560737, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
