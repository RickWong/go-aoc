package day03

import (
	_ "embed"
	"strings"
	"testing"
	"unicode"

	"github.com/RickWong/go-aoc/common"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Number struct {
	y,
	x1, x2 int
	value int
}

func irange(start int, end int) []int {
	arr := make([]int, 0, end-start)
	for i := start; i <= end; i++ {
		arr = append(arr, i)
	}
	return arr
}

func part1() int {
	lines := strings.Split(data, "\n")
	sum := 0

	// Researched how to do this with multiline regexp.
	// Below will return all numbers and their start & end indices,
	// then those indices would need to be divided & modulo-ed by the line width.
	//
	//   r := regexp.MustCompile(`(?m)\d+`)
	//   m := r.FindAllStringIndex(multi, -1)

	numbers := make([]Number, 0, 100)
	for y, line := range lines {
		parsed := ""
		x1 := 0

		for i := 0; i < len(line); i++ {
			isDigit := unicode.IsDigit(rune(line[i]))
			isLast := i == len(line)-1

			if isDigit {
				if parsed == "" {
					x1 = i
				}
				parsed += line[i : i+1]
			}

			if len(parsed) > 0 &&
				(!isDigit || isLast) {
				numbers = append(numbers, Number{y, x1, i - 1, common.Atoi(parsed)})
				parsed = ""
			}
		}
	}

	for _, number := range numbers {
		yRange := irange(number.y-1, number.y+1)
		xRange := irange(number.x1-1, number.x2+1)
		symbolFound := false
	scan:
		for _, y := range yRange {
			if y < 0 || y >= len(lines) {
				continue
			}

			for _, x := range xRange {
				if x < 0 || x >= len(lines[0]) {
					continue
				}

				r := rune(lines[y][x])

				if !unicode.IsDigit(r) && r != '.' {
					symbolFound = true
					break scan
				}
			}
		}
		if symbolFound {
			sum += number.value
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()
	expect := 4361
	if data == Input {
		expect = 559667
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	sum := 0

	numbers := make([]Number, 0, 100)
	for y, line := range lines {
		parsed := ""
		x1 := 0

		for i := 0; i < len(line); i++ {
			isDigit := unicode.IsDigit(rune(line[i]))
			isLast := i == len(line)-1

			if isDigit {
				if parsed == "" {
					x1 = i
				}
				parsed += line[i : i+1]
			}

			if len(parsed) > 0 &&
				(!isDigit || isLast) {
				numbers = append(numbers, Number{y, x1, i - 1, common.Atoi(parsed)})
				parsed = ""
			}
		}
	}

	gears := make(map[int][]Number)

	for _, number := range numbers {
		yRange := irange(number.y-1, number.y+1)
		xRange := irange(number.x1-1, number.x2+1)
	scan:
		for _, y := range yRange {
			if y < 0 || y >= len(lines) {
				continue
			}

			for _, x := range xRange {
				if x < 0 || x >= len(lines[0]) {
					continue
				}

				r := rune(lines[y][x])

				if r == '*' {
					gears[y*10000000+x] = append(gears[y*10000000+x], number)
					break scan
				}
			}
		}
	}

	for _, parts := range gears {
		if len(parts) == 2 {
			sum += parts[0].value * parts[1].value
		}
	}

	return sum
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()
	expect := 467835
	if data == Input {
		expect = 86841457
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
