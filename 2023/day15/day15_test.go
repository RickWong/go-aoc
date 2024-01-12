package day15

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"regexp"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Lens struct {
	Label string
	Focal int
}

// Helper functions.

func hash(s string) int {
	v := 0
	for _, c := range s {
		v += int(c)
		v *= 17
		v %= 256
	}
	return v
}

// Part 1.

func part1() int {
	steps := strings.Split(strings.TrimSpace(data), ",")
	sum := 0

	for _, step := range steps {
		sum += hash(step)
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 1320, result)
	} else {
		assert.Equal(t, 511257, result)
	}
}

// Part 2.

func part2() int {
	steps := strings.Split(strings.TrimSpace(data), ",")
	boxes := make(map[int][]Lens, 256)
	for b := range boxes {
		boxes[b] = make([]Lens, 16)
	}
	sum := 0

	stepRe := regexp.MustCompile(`(\w+)(-?)(=?)(\d+)?`)
	for _, step := range steps {
		matches := stepRe.FindStringSubmatch(step)
		label, dash, equals, focal := matches[1], matches[2], matches[3], common.Atoi(matches[4])

		boxIdx := hash(label)
		box := boxes[boxIdx]
		lensIdx := slices.IndexFunc(box, func(lens Lens) bool {
			return lens.Label == label
		})

		if dash == "-" && lensIdx >= 0 {
			boxes[boxIdx] = append(box[:lensIdx], box[lensIdx+1:]...)
			continue
		}

		if equals == "=" {
			if lensIdx >= 0 {
				box[lensIdx] = Lens{label, focal}
			} else {
				boxes[boxIdx] = append(box, Lens{label, focal})
			}
		}
	}

	for b, box := range boxes {
		for l, lens := range box {
			sum += (1 + b) * (1 + l) * lens.Focal
		}
	}

	return sum
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 145, result)
	} else {
		assert.Equal(t, 239484, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
