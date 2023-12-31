package day16

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Beam struct {
	y, x   int
	dy, dx int
}

// Helper functions.

func Map[T, R any](collection []T, fn func(a T) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v)
	}
	return m
}

func makeLut(height int) map[int]map[int]int {
	lut := make(map[int]map[int]int)
	for y := 0; y < height; y++ {
		lut[y] = make(map[int]int)
	}
	return lut
}

func traceBeam(grid [][]string, beam Beam) map[int]map[int]int {
	beams := make([]Beam, 0, 128)
	history := make(map[Beam]bool, 1024)
	lut := makeLut(len(grid))

	beams = append(beams, beam)
	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		if _, ok := history[beam]; ok {
			continue
		}
		history[beam] = true

		steps := 0
		for {
			steps++
			y := beam.y + beam.dy*steps
			x := beam.x + beam.dx*steps
			if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
				break
			}

			tile := grid[y][x]
			lut[y][x]++

			if tile == "|" {
				beams = append(beams, Beam{y, x, -1, 0})
				beams = append(beams, Beam{y, x, 1, 0})
				break
			}
			if tile == "-" {
				beams = append(beams, Beam{y, x, 0, -1})
				beams = append(beams, Beam{y, x, 0, 1})
				break
			}
			if tile == "/" {
				if beam.dy != 0 {
					beams = append(beams, Beam{y, x, 0, -beam.dy})
				} else {
					beams = append(beams, Beam{y, x, -beam.dx, 0})
				}
				break
			}
			if tile == "\\" {
				if beam.dy != 0 {
					beams = append(beams, Beam{y, x, 0, beam.dy})
				} else {
					beams = append(beams, Beam{y, x, beam.dx, 0})
				}
				break
			}
		}
	}
	return lut
}

func printAndSum(lut map[int]map[int]int) int {
	sum := 0

	for y := range lut {
		sum += len(lut[y])
	}

	return sum
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := Map(lines, func(line string) []string { return strings.Split(line, "") })

	lut := traceBeam(grid, Beam{0, -1, 0, 1})
	sum := printAndSum(lut)

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 46, result)
	} else {
		assert.Equal(t, 7034, result)
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := Map(lines, func(line string) []string { return strings.Split(line, "") })
	sum := 0

	beams := make([]Beam, 0, 2*len(grid)+2*(len(grid[0])))
	for y := range grid {
		beams = append(beams, Beam{y, -1, 0, 1})
		beams = append(beams, Beam{y, len(grid[0]), 0, -1})
	}
	for x := range grid[0] {
		beams = append(beams, Beam{-1, x, 1, 0})
		beams = append(beams, Beam{len(grid), x, -1, 0})
	}

	eg := errgroup.Group{}
	mu := sync.Mutex{}

	for _, beam := range beams {
		beam := beam

		eg.Go(func() error {
			lut := traceBeam(grid, beam)
			mu.Lock()
			defer mu.Unlock()
			sum = max(sum, printAndSum(lut))
			return nil
		})
	}

	_ = eg.Wait()

	return sum
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 51, result)
	} else {
		assert.Equal(t, 7759, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
