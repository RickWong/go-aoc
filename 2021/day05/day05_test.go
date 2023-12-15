package day05

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var input string

var data = input

type Vent struct {
	x1, y1, x2, y2 int
}

func part1() int {
	vents, gridWidth, gridHeight := readData(strings.Split(data, "\n"))
	grid := make([]int, gridHeight*gridWidth)
	count := 0

	for _, vent := range vents {
		x1, y1, x2, y2 := vent.x1, vent.y1, vent.x2, vent.y2
		if y1 == y2 {
			count += visitAndCountHorizontal(x1, x2, y1, grid, gridWidth)
		} else if x1 == x2 {
			count += visitAndCountVertical(y1, y2, x1, grid, gridWidth)
		}
	}

	return count
}

func readData(lines []string) ([]Vent, int, int) {
	vents := make([]Vent, 0, len(lines))
	gridWidth := 0
	gridHeight := 0

	for _, line := range lines {
		var x1, y1, x2, y2 int
		_, _ = fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		vents = append(vents, Vent{x1, y1, x2, y2})

		gridWidth = max(gridWidth, x1, x2)
		gridHeight = max(gridHeight, y1, y2)
	}

	gridWidth++
	gridHeight++

	return vents, gridWidth, gridHeight
}

func visitAndCountVertical(y1 int, y2 int, x1 int, grid []int, gridWidth int) int {
	count := 0
	step := int(math.Copysign(1, float64(y2-y1)))
	for y := y1; y != y2+step; y += step {
		grid[y*gridWidth+x1]++
		if grid[y*gridWidth+x1] == 2 {
			count++
		}
	}
	return count
}

func visitAndCountHorizontal(x1 int, x2 int, y1 int, grid []int, gridWidth int) int {
	count := 0
	step := int(math.Copysign(1, float64(x2-x1)))
	for x := x1; x != x2+step; x += step {
		grid[y1*gridWidth+x]++
		if grid[y1*gridWidth+x] == 2 {
			count++
		}
	}
	return count
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 5
	if data == input {
		expect = 6267
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	vents, gridWidth, gridHeight := readData(strings.Split(data, "\n"))
	grid := make([]int, gridHeight*gridWidth)
	count := 0

	for _, vent := range vents {
		x1, y1, x2, y2 := vent.x1, vent.y1, vent.x2, vent.y2
		if y1 == y2 {
			count += visitAndCountHorizontal(x1, x2, y1, grid, gridWidth)
		} else if x1 == x2 {
			count += visitAndCountVertical(y1, y2, x1, grid, gridWidth)
		} else {
			count += visitAndCountDiagonal(x1, y1, x2, y2, grid, gridWidth)
		}
	}

	return count
}

func visitAndCountDiagonal(x1 int, y1 int, x2 int, y2 int, grid []int, gridWidth int) int {
	count := 0
	stepY := int(math.Copysign(1, float64(y2-y1)))
	stepX := int(math.Copysign(1, float64(x2-x1)))
	for i := 0; i <= int(math.Abs(float64(y2-y1))); i++ {
		y := y1 + i*stepY
		x := x1 + i*stepX
		grid[y*gridWidth+x]++
		if grid[y*gridWidth+x] == 2 {
			count++
		}
	}
	return count
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 12
	if data == input {
		expect = 20196
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
