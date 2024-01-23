package day05

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
	"regexp"
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

func part1() int {
	re := regexp.MustCompile(`((?:\d+\s?){3,})`)
	maps := re.FindAllStringSubmatch(data, -1)
	seeds := common.Map(strings.Fields(strings.TrimSpace(maps[0][0])), common.Atoi)

	for _, m := range maps[1:] {
		lines := lo.Chunk(strings.Fields(strings.TrimSpace(m[0])), 3)

		for i := range seeds {
			for _, line := range lines {
				dst, src, num := common.Atoi(line[0]), common.Atoi(line[1]), common.Atoi(line[2])
				v := seeds[i]

				if v >= src && v < src+num {
					seeds[i] += dst - src
					break
				}
			}
		}
	}

	return lo.Min(seeds)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 35
	if data == Input {
		expect = 265018614
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	re := regexp.MustCompile(`((?:\d+\s?){3,})`)
	maps := re.FindAllStringSubmatch(data, -1)
	ranges := lo.Chunk(common.Map(strings.Fields(strings.TrimSpace(maps[0][0])), common.Atoi), 2)
	transforms := lo.Reverse(common.Map(maps[1:], func(m []string) [][]int {
		lines := lo.Chunk(strings.Fields(strings.TrimSpace(m[0])), 3)
		return lo.Reverse(common.Map(lines, func(line []string) []int {
			return []int{common.Atoi(line[0]), common.Atoi(line[1]), common.Atoi(line[2])}
		}))
	}))

	runTransforms := func(input int) int {
		value := input
		for _, rules := range transforms {
			for _, rule := range rules {
				from := rule[0]
				till := from + rule[2]
				move := rule[1] - from
				if value >= from && value < till {
					value += move
					break
				}
			}
		}
		for _, r := range ranges {
			if value >= r[0] && value < r[0]+r[1] {
				return input
			}
		}
		return -1
	}

	eg := errgroup.Group{}
	firstHit := int64(0)

	// Use all cores and big steps to find the first hit.
	numThreads := runtime.NumCPU()
	stepSize := 1000
	for i := 0; i < numThreads; i++ {
		i := i
		eg.Go(func() error {
			// Use steps of 10 to find the first hit.
			for j := 0; j < 1000*stepSize; j += 10 {
				if j%numThreads == i && runTransforms(j*stepSize) >= 0 {
					atomic.StoreInt64(&firstHit, int64(j*stepSize))
					break
				}
			}
			return nil
		})
	}

	_ = eg.Wait()

	// Go backwards to find the lowest result.
	lowest := int(firstHit)
	for ; lowest > 0; lowest-- {
		if runTransforms(lowest-1) == -1 {
			break
		}
	}

	return lowest
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 46
	if data == Input {
		expect = 63179500
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
