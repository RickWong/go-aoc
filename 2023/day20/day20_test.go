package day20

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/2021/common"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Module struct {
	name        string
	outputs     []string
	flipflop    bool
	conjunction bool
}

type Pulse struct {
	input  string
	output string
	power  bool
}

// Helper functions.

// Part 1.

func part1() int {
	data := data
	modulesRegex := regexp.MustCompile(`(?m)^([\w%&]+) -> ([\w, ]+)$`)
	modules := make(map[string]Module)
	for _, match := range modulesRegex.FindAllStringSubmatch(data, -1) {
		input, outputs := match[1], strings.Split(match[2], ", ")
		flipflop := strings.HasPrefix(input, "%")
		conjunction := strings.HasPrefix(input, "&")
		if flipflop || conjunction {
			input = input[1:]
		}
		modules[input] = Module{input, outputs, flipflop, conjunction}
	}

	// For each output, store false values for each input.
	memory := make(map[string]map[string]bool)
	for _, module := range modules {
		for _, output := range module.outputs {
			if memory[output] == nil {
				memory[output] = make(map[string]bool)
			}
			memory[output][module.name] = false
		}
	}

	history := make([]Pulse, 0, 1024)

	for i := 0; i < 1000; i++ {
		pulses := make([]Pulse, 1, 64)
		pulses[0] = Pulse{"button", "broadcaster", false}

		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]
			history = append(history, pulse)

			module := modules[pulse.output]
			if module.flipflop {
				if pulse.power {
					continue
				}
				memory[module.name][pulse.input] = !memory[module.name][pulse.input]
				pulse.power = memory[module.name][pulse.input]
			} else if module.conjunction {
				memory[module.name][pulse.input] = pulse.power
				pulse.power = !AllValues(memory[module.name], true)
			}

			for _, output := range module.outputs {
				pulses = append(pulses, Pulse{module.name, output, pulse.power})
			}
		}
	}

	low := lo.CountBy(history, func(p Pulse) bool { return !p.power })
	high := lo.CountBy(history, func(p Pulse) bool { return p.power })
	sum := low * high
	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 11687500, result)
	} else {
		assert.Equal(t, 980457412, result)
	}
}

// Part 2.

func part2() int {
	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result)
	} else {
		assert.Equal(t, 357134560737, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
