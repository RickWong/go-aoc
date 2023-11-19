package day10

import (
	_ "embed"
	"github.com/samber/lo"
	stack2 "github.com/zyedidia/generic/stack"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var data = input

var pairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var corruptionPoints = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var incompletePoints = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func part1() int {
	lines := strings.Split(data, "\n")
	scores := make([]int, 0)

	for _, line := range lines {
		stack := stack2.New[string]()

		for _, c := range strings.Split(line, "") {
			if len(pairs[c]) > 0 {
				stack.Push(pairs[c])
				continue
			}

			if stack.Pop() != c {
				scores = append(scores, corruptionPoints[c])
				break
			}
		}
	}

	return lo.Sum(scores)
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 26397
	if data == input {
		expect = 288291
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	scores := make([]int, 0)

	for _, line := range lines {
		stack := stack2.New[string]()
		corrupt := false
		score := 0

		for _, c := range strings.Split(line, "") {
			if len(pairs[c]) > 0 {
				stack.Push(pairs[c])
				continue
			}

			if stack.Pop() != c {
				corrupt = true
				break
			}
		}

		if corrupt {
			continue
		}

		for {
			c := stack.Pop()
			if len(c) == 0 {
				break
			}

			score *= 5
			score += incompletePoints[c]
		}

		scores = append(scores, score)
	}

	slices.Sort(scores)
	return scores[len(scores)/2]
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 288957
	if data == input {
		expect = 820045242
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
