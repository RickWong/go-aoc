package day13

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

// Helper functions.

func TransposeStringSlice(slice []string) []string {
	transpose := make([]string, len(slice[0]))
	for i := range slice[0] {
		for j := range slice {
			transpose[i] += string(slice[j][i])
		}
	}
	return transpose
}

// Part 1.

func part1() int {
	patterns := strings.Split(strings.TrimSpace(data), "\n\n")
	sum := 0

	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")

		for i := 0; i < len(lines)-1; i++ {
			if lines[i] == lines[i+1] {
				same := true
				for j := 1; j < min(i+1, len(lines)-i-1); j++ {
					if lines[i-j] != lines[i+1+j] {
						same = false
					}
				}

				if same {
					sum += 100 * (i + 1)
				}
			}
		}

		columns := TransposeStringSlice(lines)

		for i := 0; i < len(columns)-1; i++ {
			if columns[i] == columns[i+1] {
				same := true
				for j := 1; j < min(i+1, len(columns)-i-1); j++ {
					if columns[i-j] != columns[i+1+j] {
						same = false
					}
				}

				if same {
					sum += i + 1
				}
			}
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 405, result, "Result was incorrect")
	} else {
		assert.Equal(t, 9274989, result, "Result was incorrect")
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
