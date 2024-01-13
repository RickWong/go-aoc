package day20

import (
	_ "embed"
	. "github.com/RickWong/go-aoc/common"
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

	low := 0
	high := 0

	pulses := make([]Pulse, 0, 768)
	for i := 0; i < 1000; i++ {
		pulses = pulses[:0]
		pulses = append(pulses, Pulse{"button", "broadcaster", false})

		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]

			if pulse.power {
				high++
			} else {
				low++
			}

			module := modules[pulse.output]
			mem := memory[module.name]

			if module.flipflop {
				if pulse.power {
					continue
				}
				power := !mem[pulse.input]
				mem[pulse.input] = power
				pulse.power = power
			} else if module.conjunction {
				mem[pulse.input] = pulse.power
				pulse.power = !AllValues(mem, true)
			}

			for _, output := range module.outputs {
				pulses = append(pulses, Pulse{module.name, output, pulse.power})
			}
		}
	}

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
	data := data
	modulesRegex := regexp.MustCompile(`(?m)^([\w%&]+) -> ([\w, ]+)$`)
	modules := make(map[string]Module, 1024)
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

	// Inspected input to find conjunctions that lead to kl and rx.
	modulesInLCM := map[string]bool{"fp": true, "mk": true, "xt": true, "zc": true}
	lcm := NewLCMCalculator[string](len(modulesInLCM))

	pulses := make([]Pulse, 0, 1024*256)
	for i := 0; i < 10000; i++ {
		pulses = pulses[:0]
		pulses = append(pulses, Pulse{"button", "broadcaster", false})

		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]

			module := modules[pulse.output]
			mem := memory[module.name]

			if module.flipflop {
				if pulse.power {
					continue
				}
				power := !mem[pulse.input]
				mem[pulse.input] = power
				pulse.power = power
			} else if module.conjunction {
				mem[pulse.input] = pulse.power
				pulse.power = !AllValues(mem, true)

				if pulse.power && modulesInLCM[module.name] && lcm.Detect(module.name, i+1) {
					return lcm.Calc()
				}
			}

			for _, output := range module.outputs {
				pulses = append(pulses, Pulse{module.name, output, pulse.power})
			}
		}
	}

	return 0
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result)
	} else {
		assert.Equal(t, 232774988886497, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
