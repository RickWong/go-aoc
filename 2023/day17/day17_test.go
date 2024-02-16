package day17

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
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
	y, x      int
	cumLoss   int
	direction int
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := common.Map2(lines, func(line string, y int) []Block {
		return common.Map2([]byte(line), func(c byte, x int) Block {
			return Block{y, x, int(c - '0')}
		})
	})

	start := Trail{}
	end := &grid[len(grid)-1][len(grid[0])-1]
	result := common.BucketSearch(
		&start,
		func(cur *Trail) []*Trail {
			minSteps := 1
			maxSteps := 3
			branches := make([]*Trail, 0, maxSteps*2)
			cumLoss := 0

			if cur.direction == 0 || cur.direction%2 == 0 {
				cumLoss = cur.cumLoss
				for i := minSteps; i <= maxSteps; i++ {
					if cur.x-i >= 0 {
						cumLoss += grid[cur.y][cur.x-i].loss
						branches = append(branches,
							&Trail{cur.y, cur.x - i, cumLoss, 1})
					}
				}
				cumLoss = cur.cumLoss
				for i := minSteps; i <= maxSteps; i++ {
					if cur.x+i < len(grid[cur.y]) {
						cumLoss += grid[cur.y][cur.x+i].loss
						branches = append(branches,
							&Trail{cur.y, cur.x + i, cumLoss, 1})
					}
				}
			}

			if cur.direction == 0 || cur.direction%2 == 1 {
				cumLoss = cur.cumLoss
				for i := minSteps; i <= maxSteps; i++ {
					if cur.y-i >= 0 {
						cumLoss += grid[cur.y-i][cur.x].loss
						branches = append(branches,
							&Trail{cur.y - i, cur.x, cumLoss, 2})
					}
				}
				cumLoss = cur.cumLoss
				for i := minSteps; i <= maxSteps; i++ {
					if cur.y+i < len(grid) {
						cumLoss += grid[cur.y+i][cur.x].loss
						branches = append(branches,
							&Trail{cur.y + i, cur.x, cumLoss, 2})
					}
				}
			}

			return branches
		},
		func(cur *Trail) bool {
			return cur.y == end.y && cur.x == end.x
		},
		func(cur *Trail) int {
			return cur.y<<20 | cur.x<<10 | cur.direction
		},
		func(cur *Trail, _ int) int {
			return cur.cumLoss
		},
		func(cur *Trail) int {
			return common.Manhattan(cur.y, cur.x, end.y, end.x)
		},
		8,
		true,
		false,
	)

	return result.Best.cumLoss
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 102, result)
	} else {
		assert.Equal(t, 1128, result)
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := common.Map2(lines, func(line string, y int) []Block {
		return common.Map2([]byte(line), func(c byte, x int) Block {
			return Block{y, x, int(c - '0')}
		})
	})

	start := Trail{}
	end := &grid[len(grid)-1][len(grid[0])-1]
	minSteps := 4
	maxSteps := 10

	result := common.BucketSearch(
		&start,
		func(cur *Trail) []*Trail {
			branches := make([]*Trail, 0, maxSteps*2)
			cumLoss := 0

			if cur.direction == 0 || cur.direction%2 == 0 {
				cumLoss = cur.cumLoss
				for i := 1; i <= maxSteps; i++ {
					if cur.x-i >= 0 {
						cumLoss += grid[cur.y][cur.x-i].loss
						if i >= minSteps {
							branches = append(branches,
								&Trail{cur.y, cur.x - i, cumLoss, 1})
						}
					}
				}
				cumLoss = cur.cumLoss
				for i := 1; i <= maxSteps; i++ {
					if cur.x+i < len(grid[cur.y]) {
						cumLoss += grid[cur.y][cur.x+i].loss
						if i >= minSteps {
							branches = append(branches,
								&Trail{cur.y, cur.x + i, cumLoss, 1})
						}
					}
				}
			}

			if cur.direction == 0 || cur.direction%2 == 1 {
				cumLoss = cur.cumLoss
				for i := 1; i <= maxSteps; i++ {
					if cur.y-i >= 0 {
						cumLoss += grid[cur.y-i][cur.x].loss
						if i >= minSteps {
							branches = append(branches,
								&Trail{cur.y - i, cur.x, cumLoss, 2})
						}
					}
				}
				cumLoss = cur.cumLoss
				for i := 1; i <= maxSteps; i++ {
					if cur.y+i < len(grid) {
						cumLoss += grid[cur.y+i][cur.x].loss
						if i >= minSteps {
							branches = append(branches,
								&Trail{cur.y + i, cur.x, cumLoss, 2})
						}
					}
				}
			}

			return branches
		},
		func(cur *Trail) bool {
			return cur.y == end.y && cur.x == end.x
		},
		func(cur *Trail) int {
			return cur.y<<20 | cur.x<<10 | cur.direction
		},
		func(cur *Trail, _ int) int {
			return cur.cumLoss
		},
		func(cur *Trail) int {
			return common.Manhattan(cur.y, cur.x, end.y, end.x)
		},
		8,
		true,
		false,
	)

	return result.Best.cumLoss
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 94, result)
	} else {
		assert.Equal(t, 1268, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
