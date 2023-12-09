package day09

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Helper functions.

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Map[T, R any](collection []T, fn func(a T) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v)
	}
	return m
}

func Every(collection []int, is int) bool {
	for _, v := range collection {
		if v != is {
			return false
		}
	}
	return true
}

// Part 1.

func part1() int {
	lines := strings.Split(data, "\n")
	sum := 0

	for _, line := range lines {
		numbers := Map(strings.Fields(line), Atoi)
		current := numbers
		sum += current[len(current)-1]

		for {
			// Going from right to left. Overwriting the last item.
			for i := len(current) - 1; i > 0; i-- {
				current[i] = current[i] - current[i-1]
			}

			// Trim most-left item to items to the right as new iteration.
			current = current[1:]

			sum += current[len(current)-1]

			if Every(current, 0) {
				break
			}
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	result := part1()

	if data == Example {
		assert.Equal(t, 114, result, "Result was incorrect")
	} else {
		assert.Equal(t, 1708206096, result, "Result was incorrect")
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(data, "\n")
	sum := 0

	for _, line := range lines {
		numbers := Map(strings.Fields(line), Atoi)
		current := numbers
		firsts := []int{current[0]}

		for {
			// Going from right to left. Overwriting the last item.
			for i := len(current) - 1; i > 0; i-- {
				current[i] = current[i] - current[i-1]
			}

			// Trim most-left item to items to the right as new iteration.
			current = current[1:]

			firsts = append(firsts, current[0])

			if Every(current, 0) {
				break
			}
		}

		first := 0
		for i := len(firsts) - 1; i > 0; i-- {
			first = firsts[i-1] - first
		}

		sum += first
	}

	return sum
}

func TestPart2(t *testing.T) {
	result := part2()

	if data == Example {
		assert.Equal(t, 2, result, "Result was incorrect")
	} else {
		assert.Equal(t, 1050, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
