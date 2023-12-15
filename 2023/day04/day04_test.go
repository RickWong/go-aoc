package day04

import (
	_ "embed"
	"github.com/samber/lo"
	"regexp"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

func part1() int {
	re := regexp.MustCompile(`(?m)Card\s+\d+:\s+((?:\d+ *)+)\s+\|\s+((?:\d+ *)+)`)
	matches := re.FindAllStringSubmatch(data, -1)
	sum := 0

	for _, m := range matches {
		winners, numbers := strings.Fields(m[1]), strings.Fields(m[2])
		numWinners := len(lo.Intersect(winners, numbers))

		sum += (1 << numWinners) >> 1
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 13
	if data == Input {
		expect = 25571
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	re := regexp.MustCompile(`(?m)Card\s+\d+:\s+((?:\d+ *)+)\s+\|\s+((?:\d+ *)+)`)
	matches := re.FindAllStringSubmatch(data, -1)
	copies := make([]int, len(matches))

	for i, m := range matches {
		copies[i]++

		winners, numbers := strings.Fields(m[1]), strings.Fields(m[2])
		numWinners := len(lo.Intersect(winners, numbers))

		for k := 0; k < numWinners; k++ {
			copies[i+k+1] += copies[i]
		}
	}

	return lo.Sum(copies)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 30
	if data == Input {
		expect = 8805731
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
