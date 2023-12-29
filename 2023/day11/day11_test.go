package day11

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"math"
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
	y, x   int
	galaxy bool
}

// Helper functions.

func parseGridAndGalaxies(lines []string, repeatSpace int) ([][]Point, []*Point) {
	grid := make([][]Point, len(lines))
	galaxies := make([]*Point, 0, 100)
	rowsWithGalaxy := make(map[int]bool)
	colsWithGalaxy := make(map[int]bool)

	for y, line := range lines {
		grid[y] = make([]Point, len(line))

		for x, char := range line {
			grid[y][x] = Point{y, x, char == '#'}

			if grid[y][x].galaxy {
				galaxies = append(galaxies, &grid[y][x])
				rowsWithGalaxy[y] = true
				colsWithGalaxy[x] = true
			}
		}
	}

	// Don't grow the grid, just move the galaxies to the right and down.

	offset := 0
	for y := range lines {
		if !rowsWithGalaxy[y] {
			for _, g := range galaxies {
				if g.y > y+offset {
					g.y += repeatSpace - 1
				}
			}
			offset += repeatSpace - 1
		}
	}

	offset = 0
	for x := range lines[0] {
		if !colsWithGalaxy[x] {
			for _, g := range galaxies {
				if g.x > x+offset {
					g.x += repeatSpace - 1
				}
			}
			offset += repeatSpace - 1
		}
	}

	return grid, galaxies
}

func stepsBetweenGalaxies(start *Point, end *Point) int {
	return int(math.Abs(float64(end.y-start.y)) + math.Abs(float64(end.x-start.x)))
}

// Part 1.

func part1() int {
	lines := strings.Split(data, "\n")
	_, galaxies := parseGridAndGalaxies(lines, 2)

	sum := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += stepsBetweenGalaxies(galaxies[i], galaxies[j])
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 374, result, "Result was incorrect")
	} else {
		assert.Equal(t, 9563821, result, "Result was incorrect")
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(data, "\n")
	_, galaxies := parseGridAndGalaxies(lines, 1000000)
	sum := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += stepsBetweenGalaxies(galaxies[i], galaxies[j])
		}
	}

	return sum
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result, "Result was incorrect")
	} else {
		assert.Equal(t, 827009909817, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
