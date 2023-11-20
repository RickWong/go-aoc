package day12

import (
	_ "embed"
	utils "github.com/RickWong/go-aoc/2021"
	"maps"
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
	id      string
	tunnels []string
}

type Route struct {
	cave    *Cave
	visited map[string]int
	twice   bool
}

func part1() int {
	lines := strings.Split(data, "\n")
	caves := parseCaves(lines)
	ends := 0

	utils.IterativeSearch(
		&Route{caves["start"], map[string]int{"start": 1}, false},
		func(current *Route) []*Route {
			var nextCaves []*Route
			for _, id := range current.cave.tunnels {
				if id == "start" {
					continue
				}

				if id == "end" {
					ends++
					continue
				}

				if unicode.IsUpper(([]rune(id))[0]) ||
					current.visited[id] < 1 {
					visited := make(map[string]int, len(current.visited))
					maps.Copy(visited, current.visited)
					visited[id] = visited[id] + 1

					nextCaves = append(
						nextCaves,
						&Route{caves[id], visited, false},
					)
				}
			}
			return nextCaves
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

	caves := make(map[string]*Cave)

	for _, line := range lines {
		path := strings.Split(line, "-")
		if path == nil {
			panic("Invalid")
		}

		id := path[0]
		nextCave := path[1]

		if caves[id] == nil {
			caves[id] = &Cave{id, nil}
		}
		caves[id].tunnels = append(caves[id].tunnels, nextCave)

		if caves[nextCave] == nil {
			caves[nextCave] = &Cave{nextCave, nil}
		}
		caves[nextCave].tunnels = append(caves[nextCave].tunnels, id)

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
		&Route{caves["start"], map[string]int{"start": 1}, false},
		func(current *Route) []*Route {
			var nextCaves []*Route
			for _, id := range current.cave.tunnels {
				if id == "start" {
					continue
				}

				if id == "end" {
					ends++
					continue
				}

				isSmallCave := unicode.IsLower(([]rune(id))[0])

				if !isSmallCave ||
					(current.visited[id] == 0) ||
					(current.visited[id] == 1 && !current.twice) {
					visited := make(map[string]int, len(current.visited))
					maps.Copy(visited, current.visited)
					visited[id]++
					twice := current.twice || (isSmallCave && visited[id] == 2)

					nextCaves = append(
						nextCaves,
						&Route{caves[id], visited, twice},
					)
				}
			}
			return nextCaves
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
