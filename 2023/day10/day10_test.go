package day10

import (
	_ "embed"
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
	y, x     int
	next     []*Point
	tile     byte
	distance int
	loop     bool
}

// Helper functions.

// Part 1.

func parsePoints(lines []string) (start *Point, points [][]*Point) {
	points = make([][]*Point, len(lines))

	for y := 0; y < len(lines); y++ {
		points[y] = make([]*Point, len(lines[y]))

		for x := 0; x < len(lines[y]); x++ {
			points[y][x] = &Point{y, x, make([]*Point, 0, 4), lines[y][x], 0, false}

			if start == nil && lines[y][x] == 'S' {
				start = points[y][x]
			}
		}
	}

	AROUND := []struct {
		y, x  int
		self  []byte
		match []byte
	}{
		{1, 0, []byte("S|7F"), []byte("|JL")},
		{-1, 0, []byte("S|JL"), []byte("|7F")},
		{0, 1, []byte("S-LF"), []byte("-J7")},
		{0, -1, []byte("S-J7"), []byte("-LF")},
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			p := points[y][x]

			for _, a := range AROUND {
				if y+a.y < 0 || y+a.y >= len(lines) || x+a.x < 0 || x+a.x >= len(lines[y]) {
					continue
				}

				next := points[y+a.y][x+a.x]
				if p.tile == a.self[0] || p.tile == a.self[1] || p.tile == a.self[2] || p.tile == a.self[3] {
					if next.tile == a.match[0] || next.tile == a.match[1] || next.tile == a.match[2] {
						p.next = append(p.next, next)
					}
				}
			}
		}
	}

	safeGet := func(y, x int) byte {
		if y < 0 || y >= len(points) || x < 0 || x >= len(points[y]) {
			return 0
		}
		return points[y][x].tile
	}

	belowStart := safeGet(start.y+1, start.x)
	rightStart := safeGet(start.y, start.x+1)
	aboveStart := safeGet(start.y-1, start.x)
	leftStart := safeGet(start.y, start.x-1)

	switch leftStart {
	case '-', 'F', 'L':
		switch {
		case aboveStart == '|', aboveStart == '7', aboveStart == 'F':
			start.tile = 'J'
		case belowStart == '|', belowStart == 'J', belowStart == 'L':
			start.tile = '7'
		}
	}

	switch rightStart {
	case '-', 'J', '7':
		switch {
		case aboveStart == '|', aboveStart == '7', aboveStart == 'F':
			start.tile = 'L'
		case belowStart == '|', belowStart == 'J', belowStart == 'L':
			start.tile = 'F'
		}
	}

	return start, points
}

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	start, _ := parsePoints(lines)
	maxDistance := calculateDistances(start, len(lines)*len(lines[0]))
	return maxDistance
}

func calculateDistances(start *Point, size int) int {
	unvisited := make([]*Point, 0, size)
	unvisited = append(unvisited, start)
	visited := make(map[*Point]bool, size)
	maxDistance := 0

	for len(unvisited) > 0 {
		current := unvisited[0]
		unvisited = unvisited[1:]
		visited[current] = true

		current.loop = true

		for _, next := range current.next {
			if !visited[next] {
				maxDistance = max(maxDistance, current.distance+1)
				next.distance = current.distance + 1
				unvisited = append(unvisited, next)
			}
		}
	}

	return maxDistance
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 80, result)
	} else {
		assert.Equal(t, 6778, result)
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	start, points := parsePoints(lines)
	calculateDistances(start, len(lines)*len(lines[0]))

	//for y := 0; y < len(lines); y++ {
	//	for x := 0; x < len(points[y]); x++ {
	//		if points[y][x].loop {
	//			print(points[y][x].tile)
	//		} else {
	//			print("_")
	//		}
	//	}
	//	println()
	//}

	endOfCorner := map[byte]byte{'F': '7', 'L': 'J'}
	n := 0

	for y := 0; y < len(lines); y++ {
		corner := byte(0)
		inside := false
		m := 0

		for x := 0; x < len(lines[y]); x++ {
			// Inside loop and not part of the loop.
			if inside && !points[y][x].loop {
				m++
				continue
			}

			// Crossing part of the loop.
			if points[y][x].loop {
				switch points[y][x].tile {
				// Loop border.
				case '|':
					inside = !inside
					corner = 0
				// Track start of a corner.
				case 'F', 'L':
					if corner == 0 {
						inside = !inside
						corner = points[y][x].tile
					}
				// End of a corner.
				case '7', 'J':
					if points[y][x].tile == endOfCorner[corner] {
						inside = !inside
					}
					corner = 0
				}
			}
		}
		//println("row", y, "has", m, "inside", corner)
		n += m
	}

	return n
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 10, result)
	} else {
		assert.Equal(t, 433, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
