package day01

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

func calibrate(s *string) int {
	firstIndex := strings.IndexAny(*s, "0123456789")
	lastIndex := strings.LastIndexAny(*s, "0123456789")
	if firstIndex < 0 || lastIndex < 0 {
		return -1
	}

	firstDigit, _ := strconv.Atoi((*s)[firstIndex : firstIndex+1])
	lastDigit, _ := strconv.Atoi((*s)[lastIndex : lastIndex+1])

	return firstDigit*10 + lastDigit
}

func part1() int {
	lines := strings.Split(data, "\n")
	sum := 0
	for _, line := range lines {
		sum += calibrate(&line)
	}
	return sum
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

func translate(s *string) {
	// Can't use strings.Replace (need to keep string order) nor
	// strings.NewReplacer (cannot replace in reverse order "oneight").
	// Let's write our own translator.
	replacements := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

first:
	for i := 0; i < len(*s); i++ {
		for needle, replacement := range replacements {
			nLen := uint(len(needle))
			if lo.Substring(*s, i, nLen) == needle {
				rLen := len(replacement)
				*s = (*s)[0:i] + replacement + (*s)[i+rLen:]
				break first
			}
		}
	}
last:
	for i := len(*s) - 1; i >= 0; i-- {
		for needle, replacement := range replacements {
			nLen := uint(len(needle))
			if lo.Substring(*s, i, nLen) == needle {
				rLen := len(replacement)
				*s = (*s)[0:i] + replacement + (*s)[i+rLen:]
				break last
			}
		}
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	sum := 0
	for _, line := range lines {
		translate(&line)
		sum += calibrate(&line)
	}
	return sum
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
