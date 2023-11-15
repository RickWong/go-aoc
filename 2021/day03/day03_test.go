package day03

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var data = input

func part1() int {
	lines := strings.Split(data, "\n")
	gamma_str := ""
	epsilon_str := ""

	for i := 0; i < len(lines[0]); i++ {
		bits := ""
		for _, line := range lines {
			bits += line[i : i+1]
		}

		if strings.Count(bits, "0") > strings.Count(bits, "1") {
			gamma_str += "0"
			epsilon_str += "1"
		} else {
			gamma_str += "1"
			epsilon_str += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gamma_str, 2, 0)
	epsilon, _ := strconv.ParseInt(epsilon_str, 2, 0)
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

func BenchmarkPart1(b *testing.B) {
	part1()
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
	o2_str := ""
	co2_str := ""

	for i := 0; len(lines) > 0 && i < len(lines[0]); i++ {
		bits := ""
		for _, line := range lines {
			bits += string(line[i])
		}

		if strings.Count(bits, "0") > strings.Count(bits, "1") {
			o2_str += "0"
			lines = filter(lines, func(line string) bool {
				return line[i] == '0'
			})
		} else {
			o2_str += "1"
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
		co2_str = lines2[0]
	}

	o2, _ := strconv.ParseInt(o2_str, 2, 0)
	co2, _ := strconv.ParseInt(co2_str, 2, 0)
	return int(o2 * co2)

	return -1
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

func BenchmarkPart2(b *testing.B) {
	part2()
}
