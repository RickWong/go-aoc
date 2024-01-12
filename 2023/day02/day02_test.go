package day02

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
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
	lines := strings.Split(data, "\n")
	r := regexp.MustCompile(`(\d+) (\w+)`)
	sum := 0

outer:
	for i, line := range lines {
		draws := r.FindAllStringSubmatch(line, -1)

		for _, draw := range draws {
			count, color := draw[1], draw[2]
			switch {
			case color == "red" && common.Atoi(count) > 12,
				color == "green" && common.Atoi(count) > 13,
				color == "blue" && common.Atoi(count) > 14:
				continue outer
			}
		}

		sum += i + 1
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 8
	if data == Input {
		expect = 2406
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	r := regexp.MustCompile(`(\d+) (\w+)`)
	sum := 0

	for _, line := range lines {
		maxes := make(map[string]int)
		draws := r.FindAllStringSubmatch(line, -1)

		for _, draw := range draws {
			count, color := draw[1], draw[2]
			maxes[color] = max(maxes[color], common.Atoi(count))
		}

		sum += maxes["red"] * maxes["green"] * maxes["blue"]
	}

	return sum
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 2286
	if data == Input {
		expect = 78375
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
