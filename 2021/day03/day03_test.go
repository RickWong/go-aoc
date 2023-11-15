package day03

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var data = input

func part1() int {
	return -1
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 15 * 10
	if data == input {
		expect = 2065 * 917 // 1_893_605
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func BenchmarkPart1(b *testing.B) {
	part1()
}

func part2() int {
	return -1
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 15 * 60
	if data == input {
		expect = 2_120_734_350
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}

}

func BenchmarkPart2(b *testing.B) {
	part2()
}
