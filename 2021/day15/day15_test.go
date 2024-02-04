package day15

import (
	_ "embed"
	utils "github.com/RickWong/go-aoc/common"
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Position struct {
	y, x  int
	risk  int
	steps []*Position
}

func part1() int {
	lines := strings.Split(data, "\n")
	grid := parseGrid(lines, 1)
	start := grid[0]
	end := grid[len(grid)-1]

	result := utils.IterativeSearch(
		start,
		func(p *Position) []*Position {
			return p.steps
		},
		func(p *Position) bool {
			return p == end
		},
		func(p *Position) int {
			return (p.y << 16) + p.x
		},
		func(p *Position, curWeight float64) float64 {
			return curWeight + float64(p.risk)
		},
		func(p *Position) float64 {
			return math.Abs(float64(end.y-p.y)) +
				math.Abs(float64(end.x-p.x))
		},
		0,
		true,
		false,
	)

	return lo.Sum(lo.Map(result.BestPath[1:], func(p *Position, _ int) int {
		return p.risk
	}))
}

func parseGrid(lines []string, repeat int) []*Position {
	if lines == nil {
		panic("no data")
	}

	height := len(lines) * repeat
	width := len(lines[0]) * repeat
	grid := make([]*Position, height*width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			riskStr := lines[y%len(lines)][x%len(lines[0]) : x%len(lines[0])+1]
			risk, _ := strconv.Atoi(riskStr)
			risk = (risk+y/len(lines)+x/len(lines[0])-1)%9 + 1
			grid[y*width+x] = &Position{y, x, risk, nil}
		}
	}

	for _, p := range grid {
		if p.y >= 1 {
			p.steps = append(p.steps, grid[(p.y-1)*width+p.x])
		}
		if p.y < height-1 {
			p.steps = append(p.steps, grid[(p.y+1)*width+p.x])
		}
		if p.x >= 1 {
			p.steps = append(p.steps, grid[p.y*width+p.x-1])
		}
		if p.x < width-1 {
			p.steps = append(p.steps, grid[p.y*width+p.x+1])
		}
	}

	return grid
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 40
	if data == Input {
		expect = 540
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	grid := parseGrid(lines, 5)
	start := grid[0]
	end := grid[len(grid)-1]

	result := utils.IterativeSearch(
		start,
		func(p *Position) []*Position {
			return p.steps
		},
		func(p *Position) bool {
			return p == end
		},
		func(p *Position) int {
			return (p.y << 16) + p.x
		},
		func(p *Position, curWeight float64) float64 {
			return curWeight + float64(p.risk)
		},
		func(p *Position) float64 {
			return math.Abs(float64(end.y-p.y)) +
				math.Abs(float64(end.x-p.x))
		},
		0,
		true,
		false,
	)

	return lo.Sum(lo.Map(result.BestPath[1:], func(p *Position, _ int) int {
		return p.risk
	}))
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 315
	if data == Input {
		expect = 2879
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
