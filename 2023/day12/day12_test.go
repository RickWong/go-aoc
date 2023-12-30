package day12

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

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

func validate(springs string, groups []int) bool {
	curGroup := 0
	curSize := 0

	for index := range springs {
		switch springs[index] {
		case '?':
			return false
		case '#':
			curSize++
		}

		if springs[index] == '.' || index == len(springs)-1 {
			if curSize > 0 {
				// Found more groups than expected. Or current length doesn't match.
				if curGroup >= len(groups) || curSize != groups[curGroup] {
					return false
				}

				// Reset to count the next group.
				curSize = 0
				curGroup++
			}
		}
	}

	return curGroup == len(groups)
}

// Part 1.

func part1() int {
	lines := strings.Split(data, "\n")
	sum := int64(0)
	eg := errgroup.Group{}
	eg.SetLimit(64)

	for _, line := range lines {
		line := line

		eg.Go(func() error {
			row := strings.Split(line, " ")
			springs, sizes_ := row[0], row[1]
			sizes := Map(strings.Split(sizes_, ","), Atoi)

			common.IterativeSearch(
				&springs,
				func(s *string) []*string {
					index := strings.Index(*s, "?")
					if index >= 0 {
						branches := make([]*string, 0, 2)
						variant1 := (*s)[:index] + "#" + (*s)[index+1:]
						variant2 := (*s)[:index] + "." + (*s)[index+1:]
						branches = append(branches, &variant1, &variant2)
						return branches
					} else if validate(*s, sizes) {
						atomic.AddInt64(&sum, 1)
					}
					return nil
				},
				nil,
				nil,
				nil,
				nil,
				0,
				true,
				false,
			)

			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		panic(err)
	}

	return int(sum)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 21, result)
	} else {
		assert.Equal(t, 7286, result)
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
		assert.Equal(t, 82000210, result)
	} else {
		assert.Equal(t, 357134560737, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
