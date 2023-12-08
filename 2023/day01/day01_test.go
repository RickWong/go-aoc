package day01

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

func mapFn[T, R any](collection []T, fn func(a T) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v)
	}
	return m
}

func sum(collection []int) int {
	m := 0
	for _, v := range collection {
		m += v
	}
	return m
}

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
	return sum(mapFn(lines, calibrate))
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 208
	if data == Input {
		expect = 55208
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

func transform(s string) int {
	return calibrate(translate(s))
}

func part2() int {
	lines := strings.Split(data, "\n")
	return sum(mapFn(lines, transform))
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 281
	if data == Input {
		expect = 54578
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
