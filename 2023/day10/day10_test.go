package day10

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
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

func parsePoints(lines []string) (start *Point, points []*Point) {
	height := len(lines)
	width := len(lines[0])
	points = make([]*Point, height*width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			points[idx] = &Point{y, x, make([]*Point, 0, 4), lines[y][x], 0, false}

			if start == nil && lines[y][x] == 'S' {
				start = points[idx]
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

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			p := points[idx]

			for _, a := range AROUND {
				if y+a.y < 0 || y+a.y >= len(lines) || x+a.x < 0 || x+a.x >= len(lines[y]) {
					continue
				}

				idx := (y+a.y)*width + (x + a.x)
				next := points[idx]
				if p.tile == a.self[0] || p.tile == a.self[1] || p.tile == a.self[2] || p.tile == a.self[3] {
					if next.tile == a.match[0] || next.tile == a.match[1] || next.tile == a.match[2] {
						p.next = append(p.next, next)
					}
				}
			}
		}
	}

	safeGet := func(y, x int) byte {
		if y < 0 || y >= height || x < 0 || x >= width {
			return 0
		}
		idx := y*width + x
		return points[idx].tile
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
	visited := bitmap.Bitmap{}
	maxDistance := 0

	for len(unvisited) > 0 {
		current := unvisited[0]
		unvisited[0] = nil
		unvisited = unvisited[1:]
		visited.Set(uint32(current.y<<8 | current.x))

		current.loop = true
		maxDistance = max(maxDistance, current.distance)

		for _, next := range current.next {
			if !visited.Contains(uint32(next.y<<8 | next.x)) {
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
	//idx := y*width + x
	//		if points[idx].loop {
	//			print(points[idx].tile)
	//		} else {
	//			print("_")
	//		}
	//	}
	//	println()
	//}

	endOfCorner := map[byte]byte{'F': '7', 'L': 'J'}
	n := 0

	height := len(lines)
	width := len(lines[0])

	for y := 0; y < height; y++ {
		corner := byte(0)
		inside := false
		m := 0

		for x := 0; x < width; x++ {
			idx := y*width + x
			// Inside loop and not part of the loop.
			if inside && !points[idx].loop {
				m++
				continue
			}

			// Crossing part of the loop.
			if points[idx].loop {
				switch points[idx].tile {
				// Loop border.
				case '|':
					inside = !inside
					corner = 0
				// Track start of a corner.
				case 'F', 'L':
					if corner == 0 {
						inside = !inside
						corner = points[idx].tile
					}
				// End of a corner.
				case '7', 'J':
					if points[idx].tile == endOfCorner[corner] {
						inside = !inside
					}
					corner = 0
				}
			}
		}
		// println("row", y, "has", m, "inside", corner)
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
