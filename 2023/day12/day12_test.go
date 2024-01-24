package day12

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"runtime"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type State struct {
	Springs string
	Groups  [32]uint8
}

// Helper functions.

func memoizedCountPossible(cache *map[State]int) func(st State) int {
	var countPossible func(State) int

	countPossible = common.Memo(
		func(st State) int {
			s := st.Springs
			c := st.Groups
			s = strings.TrimLeft(s, ".")

			if len(s) == 0 {
				if c[0] == 0 {
					return 1
				} else {
					return 0
				}
			}

			if c[0] == 0 {
				if strings.Contains(s, "#") {
					return 0
				} else {
					return 1
				}
			}

			if s[0] == '#' {
				if len(s) < int(c[0]) || strings.Contains(s[:c[0]], ".") {
					return 0
				}

				if len(s) == int(c[0]) {
					if c[0] != 0 && c[1] == 0 {
						return 1
					} else {
						return 0
					}
				}

				if s[c[0]] == '#' {
					return 0
				}

				nextC := [32]uint8{}
				copy(nextC[:], c[1:])
				return countPossible(State{s[c[0]+1:], nextC})
			}

			return countPossible(State{"#" + s[1:], c}) + countPossible(State{s[1:], c})
		}, cache)

	return countPossible
}

// Part 1.

func part1() int {
	lines := strings.Split(data, "\n")
	numThreads := runtime.NumCPU()
	eg := errgroup.Group{}
	eg.SetLimit(numThreads)

	caches := make([]map[State]int, numThreads)
	pageSize := len(lines) / numThreads
	results := make([]int, numThreads)
	for i := 0; i < numThreads; i++ {
		i := i
		caches[i] = make(map[State]int, 1024)
		start := i * pageSize
		end := (i + 1) * pageSize
		page := lines[start:end:end]

		eg.Go(func() error {
			for _, line := range page {
				row := strings.Split(line, " ")

				springs := row[0]
				sizes := row[1]
				groups := common.Map(strings.Split(sizes, ","), common.Atoc)
				countPossible := memoizedCountPossible(&caches[i])

				groupsArr := [32]uint8{}
				copy(groupsArr[:], groups)
				res := countPossible(State{springs, groupsArr})
				if res >= 0 {
					results[i] += res
				}
			}
			return nil
		})
	}

	_ = eg.Wait()
	return common.Sum(results)
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

	caches := make([]map[State]int, numThreads)
	pageSize := len(lines) / numThreads
	results := make([]int, numThreads)
	for i := 0; i < numThreads; i++ {
		i := i
		caches[i] = make(map[State]int, 64*1024)
		start := i * pageSize
		end := (i + 1) * pageSize
		page := lines[start:end:end]

		eg.Go(func() error {
			for _, line := range page {
				row := strings.Split(line, " ")

				springs := row[0] + "?" + row[0] + "?" + row[0] + "?" + row[0] + "?" + row[0]
				sizes := row[1] + "," + row[1] + "," + row[1] + "," + row[1] + "," + row[1]
				groups := common.Map(strings.Split(sizes, ","), common.Atoc)
				countPossible := memoizedCountPossible(&caches[i])

				groupsArr := [32]uint8{}
				copy(groupsArr[:], groups)
				res := countPossible(State{springs, groupsArr})
				if res >= 0 {
					results[i] += res
				}
			}
			return nil
		})
	}

	_ = eg.Wait()
	return common.Sum(results)
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
