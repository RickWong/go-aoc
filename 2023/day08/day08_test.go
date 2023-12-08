package day08

import (
	_ "embed"
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
	result := part1()
	expect := 6
	if data == Input {
		expect = 16343
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(integers ...int) int {
	if len(integers) < 2 {
		return integers[0]
	}

	a, b := integers[0], integers[1]
	result := a * b / gcd(a, b)

	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
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

	steps := make([]int, len(positions))
	for i, position := range positions {
		for {
			L, R := nodes[position][0], nodes[position][1]
			if path[steps[i]%len(path)] == "L" {
				position = L
			} else {
				position = R
			}
			steps[i]++
			if strings.HasSuffix(position, "Z") {
				break
			}
		}
	}

	return lcm(steps...)
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 6
	if data == Input {
		expect = 15299095336639
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
