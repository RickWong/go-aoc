package day23

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
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

type Point struct {
	x, y int
	text string
}

type Trail struct {
	*Point
	steps int
	last  *Trail
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := make([][]Point, len(lines))
	for y, line := range lines {
		grid[y] = make([]Point, len(line))
		for x, text := range strings.Split(line, "") {
			grid[y][x] = Point{x, y, text}
		}
	}

	start := &grid[0][strings.Index(lines[0], ".")]
	end := &grid[len(grid)-1][strings.Index(lines[len(lines)-1], ".")]

	result := common.IterativeSearch(
		&Trail{start, 0, nil},
		func(t *Trail) []*Trail {
			for {
				branches := make([]*Trail, 0, 3)

				leftAllowed := (t.text == "<" || t.text == ".") && (t.last == nil || t.last.x != t.x-1)
				upAllowed := (t.text == "^" || t.text == ".") && (t.last == nil || t.last.y != t.y-1)
				downAllowed := (t.text == "v" || t.text == ".") && (t.last == nil || t.last.y != t.y+1)
				rightAllowed := (t.text == ">" || t.text == ".") && (t.last == nil || t.last.x != t.x+1)

				if upAllowed && t.y > 0 && grid[t.y-1][t.x].text != "#" {
					branches = append(branches, &Trail{&grid[t.y-1][t.x], t.steps + 1, t})
				}
				if downAllowed && t.y < len(grid)-1 && grid[t.y+1][t.x].text != "#" {
					branches = append(branches, &Trail{&grid[t.y+1][t.x], t.steps + 1, t})
				}
				if leftAllowed && t.x > 0 && grid[t.y][t.x-1].text != "#" {
					branches = append(branches, &Trail{&grid[t.y][t.x-1], t.steps + 1, t})
				}
				if rightAllowed && t.x < len(grid[0])-1 && grid[t.y][t.x+1].text != "#" {
					branches = append(branches, &Trail{&grid[t.y][t.x+1], t.steps + 1, t})
				}

				if len(branches) == 1 && branches[0].Point != end {
					t = branches[0]
					continue
				}

				return branches
			}
		},
		func(t *Trail) bool {
			return t.Point == end
		},
		func(t *Trail) uint32 {
			return uint32(t.y<<16 | t.x)
		},
		func(t *Trail, currentWeight int) int {
			return t.steps
		},
		nil,
		0,
		false,
		true,
	)

	return result.BestWeight
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 94, result)
	} else {
		assert.Equal(t, 2170, result)
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
