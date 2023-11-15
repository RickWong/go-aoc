package main

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"
)

//go:embed examples/day01.txt
var example string

//go:embed inputs/day01.txt
var input string

var data = input

func part1() int {
	count := 0
	lines := strings.Split(data, "\n")

	for i := 1; i < len(lines); i++ {
		prev, _ := strconv.Atoi(lines[i-1])
		curr, _ := strconv.Atoi(lines[i])

		if curr > prev {
			count++
		}
	}

	return count
}

func TestPart1(t *testing.T) {
	// How many measurements are larger than the previous measurement?
	result := part1()
	expect := 7
	if data == input {
		expect = 1713
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	count := 0
	lines := strings.Split(data, "\n")

	for i := 3; i < len(lines); i++ {
		a1, _ := strconv.Atoi(lines[i-3])
		a2, _ := strconv.Atoi(lines[i-2])
		a3, _ := strconv.Atoi(lines[i-1])
		a := a1 + a2 + a3
		b1 := a2
		b2 := a3
		b3, _ := strconv.Atoi(lines[i])
		b := b1 + b2 + b3

		if b > a {
			count++
		}
	}

	return count
}

func TestPart2(t *testing.T) {
	// How many sums are larger than the previous sum?
	result := part2()
	expect := 5
	if data == input {
		expect = 1734
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
