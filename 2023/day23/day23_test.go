package day23

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/bits-and-blooms/bitset"
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

type Trail2 struct {
	*Point
	steps   int
	visited bitset.BitSet
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
			branches := make([]*Trail, 0, 3)

			for {
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
					branches = branches[:0]
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
		func(t *Trail, _ int) int {
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

	result := common.IterativeSearch[Trail2, int, int](
		&Trail2{start, 0, common.Deref(bitset.New(uint(len(grid))))},
		func(t *Trail2) []*Trail2 {
			branches := make([]*Trail2, 0, 3)
			t.visited = common.Deref(t.visited.Clone())

			for {
				t.visited.Set(uint(t.y<<8 | t.x))

				upAllowed := t.y-1 >= 0 && !t.visited.Test(uint((t.y-1)<<8|t.x)) && grid[t.y-1][t.x].text != "#"
				downAllowed := t.y+1 < len(grid) && !t.visited.Test(uint((t.y+1)<<8|t.x)) && grid[t.y+1][t.x].text != "#"
				leftAllowed := t.x-1 >= 0 && !t.visited.Test(uint(t.y<<8|t.x-1)) && grid[t.y][t.x-1].text != "#"
				rightAllowed := t.x+1 < len(grid[0]) && !t.visited.Test(uint(t.y<<8|t.x+1)) && grid[t.y][t.x+1].text != "#"

				if upAllowed {
					branches = append(branches, &Trail2{&grid[t.y-1][t.x], t.steps + 1, t.visited})
				}
				if downAllowed {
					branches = append(branches, &Trail2{&grid[t.y+1][t.x], t.steps + 1, t.visited})
				}
				if leftAllowed {
					branches = append(branches, &Trail2{&grid[t.y][t.x-1], t.steps + 1, t.visited})
				}
				if rightAllowed {
					branches = append(branches, &Trail2{&grid[t.y][t.x+1], t.steps + 1, t.visited})
				}

				if len(branches) == 1 && branches[0].Point != end {
					t = branches[0]
					branches = branches[:0]
					continue
				}

				return branches
			}
		},
		func(t *Trail2) bool {
			return t.Point == end
		},
		nil,
		func(t *Trail2, cw int) int {
			return t.steps
		},
		nil,
		0,
		false,
		true,
	)

	return result.BestWeight
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 154, result)
	} else {
		assert.Equal(t, 6502, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
