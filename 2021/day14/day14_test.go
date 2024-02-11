package day14

import (
	_ "embed"
	"github.com/samber/lo"
	"maps"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

func part1() int {
	sections := strings.Split(data, "\n\n")
	template := sections[0]
	rules := parseRules(sections)

	polymer := template
	for i := 0; i < 10; i++ {
		next := ""
		for idx := 0; idx < len(polymer)-1; idx++ {
			key := polymer[idx : idx+2]
			if rules[key] != "" {
				next += polymer[idx:idx+1] + rules[key]
			} else {
				next += key
			}
		}
		polymer = next + polymer[len(polymer)-1:]
	}

	counts := lo.Values(lo.CountValues(strings.Split(polymer, "")))
	return lo.Max(counts) - lo.Min(counts)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 1588
	if data == Input {
		expect = 2321
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	sections := strings.Split(data, "\n\n")
	template := sections[0]
	rules := parseRules(sections)

	pairCounter := make(map[string]int)
	elementCounter := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		pairCounter[pair]++
	}

	for i := 0; i < len(template); i++ {
		element := template[i : i+1]
		elementCounter[element]++
	}

	for i := 0; i < 40; i++ {
		snapshot := make(map[string]int, len(pairCounter))
		maps.Copy(snapshot, pairCounter)

		for pair, count := range snapshot {
			if count == 0 {
				continue
			}

			inserted := rules[pair]
			elementCounter[inserted] += count

			pairCounter[pair] -= count
			pairCounter[pair[0:1]+inserted] += count
			pairCounter[inserted+pair[1:2]] += count
		}
	}

	counts := lo.Values(elementCounter)
	return lo.Max(counts) - lo.Min(counts)
}

func parseRules(sections []string) map[string]string {
	instructions := strings.Split(sections[1], "\n")
	rules := make(map[string]string)
	for _, instruction := range instructions {
		fromTo := strings.Split(instruction, " -> ")
		rules[fromTo[0]] = fromTo[1]
	}
	return rules
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 2188189693529
	if data == Input {
		expect = 2399822193707
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
