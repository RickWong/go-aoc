package day08

import (
	_ "embed"
	"github.com/samber/lo"
	"math/bits"
	"slices"
	"strings"
	"testing"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var data = input

type Segment rune
type Signal uint

var segmentToSignal = map[Segment]Signal{
	'a': 0b0000001,
	'b': 0b0000010,
	'c': 0b0000100,
	'd': 0b0001000,
	'e': 0b0010000,
	'f': 0b0100000,
	'g': 0b1000000,
}

type Entry struct {
	patterns []Signal
	outputs  []Signal
}

func parseSegments(segments string, _ int) Signal {
	var signal Signal
	for _, segment := range segments {
		signal |= segmentToSignal[Segment(segment)]
	}
	return signal
}

func parseEntries(lines []string) []*Entry {
	return lo.FilterMap(lines, func(line string, _ int) (*Entry, bool) {
		signals := strings.Split(line, " | ")
		if len(signals[0]) == 0 {
			return nil, false
		}

		patternsSegments := strings.Split(signals[0], " ")
		outputSegments := strings.Split(signals[1], " ")

		entry := Entry{
			lo.Map(patternsSegments, parseSegments),
			lo.Map(outputSegments, parseSegments),
		}
		return &entry, true
	})
}

func part1() int {
	lines := strings.Split(data, "\n")
	entries := parseEntries(lines)
	obviousSignalLengths := []int{2, 3, 4, 7}
	occurrences := 0

	for _, entry := range entries {
		for _, signal := range entry.outputs {
			if slices.Contains(obviousSignalLengths, bits.OnesCount(uint(signal))) {
				occurrences++
			}
		}
	}

	return occurrences
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 26
	if data == input {
		expect = 278
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	entries := parseEntries(lines)
	values := make([]int, 0)

	for _, entry := range entries {
		decoder := make(map[int]Signal)
		decoder[8] = Signal(0b1111111)
		signals := append(entry.patterns, entry.outputs...)
		rounds := 3
		value := 0

		for i := 0; i < rounds; i++ {
			for _, signal := range signals {
				_, known := lo.FindKey(decoder, signal)
				if known {
					continue
				}

				numSegments := bits.OnesCount(uint(signal))
				switch numSegments {
				case 2:
					decoder[1] = signal
				case 3:
					decoder[7] = signal
				case 4:
					decoder[4] = signal
				case 5:
					if decoder[2] == 0 {
						if decoder[3] != 0 && decoder[3] != signal &&
							decoder[5] != 0 && decoder[5] != signal {
							decoder[2] = signal
							continue
						} else if decoder[4] != 0 && decoder[4]|signal == decoder[8] {
							decoder[2] = signal
							continue
						}
					}

					if decoder[3] == 0 {
						if decoder[2] != 0 && decoder[2] != signal &&
							decoder[5] != 0 && decoder[5] != signal {
							decoder[3] = signal
							continue
						} else if decoder[7] != 0 && bits.OnesCount(uint(signal&^decoder[7])) == 2 {
							decoder[3] = signal
							continue
						} else if decoder[1] != 0 && bits.OnesCount(uint(signal&^decoder[1])) == 3 {
							decoder[3] = signal
							continue
						}
					}

					if decoder[5] == 0 {
						if decoder[3] != 0 && decoder[3] != signal &&
							decoder[2] != 0 && decoder[2] != signal {
							decoder[5] = signal
							continue
						} else if decoder[2] != 0 && decoder[2]|signal == decoder[8] {
							decoder[5] = signal
							continue
						}
					}
				case 6:
					if decoder[0] == 0 {
						if decoder[6] != 0 && decoder[6] != signal &&
							decoder[9] != 0 && decoder[9] != signal {
							decoder[0] = signal
							continue
						} else if decoder[5] != 0 && decoder[5]|signal == decoder[8] {
							decoder[0] = signal
							continue
						}
					}

					if decoder[6] == 0 {
						if decoder[0] != 0 && decoder[0] != signal &&
							decoder[9] != 0 && decoder[9] != signal {
							decoder[6] = signal
							continue
						} else if decoder[7] != 0 && decoder[7]|signal == decoder[8] {
							decoder[6] = signal
							continue
						}
					}

					if decoder[9] == 0 {
						if decoder[6] != 0 && decoder[6] != signal &&
							decoder[0] != 0 && decoder[0] != signal {
							decoder[9] = signal
							continue
						} else if decoder[4] != 0 && decoder[4]|signal == signal {
							decoder[9] = signal
							continue
						}
					}
				case 7:
					decoder[8] = signal
				default:
					panic("Invalid signal")
				}
			}

			if len(decoder) == 10 {
				digit1, _ := lo.FindKey(decoder, entry.outputs[0])
				digit2, _ := lo.FindKey(decoder, entry.outputs[1])
				digit3, _ := lo.FindKey(decoder, entry.outputs[2])
				digit4, _ := lo.FindKey(decoder, entry.outputs[3])
				value = 1000*digit1 + 100*digit2 + 10*digit3 + digit4
				break
			}
		}

		if len(decoder) != 10 {
			panic("Decoding error")
		}

		values = append(values, value)
	}

	return lo.Sum(values)
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 61229
	if data == input {
		expect = 986179
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
