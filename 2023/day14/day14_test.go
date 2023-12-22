package day14

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

// Helper functions.

func Map[T, R any](collection []T, fn func(a T) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v)
	}
	return m
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := Map(lines, func(line string) []string { return strings.Split(line, "") })
	sum := 0

	for x := 0; x < len(grid[0]); x++ {
		bottom := 0
		for y := bottom; y < len(grid); y++ {
			switch grid[y][x] {
			case "#":
				bottom = y + 1
			case "O":
				if bottom != y {
					grid[bottom][x] = "O"
					grid[y][x] = "."
				}
				bottom++
			}
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" {
				sum += len(grid) - y
			}
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 136, result, "Result was incorrect")
	} else {
		assert.Equal(t, 105623, result, "Result was incorrect")
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
