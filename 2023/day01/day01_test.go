package day01

import (
	_ "embed"
	common2 "github.com/RickWong/go-aoc/common"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

func calibrate(s string) int {
	first := strings.IndexAny(s, "0123456789")
	last := strings.LastIndexAny(s, "0123456789")
	if first < 0 || last < 0 {
		return -1
	}

	firstDigit := s[first] - '0'
	lastDigit := s[last] - '0'
	return int(firstDigit*10 + lastDigit)
}

func part1() int {
	lines := strings.Split(data, "\n")
	return common2.Sum(common2.Map(lines, calibrate))
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 208
	if data == Input {
		expect = 53921
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func translate(s string) string {
	replacements := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for i := 0; i < len(s); i++ {
		for needle, replacement := range replacements {
			j := i + len(needle)
			if j <= len(s) && s[i:j] == needle {
				s = s[:i] + replacement + s[j:]
				break
			}
		}
	}

	return s
}

func part2() int {
	lines := strings.Split(data, "\n")
	return common2.Sum(common2.Map(lines, func(s string) int {
		return calibrate(translate(s))
	}))
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 281
	if data == Input {
		expect = 54676
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
