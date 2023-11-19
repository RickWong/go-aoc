package day11

import (
	_ "embed"
	queue2 "github.com/zyedidia/generic/queue"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var data = input

type Octopus struct {
	energy    int
	neighbors []*Octopus
}

func part1() int {
	lines := strings.Split(data, "\n")
	octopi := parseOctopi(lines)
	flashes := 0

	for i := 0; i < 100; i++ {
		queue := queue2.New[*Octopus]()
		for _, o := range octopi {
			queue.Enqueue(o)
		}

		for !queue.Empty() {
			current := queue.Dequeue()
			current.energy++
			if current.energy == 10 {
				flashes++
				for _, o := range current.neighbors {
					queue.Enqueue(o)
				}
			}
		}

		for _, o := range octopi {
			if o.energy > 9 {
				o.energy = 0
			}
		}
	}

	return flashes
}

func parseOctopi(lines []string) []*Octopus {
	if lines == nil {
		panic("No data")
	}

	var octopi []*Octopus

	for _, line := range lines {
		for _, energyStr := range strings.Split(line, "") {
			energy, _ := strconv.Atoi(energyStr)
			octopi = append(octopi, &Octopus{energy, nil})
		}
	}

	for idx, octopus := range octopi {
		y := idx / len(lines)
		x := idx % len(lines)
		rowLength := len(lines[0])

		if y > 0 {
			octopus.neighbors = append(octopus.neighbors, octopi[(y-1)*rowLength+x])

			if x > 0 {
				octopus.neighbors = append(octopus.neighbors, octopi[(y-1)*rowLength+x-1])
			}
			if x < rowLength-1 {
				octopus.neighbors = append(octopus.neighbors, octopi[(y-1)*rowLength+x+1])
			}
		}
		if x > 0 {
			octopus.neighbors = append(octopus.neighbors, octopi[y*rowLength+x-1])
		}
		if y < rowLength-1 {
			octopus.neighbors = append(octopus.neighbors, octopi[(y+1)*rowLength+x])

			if x > 0 {
				octopus.neighbors = append(octopus.neighbors, octopi[(y+1)*rowLength+x-1])
			}
			if x < rowLength-1 {
				octopus.neighbors = append(octopus.neighbors, octopi[(y+1)*rowLength+x+1])
			}
		}
		if x < rowLength-1 {
			octopus.neighbors = append(octopus.neighbors, octopi[y*rowLength+x+1])
		}
	}
	return octopi
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 1656
	if data == input {
		expect = 1723
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	octopi := parseOctopi(lines)
	syncStep := -1

	for i := 0; i < 1000; i++ {
		queue := queue2.New[*Octopus]()
		for _, o := range octopi {
			queue.Enqueue(o)
		}

		for !queue.Empty() {
			current := queue.Dequeue()
			current.energy++
			if current.energy == 10 {
				for _, o := range current.neighbors {
					queue.Enqueue(o)
				}
			}
		}

		flashes := 0
		for _, o := range octopi {
			if o.energy > 9 {
				o.energy = 0
				flashes++
			}
		}
		if flashes == len(octopi) {
			syncStep = i + 1
			break
		}
	}

	return syncStep
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 195
	if data == input {
		expect = 327
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
