package day06

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
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

func part1() int {
	lines := strings.Split(data, "\n")
	times := common.Map(strings.Fields(lines[0])[1:], common.Atoi)
	distances := common.Map(strings.Fields(lines[1])[1:], common.Atoi)
	numWaysToWin := make([]int, len(times))

	for i, time := range times {
		for hold := 0; hold < time; hold++ {
			remain := time - hold
			if hold*remain > distances[i] {
				numWaysToWin[i]++
			}
		}
	}

	return common.Product(numWaysToWin)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 288
	if data == Input {
		expect = 4811940
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	time := common.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	distance := common.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))

	eg := errgroup.Group{}
	numThreads := runtime.NumCPU() * 2
	eg.SetLimit(numThreads)
	numWaysToWin := make([]int, numThreads)

	for i := 0; i < numThreads; i++ {
		i := i
		eg.Go(func() error {
			start := i * (time / numThreads)
			end := (i + 1) * (time / numThreads)
			for hold := start; hold < end; hold++ {
				remain := time - hold
				if hold*remain > distance {
					numWaysToWin[i]++
				}
			}
			return nil
		})
	}

	_ = eg.Wait()

	return common.Sum(numWaysToWin)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 71503
	if data == Input {
		expect = 30077773
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
