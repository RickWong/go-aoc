package day09

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"slices"
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

func nextIteration(outNumbers *[]int) (last int, done bool) {
	// Going from right to left. Overwriting the last item.
	numbers := *outNumbers
	for i := len(numbers) - 1; i > 0; i-- {
		numbers[i] = numbers[i] - numbers[i-1]
	}

	// Trim most-left item to keep items to the right as new iteration.
	numbers = numbers[1:]
	*outNumbers = numbers

	return numbers[len(numbers)-1], Every(numbers, 0)
}

func part1() int {
	lines := strings.Split(data, "\n")
	sum := 0

	for _, line := range lines {
		numbers := Map(strings.Fields(line), Atoi)

		current := numbers
		sum += current[len(current)-1]
		for {
			last, done := nextIteration(&current)
			sum += last

			if done {
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
		// Extrapolating backward is the same as extrapolating forward the reversed sequence.
		// By reversing the sequence we can reuse the exact logic in part 1.
		numbers := Map(strings.Fields(line), Atoi)
		slices.Reverse(numbers)

		current := numbers
		sum += current[len(current)-1]
		for {
			last, done := nextIteration(&current)
			sum += last

			if done {
				break
			}
		}
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