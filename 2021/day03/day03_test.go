package day03

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var input string

var data = input

func part1() int {
	lines := strings.Split(data, "\n")
	if lines == nil {
		panic("No data")
	}

	gammaStr := ""
	epsilonStr := ""

	for i := 0; i < len(lines[0]); i++ {
		bits := ""
		for _, line := range lines {
			bits += line[i : i+1]
		}

		if strings.Count(bits, "0") > strings.Count(bits, "1") {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 0)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 0)
	return int(gamma * epsilon)
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 198
	if data == input {
		expect = 3882564
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func part2() int {
	lines := strings.Split(strings.Trim(data, "\n\r "), "\n")
	lines2 := make([]string, len(lines))
	copy(lines2, lines)
	o2Str := ""
	co2Str := ""

	for i := 0; len(lines) > 0 && i < len(lines[0]); i++ {
		bits := ""
		for _, line := range lines {
			bits += string(line[i])
		}

		if strings.Count(bits, "0") > strings.Count(bits, "1") {
			o2Str += "0"
			lines = filter(lines, func(line string) bool {
				return line[i] == '0'
			})
		} else {
			o2Str += "1"
			lines = filter(lines, func(line string) bool {
				return line[i] == '1'
			})
		}
	}

	for i := 0; len(lines2) > 1 && i < len(lines2[0]); i++ {
		bits := ""
		for _, line := range lines2 {
			bits += string(line[i])
		}

		if strings.Count(bits, "0") <= strings.Count(bits, "1") {
			lines2 = filter(lines2, func(line string) bool {
				return line[i] == '0'
			})
		} else {
			lines2 = filter(lines2, func(line string) bool {
				return line[i] == '1'
			})
		}
	}

	if len(lines2) < 1 {
		panic("Out of lines")
	} else {
		co2Str = lines2[0]
	}

	o2, _ := strconv.ParseInt(o2Str, 2, 0)
	co2, _ := strconv.ParseInt(co2Str, 2, 0)
	return int(o2 * co2)
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 230
	if data == input {
		expect = 3385170
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}

}
