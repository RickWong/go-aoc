package day09

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var input string

var data = input

type Point struct {
	height    int
	neighbors []*Point
}

func parseHeights(lines []string) []*Point {
	if lines == nil {
		panic("No data")
	}

	points := make([]*Point, 0)

	for _, line := range lines {
		for _, heightStr := range line {
			height, _ := strconv.Atoi(string(heightStr))
			point := Point{height, []*Point{}}
			points = append(points, &point)
		}
	}

	for y, line := range lines {
		for x := range line {
			point := points[y*len(line)+x]

			if y > 0 {
				point.neighbors = append(point.neighbors, points[(y-1)*len(line)+x])
			}
			if x > 0 {
				point.neighbors = append(point.neighbors, points[y*len(line)+x-1])
			}
			if y < len(lines)-1 {
				point.neighbors = append(point.neighbors, points[(y+1)*len(line)+x])
			}
			if x < len(line)-1 {
				point.neighbors = append(point.neighbors, points[y*len(line)+x+1])
			}
		}
	}

	return points
}

func part1() int {
	lines := strings.Split(data, "\n")
	points := parseHeights(lines)

	risks := lo.FilterMap(points, func(point *Point, _ int) (int, bool) {
		neighborHeights := lo.Map(point.neighbors, func(point *Point, _ int) int { return point.height })
		if point.height < lo.Min(neighborHeights) {
			return point.height + 1, true
		}

		return -1, false
	})

	return lo.Sum(risks)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 15
	if data == input {
		expect = 560
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	points := parseHeights(lines)
	lows := lo.Filter(points, func(point *Point, _ int) bool {
		neighborHeights := lo.Map(point.neighbors, func(point *Point, _ int) int { return point.height })
		return point.height < lo.Min(neighborHeights)
	})

	basinSizes := make([]int, 0)

	for _, low := range lows {
		queue := []*Point{low}
		var visited []*Point

		for len(queue) > 0 {
			// TODO use ring queue
			current := queue[0]
			queue[0] = nil
			queue = queue[1:]
			if slices.Contains(visited, current) {
				continue
			}
			visited = append(visited, current)

			for _, neighbor := range current.neighbors {
				if neighbor.height < 9 {
					queue = append(queue, neighbor)
				}
			}
		}

		basinSizes = append(basinSizes, len(visited))
	}

	slices.Sort(basinSizes)

	multiplied := 1
	for _, basinSize := range basinSizes[len(basinSizes)-3:] {
		multiplied *= basinSize
	}
	return multiplied
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 1134
	if data == input {
		expect = 959136
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
