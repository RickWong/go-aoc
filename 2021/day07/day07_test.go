package day07

import (
	_ "embed"
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var input string

var data = input

func part1() int {
	items := strings.Split(data, ",")
	positions, minPos, maxPos := parsePositions(items)
	minFuel := 99999999999

	for i := minPos; i < maxPos; i++ {
		fuelUsed := 0
		for _, position := range positions {
			fuelUsed += int(math.Abs(float64(position - i)))
		}
		minFuel = min(minFuel, fuelUsed)
	}

	return minFuel
}

func parsePositions(items []string) ([]int, int, int) {
	positions := make([]int, len(items))
	for idx, item := range items {
		positions[idx], _ = strconv.Atoi(item)
	}

	minPos := lo.Min(positions)
	maxPos := lo.Max(positions)
	return positions, minPos, maxPos
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 37
	if data == input {
		expect = 337833
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	items := strings.Split(data, ",")
	positions, minPos, maxPos := parsePositions(items)
	minFuel := 99999999999

	for i := minPos; i < maxPos; i++ {
		fuelUsed := 0
		for _, position := range positions {
			n := int(math.Abs(float64(position - i)))
			fuelUsed += n * (n + 1) / 2
		}
		minFuel = min(minFuel, fuelUsed)
	}

	return minFuel
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 168
	if data == input {
		expect = 96678050
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
