package day18

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

type Hole struct {
	y, x  int
	color string
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	deltas := map[byte][2]int{
		'R': {0, 1},
		'D': {1, 0},
		'U': {-1, 0},
		'L': {0, -1},
	}

	y, x := 0, 0
	holes := []Hole{{y, x, ""}}

	// Without using an actual grid, track the min and max x for each line.
	maxY := 0
	lineMins := make(map[int]int)
	lineMaxs := make(map[int]int)
	lineMins[0] = 0
	lineMaxs[0] = 0

	for _, line := range lines {
		parts := strings.Fields(line)
		direction, meters, color := parts[0][0], int(parts[1][0]-'0'), parts[2][1:len(parts[2])-1]

		for i := 0; i < meters; i++ {
			y += deltas[direction][0]
			x += deltas[direction][1]

			maxY = max(maxY, y)
			// Go maps don't return nil. Sadly.
			_, ok := lineMins[y]
			if !ok {
				lineMins[y] = 999999999
			}
			lineMins[y] = min(lineMins[y], x)
			lineMaxs[y] = max(lineMaxs[y], x)

			holes = append(holes, Hole{y, x, color})
		}
	}

	sum := 0
	for y := 0; y <= maxY; y++ {
		sum += lineMaxs[y] - lineMins[y] + 1
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 62, result, "Result was incorrect")
	} else {
		assert.Equal(t, 1790, result, "Result was incorrect")
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
