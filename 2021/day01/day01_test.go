package day01

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 7, result)
	} else {
		assert.Equal(t, 1713, result)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 5, result)
	} else {
		assert.Equal(t, 1734, result)
	}
}
