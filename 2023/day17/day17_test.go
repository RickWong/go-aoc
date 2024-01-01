package day17

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
	"github.com/samber/lo"
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

type Block struct {
	y, x int
	loss int
}

type Trail struct {
	y, x       int
	N, E, S, W int // Number of steps in CURRENT direction. RESETS at corners.
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := lo.Map(lines, func(line string, y int) []Block {
		return lo.Map(strings.Split(line, ""), func(c string, x int) Block {
			return Block{y, x, int(c[0] - '0')}
		})
	})

	start := Trail{}
	end := &grid[len(grid)-1][len(grid[0])-1]
	result := common.IterativeSearch(
		&start,
		func(cur *Trail) []*Trail {
			branches := make([]*Trail, 0, 4)

			if cur.y > start.y && cur.N < 3 && cur.S == 0 {
				branches = append(branches,
					&Trail{cur.y - 1, cur.x, cur.N + 1, 0, 0, 0})
			}
			if cur.y < end.y && cur.S < 3 && cur.N == 0 {
				branches = append(branches,
					&Trail{cur.y + 1, cur.x, 0, 0, cur.S + 1, 0})
			}
			if cur.x > start.x && cur.W < 3 && cur.E == 0 {
				branches = append(branches,
					&Trail{cur.y, cur.x - 1, 0, 0, 0, cur.W + 1})
			}
			if cur.x < end.x && cur.E < 3 && cur.W == 0 {
				branches = append(branches,
					&Trail{cur.y, cur.x + 1, 0, cur.E + 1, 0, 0})
			}
			return branches
		},
		func(cur *Trail) bool {
			return cur.y == end.y && cur.x == end.x
		},
		func(cur *Trail) any {
			// Encode trail state as a single 32 bit integer.
			return cur.y<<0 | cur.x<<10 | cur.N<<20 | cur.E<<23 | cur.S<<26 | cur.W<<29
		},
		func(cur *Trail) float64 {
			return float64(grid[cur.y][cur.x].loss)
		},
		func(cur *Trail) float64 {
			return float64(end.y - cur.y + end.x - cur.x)
		},
		3,
		false,
		false,
	)

	return int(result.BestWeight)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 102, result, "Result was incorrect")
	} else {
		assert.Equal(t, 1128, result, "Result was incorrect")
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := lo.Map(lines, func(line string, y int) []Block {
		return lo.Map(strings.Split(line, ""), func(c string, x int) Block {
			return Block{y, x, int(c[0] - '0')}
		})
	})

	start := Trail{}
	end := &grid[len(grid)-1][len(grid[0])-1]
	result := common.IterativeSearch(
		&start,
		func(cur *Trail) []*Trail {
			branches := make([]*Trail, 0, 4)

			northAllowed := cur.y > start.y && cur.N < 10 && cur.S == 0 &&
				((cur.S|cur.W|cur.E) == 0 || (cur.S|cur.W|cur.E) >= 4)
			southAllowed := cur.y < end.y && cur.S < 10 && cur.N == 0 &&
				((cur.N|cur.W|cur.E) == 0 || (cur.N|cur.W|cur.E) >= 4)
			westAllowed := cur.x > start.x && cur.W < 10 && cur.E == 0 &&
				((cur.N|cur.S|cur.E) == 0 || (cur.N|cur.S|cur.E) >= 4)
			eastAllowed := cur.x < end.x && cur.E < 10 && cur.W == 0 &&
				((cur.N|cur.S|cur.W) == 0 || (cur.N|cur.S|cur.W) >= 4)

			if northAllowed {
				branches = append(branches,
					&Trail{cur.y - 1, cur.x, cur.N + 1, 0, 0, 0})
			}
			if southAllowed {
				branches = append(branches,
					&Trail{cur.y + 1, cur.x, 0, 0, cur.S + 1, 0})
			}
			if westAllowed {
				branches = append(branches,
					&Trail{cur.y, cur.x - 1, 0, 0, 0, cur.W + 1})
			}
			if eastAllowed {
				branches = append(branches,
					&Trail{cur.y, cur.x + 1, 0, cur.E + 1, 0, 0})
			}
			return branches
		},
		func(cur *Trail) bool {
			return cur.y == end.y && cur.x == end.x && (cur.N|cur.S|cur.W|cur.E) >= 4
		},
		func(cur *Trail) any {
			// Encode trail state as a single 64 bit integer.
			ret := int64(cur.y)
			ret |= int64(cur.x) << 10
			ret |= int64(cur.N) << 20
			ret |= int64(cur.E) << 24
			ret |= int64(cur.S) << 28
			ret |= int64(cur.W) << 32
			return ret
		},
		func(cur *Trail) float64 {
			return float64(grid[cur.y][cur.x].loss)
		},
		func(cur *Trail) float64 {
			return float64(end.y - cur.y + end.x - cur.x)
		},
		0,
		false,
		false,
	)

	return int(result.BestWeight)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 94, result, "Result was incorrect")
	} else {
		assert.Equal(t, 1268, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
