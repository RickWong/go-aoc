package day09

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Helper functions.

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
		numbers := common.Map(strings.Fields(line), common.Atoi)

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
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 114, result, "Result was incorrect")
	} else {
		assert.Equal(t, 1637452029, result, "Result was incorrect")
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(data, "\n")
	sum := 0

	for _, line := range lines {
		// Extrapolating backward is the same as extrapolating forward the reversed sequence.
		// By reversing the sequence we can reuse the exact logic in part 1.
		numbers := common.Map(strings.Fields(line), common.Atoi)
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
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 2, result, "Result was incorrect")
	} else {
		assert.Equal(t, 908, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
