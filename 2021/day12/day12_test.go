package day12

import (
	_ "embed"
	utils "github.com/RickWong/go-aoc/2021/common"
	"strings"
	"testing"
	"unicode"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Cave struct {
	name    string
	id      int64
	tunnels []*Cave
}

type Route struct {
	cave         *Cave
	visitedOnce  int64
	visitedTwice int64
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func part1() int {
	lines := strings.Split(data, "\n")
	caves := parseCaves(lines)
	ends := 0

	utils.IterativeSearch(
		&Route{caves["start"], caves["start"].id, 0},
		func(current *Route) []*Route {
			routes := make([]*Route, 0, len(current.cave.tunnels))
			for _, next := range current.cave.tunnels {
				if next.name == "start" {
					continue
				}

				if next.name == "end" {
					ends++
					continue
				}

				if current.visitedOnce&current.cave.id == 0 || unicode.
					IsUpper(([]rune(next.name))[0]) {
					routes = append(
						routes,
						&Route{next, current.visitedOnce | current.cave.
							id, current.visitedTwice},
					)
				}
			}
			return routes
		},
		func(current *Route) bool { return false },
		nil,
		nil,
		nil,
		0,
	)

	return ends
}

func parseCaves(lines []string) map[string]*Cave {
	if lines == nil {
		panic("No data")
	}

	caves := make(map[string]*Cave, len(lines))

	for idx, line := range lines {
		path := strings.Split(line, "-")
		if path == nil {
			panic("Invalid")
		}

		id := int64(1 << idx)
		current := path[0]
		next := path[1]

		if caves[current] == nil {
			caves[current] = &Cave{current, id, nil}
		}

		if caves[next] == nil {
			caves[next] = &Cave{next, id, nil}
		}

		caves[current].tunnels = append(caves[current].tunnels, caves[next])
		caves[next].tunnels = append(caves[next].tunnels, caves[current])
	}
	return caves
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 226
	if data == Input {
		expect = 4338
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	caves := parseCaves(lines)
	ends := 0

	utils.IterativeSearch(
		&Route{caves["start"], caves["start"].id, 0},
		func(current *Route) []*Route {
			var routes []*Route
			for _, next := range current.cave.tunnels {
				if next.name == "start" {
					continue
				}

				if next.name == "end" {
					ends++
					continue
				}

				isSmallCave := unicode.IsLower(([]rune(next.name))[0])

				if !isSmallCave ||
					(current.visitedOnce&current.cave.id == 0) ||
					(current.visitedOnce&current.cave.id == 1 && current.
						visitedTwice&current.cave.id == 0) {
					visitedTwice := current.visitedTwice
					if current.visitedOnce&current.cave.id == 1 {
						visitedTwice |= current.cave.id
					}
					routes = append(
						routes,
						&Route{next, current.visitedOnce | current.cave.
							id, visitedTwice},
					)
				}
			}
			return routes
		},
		func(current *Route) bool { return false },
		nil,
		nil,
		nil,
		0,
	)

	return ends
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 3509
	if data == Input {
		expect = 114189
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
