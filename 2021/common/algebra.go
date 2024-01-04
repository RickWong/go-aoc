package common

import (
	"golang.org/x/exp/constraints"
	"math"
)

// LCM returns the least common multiple of the integers.
func LCM(integers ...int) int {
	if len(integers) < 2 {
		return integers[0]
	}

	a, b := integers[0], integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// GCD returns the greatest common divisor of the integers.
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type Point2D[N Number] struct {
	x, y N
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[N Number](collection []N) N {
	var m N = 0
	for _, v := range collection {
		m += v
	}
	return m
}

func Product[N Number](collection []N) N {
	var m N = 1
	for _, v := range collection {
		m *= v
	}
	return m
}

func Manhattan[N Number](x1 N, y1 N, x2 N, y2 N) N {
	diffY := y2 - y1
	diffX := x2 - x1
	if diffY < 0 {
		diffY = -diffY
	}
	if diffX < 0 {
		diffX = -diffX
	}
	return diffY + diffX
}

func Pythagoras[N Number](x1 N, y1 N, x2 N, y2 N) N {
	diffY := y2 - y1
	diffX := x2 - x1
	return Sqrt[N](diffY*diffY + diffX*diffX)
}

func Sqrt[N Number](x N) (z N) {
	z = 100.0
	step := func() N {
		return z - (z*z-x)/(2*z)
	}
	for zz := step(); math.Abs(float64(zz-z)) > 0.00001; {
		z = zz
		zz = step()
	}
	return
}
