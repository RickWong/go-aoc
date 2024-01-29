package day16

import (
	_ "embed"
	"github.com/RickWong/go-aoc/common"
	"github.com/edwingeng/deque/v2"
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

type LUT []int

// Helper functions.

func traceBeam(grid [][]byte, beam Beam) LUT {
	history := make(map[Beam]bool, 1024)
	height := len(grid)
	width := len(grid[0])
	lut := make(LUT, height*width)

	beams := deque.NewDeque[Beam](deque.WithChunkSize(1024))
	beams.PushBack(beam)
	for !beams.IsEmpty() {
		beam := beams.PopBack()

		if _, ok := history[beam]; ok {
			continue
		}
		history[beam] = true

		for steps := 1; ; steps++ {
			y := beam.y + beam.dy*steps
			x := beam.x + beam.dx*steps
			if y < 0 || y >= height || x < 0 || x >= width {
				break
			}

			lut[y*width+x] = 1

			tile := grid[y][x]
			if tile == '|' {
				beams.PushBack(Beam{y, x, -1, 0})
				beams.PushBack(Beam{y, x, 1, 0})
				break
			}
			if tile == '-' {
				beams.PushBack(Beam{y, x, 0, -1})
				beams.PushBack(Beam{y, x, 0, 1})
				break
			}
			if tile == '/' {
				if beam.dy != 0 {
					beams.PushBack(Beam{y, x, 0, -beam.dy})
				} else {
					beams.PushBack(Beam{y, x, -beam.dx, 0})
				}
				break
			}
			if tile == '\\' {
				if beam.dy != 0 {
					beams.PushBack(Beam{y, x, 0, beam.dy})
				} else {
					beams.PushBack(Beam{y, x, beam.dx, 0})
				}
				break
			}
		}
	}
	return lut
}

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	grid := common.Map(lines, func(line string) []byte { return []byte(line) })

	lut := traceBeam(grid, Beam{0, -1, 0, 1})
	sum := common.Sum(lut)

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
	grid := common.Map(lines, func(line string) []byte { return []byte(line) })

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
	sum := 0

	for _, beam := range beams {
		beam := beam

		eg.Go(func() error {
			lut := traceBeam(grid, beam)
			mu.Lock()
			defer mu.Unlock()
			sum = max(sum, common.Sum(lut))
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
