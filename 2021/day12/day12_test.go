package day12

import (
	_ "embed"
	utils "github.com/RickWong/go-aoc/2021/common"
	"strings"
	"testing"
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
	cave    *Cave
	visited int64
	twice   bool
}

func part1() int {
	lines := strings.Split(data, "\n")
	caves := parseCaves(lines)
	ends := 0

	utils.IterativeSearch(
		&Route{caves["start"], caves["start"].id, false},
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

				if isUpper(next.name[0]) ||
					current.visited&next.id == 0 {
					routes = append(
						routes,
						&Route{next,
							current.visited | next.id,
							false},
					)
				}
			}
			return routes
		},
		nil,
		nil,
		nil,
		nil,
		0,
		true,
	)

	return ends
}

func isUpper(c uint8) bool {
	return c >= 65 && c <= 90
}

func parseCaves(lines []string) map[string]*Cave {
	if lines == nil {
		panic("No data")
	}

	caves := make(map[string]*Cave, len(lines))
	id := int64(1)

	for _, line := range lines {
		seperatorIdx := strings.Index(line, "-")
		current := line[0:seperatorIdx]
		next := line[seperatorIdx+1:]

		if caves[current] == nil {
			caves[current] = &Cave{current, id, make([]*Cave, 0, 6)}
			id <<= 1
		}

		if caves[next] == nil {
			caves[next] = &Cave{next, id, make([]*Cave, 0, 6)}
			id <<= 1
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
		&Route{caves["start"], caves["start"].id, false},
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

				isBigCave := isUpper(next.name[0])
				if isBigCave ||
					current.visited&next.id == 0 ||
					!current.twice {
					routes = append(
						routes,
						&Route{
							next,
							current.visited | next.id,
							current.twice || (!isBigCave &&
								current.visited&next.id != 0),
						},
					)
				}
			}
			return routes
		},
		nil,
		nil,
		nil,
		nil,
		0,
		true,
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

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
