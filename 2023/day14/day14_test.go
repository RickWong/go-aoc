package day14

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"hash/crc32"
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

func printAndSum(grid [][]string) int {
	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "O" {
				sum += len(grid) - y
			}
			//print(grid[y][x])
		}
		//println()
	}
	return sum
}

func tiltEast(grid [][]string) {
	for y := 0; y < len(grid[0]); y++ {
		right := len(grid) - 1
		for x := right; x >= 0; x-- {
			switch grid[y][x] {
			case "#":
				right = x - 1
			case "O":
				if right != x {
					grid[y][right] = "O"
					grid[y][x] = "."
				}
				right--
			}
		}
	}
}

func tiltSouth(grid [][]string) {
	for x := 0; x < len(grid[0]); x++ {
		top := len(grid) - 1
		for y := top; y >= 0; y-- {
			switch grid[y][x] {
			case "#":
				top = y - 1
			case "O":
				if top != y {
					grid[top][x] = "O"
					grid[y][x] = "."
				}
				top--
			}
		}
	}
}

func tiltWest(grid [][]string) {
	for y := 0; y < len(grid[0]); y++ {
		left := 0
		for x := left; x < len(grid); x++ {
			switch grid[y][x] {
			case "#":
				left = x + 1
			case "O":
				if left != x {
					grid[y][left] = "O"
					grid[y][x] = "."
				}
				left++
			}
		}
	}
}

func tiltNorth(grid [][]string) {
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
		assert.Equal(t, 136, result)
	} else {
		assert.Equal(t, 113424, result)
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := Map(lines, func(line string) []string { return strings.Split(line, "") })

	genCycles := 1000
	checksums := make(map[uint32]int, genCycles)
	sums := make([]int, genCycles)

	for i := 0; i < genCycles; i++ {
		tiltNorth(grid)
		tiltWest(grid)
		tiltSouth(grid)
		tiltEast(grid)

		bytes := []byte(
			strings.Join(
				Map(grid, func(row []string) string { return strings.Join(row, "") }),
				"",
			),
		)

		checksum := crc32.ChecksumIEEE(bytes)
		if start, found := checksums[checksum]; found {
			return sums[start+((1000000000-start)%(i-start))-1]
		}

		checksums[checksum] = i
		sums[i] = printAndSum(grid)
	}

	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 64, result)
	} else {
		assert.Equal(t, 96003, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
