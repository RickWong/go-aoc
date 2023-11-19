package day06

import (
	_ "embed"
	"maps"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var data = input

func part1() int {
	return solve(80)
}

func readData(lines []string) []int {
	if lines == nil {
		panic("No data")
	}

	numsStr := strings.Split(lines[0], ",")
	nums := make([]int, 0, len(numsStr))
	for i := range numsStr {
		num, _ := strconv.ParseInt(numsStr[i], 10, 0)
		nums = append(nums, int(num))
	}
	return nums
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 5934
	if data == input {
		expect = 360268
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	return solve(256)
}

func solve(days int) int {
	timers := readData(strings.Split(data, "\n"))
	occurrences := make(map[int]int)
	tmp := make(map[int]int)

	for _, timer := range timers {
		occurrences[timer]++
	}

	for i := 0; i < days; i++ {
		maps.Copy(tmp, occurrences)

		for k, v := range tmp {
			if k > 0 {
				occurrences[k] -= v
				occurrences[k-1] += v
			} else {
				occurrences[0] -= v
				occurrences[8] += v
				occurrences[6] += v
			}
		}
	}

	sum := 0
	for _, v := range occurrences {
		sum += v
	}

	return sum
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 26984457539
	if data == input {
		expect = 1632146183902
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
