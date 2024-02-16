package day24

import (
	_ "embed"
	"math"
	"regexp"
	"strings"
	"testing"

	"github.com/RickWong/go-aoc/common"
	"github.com/stretchr/testify/assert"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

// Data types.

type Vec2 struct {
	X, Y float64
}

type Vec3 struct {
	X, Y, Z float64
}

type Hailstone struct {
	P Vec3
	V Vec3
}

// Helper functions.

func parse() []Hailstone {
	hailstones := make([]Hailstone, 0, strings.Count(data, "\n")+1)

	re := regexp.MustCompile(`(?m)^(\d+), +(\d+), +(\d+) +@ +(-?\d+), +(-?\d+), +(-?\d+)`)
	for _, match := range re.FindAllStringSubmatch(data, -1) {
		px, py, pz := common.Atof(match[1]), common.Atof(match[2]), common.Atof(match[3])
		vx, vy, vz := common.Atof(match[4]), common.Atof(match[5]), common.Atof(match[6])
		hailstones = append(hailstones, Hailstone{Vec3{px, py, pz}, Vec3{vx, vy, vz}})
	}

	return hailstones
}

func CollisionPointInFuture(h1, h2 Hailstone) (Vec2, bool) {
	x1 := h1.P.X - 0*h1.V.X
	y1 := h1.P.Y - 0*h1.V.Y
	x2 := h1.P.X + 1*h1.V.X
	y2 := h1.P.Y + 1*h1.V.Y
	x3 := h2.P.X - 0*h2.V.X
	y3 := h2.P.Y - 0*h2.V.Y
	x4 := h2.P.X + 1*h2.V.X
	y4 := h2.P.Y + 1*h2.V.Y

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

func CollisionTime(h1, h2, h3 Hailstone) float64 {
	x1, y1, z1, vx1, vy1, vz1 := h1.P.X, h1.P.Y, h1.P.Z, h1.V.X, h1.V.Y, h1.V.Z
	x2, y2, z2, vx2, vy2, vz2 := h2.P.X, h2.P.Y, h2.P.Z, h2.V.X, h2.V.Y, h2.V.Z
	x3, y3, z3, vx3, vy3, vz3 := h3.P.X, h3.P.Y, h3.P.Z, h3.V.X, h3.V.Y, h3.V.Z

	yz := y1*(z2-z3) + y2*(-z1+z3) + y3*(z1-z2)
	xz := x1*(-z2+z3) + x2*(z1-z3) + x3*(-z1+z2)
	xy := x1*(y2-y3) + x2*(-y1+y3) + x3*(y1-y2)

	vxvy := vx1*(vy2-vy3) + vx2*(-vy1+vy3) + vx3*(vy1-vy2)
	vxvz := vx1*(-vz2+vz3) + vx2*(vz1-vz3) + vx3*(-vz1+vz2)
	vyvz := vy1*(vz2-vz3) + vy2*(-vz1+vz3) + vy3*(vz1-vz2)

	nominator := (vx2-vx3)*yz + (vy2-vy3)*xz + (vz2-vz3)*xy
	denominator := (z2-z3)*vxvy + (y2-y3)*vxvz + (x2-x3)*vyvz

	return nominator / denominator
}

// Part 1.

func part1() int {
	hailstones := parse()

	start := float64(200000000000000)
	end := float64(400000000000000)
	sum := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			a := hailstones[i]
			b := hailstones[j]
			collision, collides := CollisionPointInFuture(a, b)
			if !collides {
				continue
			}
			if collision.X < start || collision.X > end ||
				collision.Y < start || collision.Y > end {
				continue
			}
			if a.V.X > 0 && collision.X < a.P.X {
				continue
			}
			if b.V.X > 0 && collision.X < b.P.X {
				continue
			}
			if a.V.X < 0 && collision.X > a.P.X {
				continue
			}
			if b.V.X < 0 && collision.X > b.P.X {
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
	hailstones := parse()

	t1 := CollisionTime(hailstones[0], hailstones[1], hailstones[2])
	t2 := CollisionTime(hailstones[1], hailstones[0], hailstones[2])

	p1, p2 := hailstones[0].P, hailstones[1].P
	v1, v2 := hailstones[0].V, hailstones[1].V

	c1 := Vec3{p1.X + t1*v1.X, p1.Y + t1*v1.Y, p1.Z + t1*v1.Z}
	c2 := Vec3{p2.X + t2*v2.X, p2.Y + t2*v2.Y, p2.Z + t2*v2.Z}
	v0 := Vec3{(c2.X - c1.X) / (t2 - t1), (c2.Y - c1.Y) / (t2 - t1), (c2.Z - c1.Z) / (t2 - t1)}

	rock := Hailstone{
		Vec3{p1.X + (v1.X*t1 - v0.X*t1), p1.Y + (v1.Y*t1 - v0.Y*t1), p1.Z + (v1.Z*t1 - v0.Z*t1)},
		v0,
	}

	return int(math.Round(rock.P.X + rock.P.Y + rock.P.Z))
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result := part2()

	if data == Example {
		assert.Equal(t, 82000210, result)
	} else {
		assert.Equal(t, 1007148211789625, result)
	}
}

// Kaizen. Kaizen. Kaizen.

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}
