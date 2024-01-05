package day12

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
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

type SearchState struct {
	steps          int
	hashes         int
	numGroupsFound int
}

// Helper functions.

func countPossiblePaths(springs string, groups []int) *common.SearchResult[SearchState, int32] {
	res := common.IterativeSearch[SearchState, int, int32](
		&SearchState{0, 0, 0},
		func(state *SearchState) []*SearchState {
			// Arrived at the end but didn't find all groups.
			if state.steps >= len(springs) {
				return nil
			}

			// Found all groups but not at the end yet. Skip to end, unless more #.
			if state.numGroupsFound == len(groups) {
				if strings.Contains(springs[state.steps:], "#") {
					return nil
				}

				return []*SearchState{
					{len(springs), 0, state.numGroupsFound},
				}
			}

			for i := state.steps; i < len(springs); i++ {
				if springs[i] == '#' {
					if state.hashes < groups[state.numGroupsFound] {
						state.hashes++

						if i == len(springs)-1 && state.hashes == groups[state.numGroupsFound] {
							return []*SearchState{
								{i + 1, 0, state.numGroupsFound + 1},
							}
						}
						continue
					}
					return nil
				}

				if springs[i] == '.' {
					if state.hashes == 0 {
						continue
					}
					if state.hashes == groups[state.numGroupsFound] {
						return []*SearchState{
							{i + 1, 0, state.numGroupsFound + 1},
						}
					}
					return nil
				}

				if springs[i] == '?' {
					branches := make([]*SearchState, 0, 2)

					if state.hashes < groups[state.numGroupsFound] {
						if i == len(springs)-1 && state.hashes+1 == groups[state.numGroupsFound] {
							return []*SearchState{
								{i + 1, 0, state.numGroupsFound + 1},
							}
						}

						branches = append(branches,
							&SearchState{i + 1, state.hashes + 1, state.numGroupsFound})
					}

					if state.hashes == 0 {
						branches = append(branches,
							&SearchState{i + 1, 0, state.numGroupsFound})
					} else if state.hashes == groups[state.numGroupsFound] {
						branches = append(branches,
							&SearchState{i + 1, 0, state.numGroupsFound + 1})
					}

					return branches
				}
			}

			return nil
		},
		func(state *SearchState) bool {
			return state.steps == len(springs) && state.numGroupsFound == len(groups)
		},
		nil,
		nil,
		func(state *SearchState) int32 {
			return -int32(state.numGroupsFound)*10 - int32(state.hashes)
		},
		0,
		false,
		false,
	)
	return res
}

// Part 1.

func part1() int {
	lines := strings.Split(data, "\n")
	eg := errgroup.Group{}
	eg.SetLimit(runtime.NumCPU())
	sum := int64(0)

	for _, line := range lines {
		row := strings.Split(line, " ")

		eg.Go(func() error {
			springs, sizes := row[0], row[1]
			groups := common.Map(strings.Split(sizes, ","), common.Atoi)

			res := countPossiblePaths(springs, groups)
			if res.Paths >= 0 {
				atomic.AddInt64(&sum, int64(res.Paths))
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
	eg := errgroup.Group{}
	eg.SetLimit(runtime.NumCPU())
	sum := int64(0)

	for _, line := range lines {
		row := strings.Split(line, " ")

		eg.Go(func() error {
			springs := row[0] + "?" + row[0] + "?" + row[0] + "?" + row[0] + "?" + row[0]
			sizes := row[1] + "," + row[1] + "," + row[1] + "," + row[1] + "," + row[1]
			groups := common.Map(strings.Split(sizes, ","), common.Atoi)

			res := countPossiblePaths(springs, groups)
			if res.Paths >= 0 {
				atomic.AddInt64(&sum, int64(res.Paths))
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
		assert.Equal(t, 525152, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
