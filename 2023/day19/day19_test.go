package day19

import (
	_ "embed"
	"maps"
	"regexp"
	"strings"
	"testing"

	. "github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
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

type Series struct {
	Low  int
	High int
}

func (ser Series) Size() int {
	return max(0, ser.High-ser.Low+1)
}

func (ser Series) Cut(n int, operator byte) (*Series, *Series) {
	other := ser // Copy.
	switch operator {
	case '<':
		if n <= ser.Low {
			return nil, &other
		}
		if n > ser.High {
			return &ser, nil
		}
		ser.High = n - 1
		other.Low = n
	case '>':
		if n >= ser.High {
			return nil, &other
		}
		if n < ser.Low {
			return &ser, nil
		}
		ser.Low = n + 1
		other.High = n
	}
	return &ser, &other
}

type Combination map[byte]*Series

func (com Combination) Count() int {
	prod := 1
	for _, ser := range com {
		if ser == nil {
			return 0
		}
		prod *= ser.Size()
	}
	return prod
}

func (com Combination) Copy() Combination {
	other := make(Combination, len(com))
	maps.Copy(other, com)
	return other
}

func (com Combination) Cut(k byte, n int, operator byte) Combination {
	other := com.Copy()
	if com[k] != nil {
		com[k], other[k] = com[k].Cut(n, operator)
	}
	return other
}

// Helper functions.

func parseRatings(blob string) []Rating {
	ratingsRegex := regexp.MustCompile(`(?m)^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
	ratings := make([]Rating, 0)
	for _, match := range ratingsRegex.FindAllStringSubmatch(blob, -1) {
		x, m, a, s := Atoi(match[1]), Atoi(match[2]), Atoi(match[3]), Atoi(match[4])
		ratings = append(ratings, Rating{'x': x, 'm': m, 'a': a, 's': s})
	}
	return ratings
}

func parseWorkflows(blob string) map[string]Workflow {
	workflowsRegex := regexp.MustCompile(`(?m)^(\w+){((?:\w+[<>]\d+:\w+,?)+\w+)}$`)
	ruleRegex := regexp.MustCompile(`(\w+)([<>])(\d+):(\w+)`)
	workflows := make(map[string]Workflow)

	for _, match := range workflowsRegex.FindAllStringSubmatch(blob, -1) {
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
	return workflows
}

func countAccepted(com Combination, workflow Workflow, workflows map[string]Workflow) int {
	count := 0

	if com.Count() == 0 {
		return 0
	}

	for _, rule := range workflow.rules {
		remaining := com.Cut(rule.attribute, rule.value, rule.comparator)
		if rule.workflow == "A" {
			count += com.Count()
		} else {
			count += countAccepted(com, workflows[rule.workflow], workflows)
		}
		com = remaining
	}

	if workflow.fallback != "" && com.Count() > 0 {
		if workflow.fallback == "A" {
			count += com.Count()
		} else {
			count += countAccepted(com, workflows[workflow.fallback], workflows)
		}
	}

	return count
}

// Part 1.

func part1() int {
	blocks := strings.Split(strings.TrimSpace(data), "\n\n")
	workflows := parseWorkflows(blocks[0])
	ratings := parseRatings(blocks[1])

	sum := 0
	for _, rating := range ratings {
		root := Combination{
			'x': &Series{rating['x'], rating['x']},
			'm': &Series{rating['m'], rating['m']},
			'a': &Series{rating['a'], rating['a']},
			's': &Series{rating['s'], rating['s']},
		}

		if countAccepted(root, workflows["in"], workflows) > 0 {
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
	blocks := strings.Split(strings.TrimSpace(data), "\n\n")
	workflows := parseWorkflows(blocks[0])
	root := Combination{
		'x': &Series{1, 4000},
		'm': &Series{1, 4000},
		'a': &Series{1, 4000},
		's': &Series{1, 4000},
	}

	return countAccepted(root, workflows["in"], workflows)
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 167409079868000, result)
	} else {
		assert.Equal(t, 122112157518711, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func TestCut(t *testing.T) {
	t.Skip()
	t.Parallel()

	ser := Series{1, 10}
	yes, no := ser.Cut(5, '<')
	assert.Equal(t, Series{1, 4}, *yes)
	assert.Equal(t, Series{5, 10}, *no)

	ser = Series{1, 10}
	yes, no = ser.Cut(5, '>')
	assert.Equal(t, Series{6, 10}, *yes)
	assert.Equal(t, Series{1, 5}, *no)

	ser = Series{1, 10}
	yes, no = ser.Cut(1, '<')
	assert.Nil(t, yes)
	assert.Equal(t, Series{1, 10}, *no)

	ser = Series{1, 10}
	yes, no = ser.Cut(10, '>')
	assert.Nil(t, yes)
	assert.Equal(t, Series{1, 10}, *no)

	ser = Series{1, 10}
	yes, no = ser.Cut(0, '>')
	assert.Equal(t, Series{1, 10}, *yes)
	assert.Nil(t, no)

	ser = Series{1, 10}
	yes, no = ser.Cut(20, '<')
	assert.Equal(t, Series{1, 10}, *yes)
	assert.Nil(t, no)
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
