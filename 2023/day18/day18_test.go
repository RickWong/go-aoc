package day18

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Point struct {
	y, x int
}

type Hole struct {
	y, x  int
	color string
}

// Helper functions.

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	deltas := map[byte][2]int{
		'R': {0, 1},
		'L': {0, -1},
		'D': {1, 0},
		'U': {-1, 0},
	}

	y, x := 0, 0
	minY, minX := 0, 0
	maxY, maxX := 0, 0
	grid := make(map[int]map[int]*Hole)
	grid[y] = map[int]*Hole{x: {y, x, ""}}

	for _, line := range lines {
		parts := strings.Fields(line)
		direction, meters, color := parts[0][0], Atoi(parts[1]), parts[2][1:len(parts[2])-1]

		for i := 0; i < meters; i++ {
			y += deltas[direction][0]
			x += deltas[direction][1]
			minY, minX = min(minY, y), min(minX, x)
			maxY, maxX = max(maxY, y), max(maxX, x)

			if grid[y] == nil {
				grid[y] = make(map[int]*Hole)
			}
			grid[y][x] = &Hole{y, x, color}
		}
	}

	// Add top and bottom borders.
	grid[minY-1] = make(map[int]*Hole)
	grid[maxY+1] = make(map[int]*Hole)

	// Fill starting at the borders.
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			isBorder := y == minY-1 || y == maxY+1 || x == minX-1 || x == maxX+1

			if isBorder {
				grid[y][x] = &Hole{y, x, "outside"}

				// Use search algorithm to fill.
				common.IterativeSearch(
					&Point{y, x},
					func(cur *Point) []*Point {
						branches := make([]*Point, 0, 4)
						if cur.y > minY-1 && grid[cur.y-1][cur.x] == nil {
							grid[cur.y-1][cur.x] = &Hole{cur.y - 1, cur.x, "outside"}
							branches = append(branches, &Point{cur.y - 1, cur.x})
						}
						if cur.y < maxY+1 && grid[cur.y+1][cur.x] == nil {
							grid[cur.y+1][cur.x] = &Hole{cur.y + 1, cur.x, "outside"}
							branches = append(branches, &Point{cur.y + 1, cur.x})
						}
						if cur.x > minX-1 && grid[cur.y][cur.x-1] == nil {
							grid[cur.y][cur.x-1] = &Hole{cur.y, cur.x - 1, "outside"}
							branches = append(branches, &Point{cur.y, cur.x - 1})
						}
						if cur.x < maxX+1 && grid[cur.y][cur.x+1] == nil {
							grid[cur.y][cur.x+1] = &Hole{cur.y, cur.x + 1, "outside"}
							branches = append(branches, &Point{cur.y, cur.x + 1})
						}
						return branches
					},
					nil,
					func(cur *Point) any {
						return (cur.y << 32) | (cur.x & 0xffffffff)
					},
					nil,
					nil,
					0,
					false,
					false,
				)
			}
		}
	}

	sum := 0
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if grid[y][x] == nil || grid[y][x].color != "outside" {
				sum++
			}
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 62, result, "Result was incorrect")
	} else {
		// 28290 is too low.
		assert.Equal(t, 47527, result, "Result was incorrect")
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
		assert.Equal(t, 82000210, result, "Result was incorrect")
	} else {
		assert.Equal(t, 357134560737, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
