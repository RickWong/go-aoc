package day24

import (
	_ "embed"
	"github.com/RickWong/go-aoc/2021/common"
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

type Vec2 struct {
	x, y float64
}

type Vec3 struct {
	x, y, z float64
}

type Hailstone struct {
	pos Vec3
	vel Vec3
}

// Helper functions.

func CollisionPointInFuture(h1, h2 Hailstone) (Vec2, bool) {
	x1 := h1.pos.x - 0*h1.vel.x
	y1 := h1.pos.y - 0*h1.vel.y
	x2 := h1.pos.x + 1*h1.vel.x
	y2 := h1.pos.y + 1*h1.vel.y
	x3 := h2.pos.x - 0*h2.vel.x
	y3 := h2.pos.y - 0*h2.vel.y
	x4 := h2.pos.x + 1*h2.vel.x
	y4 := h2.pos.y + 1*h2.vel.y

	nominator := (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
	denominator := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	if denominator == 0 {
		return Vec2{}, false
	}

	time := nominator / denominator
	x := x1 + time*(x2-x1)
	y := y1 + time*(y2-y1)

	return Vec2{x, y}, true
}

// Part 1.

func part1() int {
	hailstones := make([]Hailstone, 0, strings.Count(data, "\n")+1)

	re := regexp.MustCompile(`(?m)^(\d+), +(\d+), +(\d+) +@ +(-?\d+), +(-?\d+), +(-?\d+)`)
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		px, py, pz := common.Atof(match[1]), common.Atof(match[2]), common.Atof(match[3])
		vx, vy, vz := common.Atof(match[4]), common.Atof(match[5]), common.Atof(match[6])
		hailstones = append(hailstones, Hailstone{Vec3{px, py, pz}, Vec3{vx, vy, vz}})
	}

	field := []float64{200000000000000, 400000000000000}
	sum := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			a := hailstones[i]
			b := hailstones[j]
			collision, collides := CollisionPointInFuture(a, b)
			if !collides {
				continue
			}
			if collision.x < field[0] || collision.x > field[1] ||
				collision.y < field[0] || collision.y > field[1] {
				continue
			}
			if a.vel.x > 0 && collision.x < a.pos.x {
				continue
			}
			if b.vel.x > 0 && collision.x < b.pos.x {
				continue
			}
			if a.vel.x < 0 && collision.x > a.pos.x {
				continue
			}
			if b.vel.x < 0 && collision.x > b.pos.x {
				continue
			}
			sum++
		}
	}

	return sum
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result := part1()

	if data == Example {
		assert.Equal(t, 2, result)
	} else {
		assert.Equal(t, 14799, result)
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
