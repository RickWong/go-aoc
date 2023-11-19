package day12

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Example

func part1() int {
	return 0
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 1656
	if data == Input {
		expect = 1723
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	return 0
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 195
	if data == Input {
		expect = 327
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
