package day06

import (
	_ "embed"
	"github.com/samber/lo"
	"strconv"
	"strings"
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

func product(collection []int) int {
	m := 1
	for _, v := range collection {
		m *= v
	}
	return m
}

func part1() int {
	lines := strings.Split(data, "\n")
	times := mapFn(fields(lines[0])[1:], atoi)
	distances := mapFn(fields(lines[1])[1:], atoi)
	numWaysToWin := make([]int, len(times))

	for i, time := range times {
		for hold := 0; hold < time; hold++ {
			remain := time - hold
			if hold*remain > distances[i] {
				numWaysToWin[i]++
			}
		}
	}

	return product(numWaysToWin)
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 288
	if data == Input {
		expect = 227850
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	time := atoi(strings.Join(fields(lines[0])[1:], ""))
	distance := atoi(strings.Join(fields(lines[1])[1:], ""))
	numWaysToWin := 0

	for hold := 0; hold < time; hold++ {
		remain := time - hold
		if hold*remain > distance {
			numWaysToWin++
		}
	}

	return numWaysToWin
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 71503
	if data == Input {
		expect = 42948149
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
