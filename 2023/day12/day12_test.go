package day12

import (
	_ "embed"
	common2 "github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"runtime"
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

func memoizedCountPossible(opts *common2.MemoOptions[int]) func(string, []int) int {
	var countPossible func(string, []int) int

	countPossible = common2.Memo2(
		func(s string, c []int) int {
			s = strings.TrimLeft(s, ".")

			if len(s) == 0 {
				if len(c) == 0 {
					return 1
				} else {
					return 0
				}
			}

			if len(c) == 0 {
				if strings.Contains(s, "#") {
					return 0
				} else {
					return 1
				}
			}

			if s[0] == '#' {
				if len(s) < c[0] || strings.Contains(s[:c[0]], ".") {
					return 0
				}

				if len(s) == c[0] {
					if len(c) == 1 {
						return 1
					} else {
						return 0
					}
				}

				if s[c[0]] == '#' {
					return 0
				}

				return countPossible(s[c[0]+1:], c[1:])
			}

			return countPossible("#"+s[1:], c) + countPossible(s[1:], c)
		}, opts)

	return countPossible
}

// Part 1.

func part1() int {
	lines := strings.Split(data, "\n")
	numThreads := runtime.NumCPU()
	eg := errgroup.Group{}
	eg.SetLimit(numThreads)
	sum := int64(0)

	caches := make([]map[string]int, numThreads)
	pageSize := len(lines) / numThreads
	for i := 0; i < numThreads; i++ {
		i := i
		caches[i] = make(map[string]int, 1024)
		start := i * pageSize
		end := (i + 1) * pageSize
		page := lines[start:end:end]

		eg.Go(func() error {
			for _, line := range page {
				row := strings.Split(line, " ")

				springs := row[0]
				sizes := row[1]
				groups := common2.Map(strings.Split(sizes, ","), common2.Atoi)
				countPossible := memoizedCountPossible(&common2.MemoOptions[int]{Cache: caches[i]})

				res := countPossible(springs, groups)
				if res >= 0 {
					atomic.AddInt64(&sum, int64(res))
				}
			}
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
	lines := strings.Split(data, "\n")
	numThreads := runtime.NumCPU()
	eg := errgroup.Group{}
	eg.SetLimit(numThreads)
	sum := int64(0)

	caches := make([]map[string]int, numThreads)
	pageSize := len(lines) / numThreads
	for i := 0; i < numThreads; i++ {
		i := i
		caches[i] = make(map[string]int, 64*1024)
		start := i * pageSize
		end := (i + 1) * pageSize
		page := lines[start:end:end]

		eg.Go(func() error {
			for _, line := range page {
				row := strings.Split(line, " ")

				springs := row[0] + "?" + row[0] + "?" + row[0] + "?" + row[0] + "?" + row[0]
				sizes := row[1] + "," + row[1] + "," + row[1] + "," + row[1] + "," + row[1]
				groups := common2.Map(strings.Split(sizes, ","), common2.Atoi)
				countPossible := memoizedCountPossible(&common2.MemoOptions[int]{Cache: caches[i]})

				res := countPossible(springs, groups)
				if res >= 0 {
					atomic.AddInt64(&sum, int64(res))
				}
			}
			return nil
		})
	}

	_ = eg.Wait()
	return int(sum)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 525152, result)
	} else {
		assert.Equal(t, 25470469710341, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
