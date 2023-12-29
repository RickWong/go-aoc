package day19

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Workflow struct {
	name     string
	rules    []Rule
	fallback string
}

type Rule struct {
	attribute  byte
	comparator byte
	value      int
	workflow   string
}

type Rating = map[byte]int

// Helper functions.

// Part 1.

func part1() int {
	blocks := strings.Split(strings.TrimSpace(data), "\n\n")
	workflowsRegex := regexp.MustCompile(`(?m)^(\w+){((?:\w+[<>]\d+:\w+,?)+\w+)}$`)
	ruleRegex := regexp.MustCompile(`(\w+)([<>])(\d+):(\w+)`)
	ratingsRegex := regexp.MustCompile(`(?m)^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
	workflows := make(map[string]Workflow)
	ratings := make([]Rating, 0)

	for _, match := range workflowsRegex.FindAllStringSubmatch(blocks[0], -1) {
		name, rules := match[1], match[2]
		workflow := Workflow{name, nil, ""}

		for _, rule := range strings.Split(rules, ",") {
			parts := ruleRegex.FindStringSubmatch(rule)
			if parts == nil {
				workflow.fallback = rule
				break
			}

			attribute, comparator, value, next := parts[1][0], parts[2][0], Atoi(parts[3]), parts[4]
			workflow.rules = append(workflow.rules,
				Rule{attribute, comparator, value, next})
		}

		workflows[name] = workflow
	}

	for _, match := range ratingsRegex.FindAllStringSubmatch(blocks[1], -1) {
		x, m, a, s := Atoi(match[1]), Atoi(match[2]), Atoi(match[3]), Atoi(match[4])
		ratings = append(ratings, Rating{'x': x, 'm': m, 'a': a, 's': s})
	}

	sum := 0
	for _, rating := range ratings {
		current := "in"

		for current != "A" && current != "R" {
			next := ""
			for _, rule := range workflows[current].rules {
				if rule.comparator == '<' && rating[rule.attribute] < rule.value {
					next = rule.workflow
					break
				}
				if rule.comparator == '>' && rating[rule.attribute] > rule.value {
					next = rule.workflow
					break
				}
			}

			if len(next) > 0 {
				current = next
			} else {
				current = workflows[current].fallback
			}
		}

		if current == "A" {
			sum += rating['x'] + rating['m'] + rating['a'] + rating['s']
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 19114, result)
	} else {
		assert.Equal(t, 374873, result)
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
