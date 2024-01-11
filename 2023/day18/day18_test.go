package day18

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

var deltas = map[byte][2]int{
	'R': {0, 1},
	'L': {0, -1},
	'D': {1, 0},
	'U': {-1, 0},
}

var directions = map[byte]byte{
	'0': 'R',
	'1': 'D',
	'2': 'L',
	'3': 'U',
}

// Helper functions.

// Part 1.

func part1() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	points := make([]common.Point2D[int], 1, 128)
	points[0] = common.Point2D[int]{X: 0, Y: 0}
	boundary := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		direction, meters := parts[0][0], common.Atoi(parts[1])

		lastPoint := points[len(points)-1]
		newPoint := common.Point2D[int]{
			Y: lastPoint.Y + deltas[direction][0]*meters,
			X: lastPoint.X + deltas[direction][1]*meters,
		}
		points = append(points, newPoint)
		boundary += meters
	}

	area := int(common.Shoelace(points))
	interior := common.PicksInterior(area, boundary)

	return boundary + interior
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 62, result)
	} else {
		// 28290 is too low.
		assert.Equal(t, 47527, result)
	}
}

// Part 2.

func part2() int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	points := make([]common.Point2D[int], 1, 128)
	points[0] = common.Point2D[int]{X: 0, Y: 0}
	boundary := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		color := parts[2][2 : len(parts[2])-1]
		direction, meters := directions[color[5]], common.Hexi(color[0:5])

		lastPoint := points[len(points)-1]
		newPoint := common.Point2D[int]{
			Y: lastPoint.Y + deltas[direction][0]*meters,
			X: lastPoint.X + deltas[direction][1]*meters,
		}
		points = append(points, newPoint)
		boundary += meters
	}

	area := int(common.Shoelace(points))
	interior := common.PicksInterior(area, boundary)

	return boundary + interior
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 952408144115, result)
	} else {
		assert.Equal(t, 52240187443190, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
