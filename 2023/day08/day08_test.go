package day08

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
	path := strings.Split(lines[0], "")
	steps := 0

	nodes := make(map[string][]string, len(lines[2:]))
	re := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	for _, line := range lines[2:] {
		matches := re.FindStringSubmatch(line)
		nodes[matches[1]] = []string{matches[2], matches[3]}
	}

	position := "AAA"
	for {
		L, R := nodes[position][0], nodes[position][1]
		if path[steps%len(path)] == "L" {
			position = L
		} else {
			position = R
		}
		steps++
		if position == "ZZZ" {
			break
		}
	}

	return steps
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 6
	if data == Input {
		expect = 12737
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	path := strings.Split(lines[0], "")

	nodes := make(map[string][]string, len(lines[2:]))
	re := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	for _, line := range lines[2:] {
		matches := re.FindStringSubmatch(line)
		nodes[matches[1]] = []string{matches[2], matches[3]}
	}

	positions := make([]string, 0, len(nodes)/2)
	for position := range nodes {
		if strings.HasSuffix(position, "A") {
			positions = append(positions, position)
		}
	}

	lcm := common.NewLCMCalculator[int](len(positions))
	for idx, position := range positions {
		steps := 0
		for {
			L, R := nodes[position][0], nodes[position][1]
			if path[steps%len(path)] == "L" {
				position = L
			} else {
				position = R
			}
			steps++
			if strings.HasSuffix(position, "Z") {
				break
			}
		}
		lcm.Detect(idx, steps)
	}

	return lcm.Calc()
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 6
	if data == Input {
		expect = 9064949303801
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
