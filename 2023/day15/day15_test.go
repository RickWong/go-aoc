package day15

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Example

// Data types.

// Helper functions.

// Part 1.

func part1() int {
	return 0
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 1320, result, "Result was incorrect")
	} else {
		assert.Equal(t, 105623, result, "Result was incorrect")
	}
}

// Part 2.

func part2() int {
	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result, "Result was incorrect")
	} else {
		assert.Equal(t, 357134560737, result, "Result was incorrect")
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
