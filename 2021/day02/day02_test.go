package day02

import (
	_ "embed"
	"fmt"
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
	x := 0
	y := 0

	for _, line := range lines {
		var direction string
		var steps int
		_, err := fmt.Sscanf(line, "%s %d", &direction, &steps)
		if err != nil {
			return -1
		}

		switch direction {
		case "forward":
			x += steps
		case "down":
			y += steps
		case "up":
			y -= steps
		}
	}
	return x * y
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 15 * 10
	if data == input {
		expect = 2065 * 917 // 1_893_605
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	x := 0
	y := 0
	aim := 0

	for _, line := range lines {
		var direction string
		var steps int
		_, err := fmt.Sscanf(line, "%s %d", &direction, &steps)
		if err != nil {
			return -1
		}

		switch direction {
		case "forward":
			x += steps
			y += steps * aim
		case "down":
			aim += steps
		case "up":
			aim -= steps
		}
	}
	return x * y
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 15 * 60
	if data == input {
		expect = 2_120_734_350
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}

}
