package day12

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
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

type Variant struct {
	springs      string
	nextWildcard int
}

// Helper functions.

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

	for _, line := range lines {
		line := line

		eg.Go(func() error {
			row := strings.Split(line, " ")
			springs, sizes_ := row[0], row[1]
			sizes := common.Map(strings.Split(sizes_, ","), common.Atoi)

			common.IterativeSearch[Variant, int, int](
				&Variant{springs, strings.Index(springs, "?")},
				func(v *Variant) []*Variant {
					if v.nextWildcard >= 0 {
						nextWildcard := strings.Index(v.springs[v.nextWildcard+1:], "?")
						if nextWildcard >= 0 {
							nextWildcard += v.nextWildcard + 1
						}
						bs := []byte(v.springs)
						bs[v.nextWildcard] = '#'
						variant1 := Variant{string(bs), nextWildcard}
						bs[v.nextWildcard] = '.'
						variant2 := Variant{string(bs), nextWildcard}
						return []*Variant{&variant1, &variant2}
					}

					if validate(v.springs, sizes) {
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

	_ = eg.Wait()
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
