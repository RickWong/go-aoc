package day16

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Beam struct {
	y, x   int
	dy, dx int
}

// Helper functions.

func Map[T, R any](collection []T, fn func(a T) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v)
	}
	return m
}

func makeLut(maxY int) map[int]map[int]int {
	lut := make(map[int]map[int]int)
	for y := 0; y < maxY; y++ {
		lut[y] = make(map[int]int)
	}
	return lut
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := Map(lines, func(line string) []string { return strings.Split(line, "") })
	sum := 0

	beams := make([]Beam, 0, 10)
	beams = append(beams, Beam{0, -1, 0, 1})
	history := make([]Beam, 0, 100)
	lut := makeLut(len(grid))

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		if slices.Contains(history, beam) {
			continue
		}
		history = append(history, beam)

		steps := 0
		for {
			steps++
			y := beam.y + beam.dy*steps
			x := beam.x + beam.dx*steps
			if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
				break
			}

			tile := grid[y][x]
			lut[y][x]++

			if tile == "|" {
				beams = append(beams, Beam{y, x, -1, 0})
				beams = append(beams, Beam{y, x, 1, 0})
				break
			}
			if tile == "-" {
				beams = append(beams, Beam{y, x, 0, -1})
				beams = append(beams, Beam{y, x, 0, 1})
				break
			}
			if tile == "/" {
				if beam.dy != 0 {
					beams = append(beams, Beam{y, x, 0, -beam.dy})
				} else {
					beams = append(beams, Beam{y, x, -beam.dx, 0})
				}
				break
			}
			if tile == "\\" {
				if beam.dy != 0 {
					beams = append(beams, Beam{y, x, 0, beam.dy})
				} else {
					beams = append(beams, Beam{y, x, beam.dx, 0})
				}
				break
			}
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if lut[y][x] > 0 {
				sum++
				print("#")
			} else {
				print(".")
			}
		}
		print("\n")
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 46, result, "Result was incorrect")
	} else {
		assert.Equal(t, 8098, result, "Result was incorrect")
	}
}

// Part 2.

func part2() int {
	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result, "Result was incorrect")
	} else {
		assert.Equal(t, 357134560737, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
