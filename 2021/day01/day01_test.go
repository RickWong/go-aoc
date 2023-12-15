package day01

import (
	_ "embed"
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
	lines := strings.Split(data, "\n")

	nums := make([]int, len(lines))
	for i, line := range lines {
		nums[i], _ = strconv.Atoi(line)
	}

	count := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}

	return count
}

func TestPart1(t *testing.T) {
	t.Parallel()

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
	lines := strings.Split(data, "\n")

	nums := make([]int, len(lines))
	for i, line := range lines {
		nums[i], _ = strconv.Atoi(line)
	}

	count := 0

	for i := 3; i < len(nums); i++ {
		a := nums[i-3] + nums[i-2] + nums[i-1]
		b := nums[i-2] + nums[i-1] + nums[i]
		if b > a {
			count++
		}
	}

	return count
}

func TestPart2(t *testing.T) {
	t.Parallel()

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
