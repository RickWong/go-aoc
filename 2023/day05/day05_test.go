package day05

import (
	_ "embed"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func fields(s string) []string {
	return strings.Fields(s)
}

func mapFn[T, R any](collection []T, fn func(a T) R) []R {
	return lo.Map[T, R](collection, func(v T, _ int) R {
		return fn(v)
	})
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func part1() int {
	re := regexp.MustCompile(`((?:\d+\s?){3,})`)
	maps := re.FindAllStringSubmatch(data, -1)
	seeds := mapFn(fields(trim(maps[0][0])), atoi)

	for _, m := range maps[1:] {
		lines := lo.Chunk(fields(trim(m[0])), 3)

		for i := range seeds {
			for _, line := range lines {
				dst, src, num := atoi(line[0]), atoi(line[1]), atoi(line[2])
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
	result := part1()
	expect := 35
	if data == Input {
		expect = 31599214
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	re := regexp.MustCompile(`((?:\d+\s?){3,})`)
	maps := re.FindAllStringSubmatch(data, -1)
	ranges := lo.Chunk(mapFn(fields(trim(maps[0][0])), atoi), 2)
	transforms := lo.Reverse(mapFn(maps[1:], func(m []string) [][]int {
		lines := lo.Chunk(fields(trim(m[0])), 3)
		return lo.Reverse(mapFn(lines, func(line []string) []int {
			return []int{atoi(line[0]), atoi(line[1]), atoi(line[2])}
		}))
	}))

	eg := new(errgroup.Group)
	eg.SetLimit(runtime.NumCPU())

	lowest := 1 << 31
	count := 0
	mu := sync.Mutex{}

	for i := 0; i < 100000000000 && count < runtime.NumCPU(); i++ {
		i := i

		eg.Go(func() error {
			v := i
			for _, rules := range transforms {
				for _, rule := range rules {
					from := rule[0]
					till := from + rule[2]
					move := rule[1] - from
					if v >= from && v < till {
						v += move
						break
					}
				}
			}
			for _, r := range ranges {
				if v >= r[0] && v < r[0]+r[1] {
					mu.Lock()
					lowest = min(lowest, i)
					count++
					mu.Unlock()
					break
				}
			}
			return nil
		})
	}

	_ = eg.Wait()
	if count == 0 {
		return 0
	}

	return lowest
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 46
	if data == Input {
		expect = 20358599
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
