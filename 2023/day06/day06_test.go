package day06

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"math"
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
	time := common.Atof(strings.Join(strings.Fields(lines[0])[1:], ""))
	distance := common.Atof(strings.Join(strings.Fields(lines[1])[1:], ""))

	maxRoot := math.Floor(time + math.Sqrt(time*time-4*distance)/2)
	minRoot := math.Ceil(time - math.Sqrt(time*time-4*distance)/2)

	return int(maxRoot-minRoot) + 1
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
