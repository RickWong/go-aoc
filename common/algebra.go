package common

import (
	"golang.org/x/exp/constraints"
	"math"
)

// LCM returns the least common multiple of the integers.
func LCM(integers ...int) int {
	if len(integers) < 1 {
		return 0
	}

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
	X, Y N
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

func Octile[N Number](x1 N, y1 N, x2 N, y2 N) N {
	diffY := y2 - y1
	diffX := x2 - x1
	if diffY < 0 {
		diffY = -diffY
	}
	if diffX < 0 {
		diffX = -diffX
	}
	if diffY > diffX {
		return diffY + N(float64(diffX)*0.41421356237)
	}
	return diffX + N(float64(diffY)*0.41421356237)
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

func Shoelace[N Number](points []Point2D[N]) float64 {
	P := len(points)
	sum := N(0)
	for p := 0; p < len(points); p++ {
		p1 := points[p]
		p2 := points[(p+1)%P]
		sum += p1.X*p2.Y - p2.X*p1.Y
	}
	return math.Abs(float64(sum)) / 2
}

func PicksInterior[N Number](area N, boundary N) N {
	return area - boundary/2 + 1
}

func EuclideanMod[N constraints.Integer](n N, div N) N {
	remainder := ((n % div) + div) % div
	return remainder
}

func FindQuadraticCoefficients[N Number](y0, y1, y2 N) (N, N, N) {
	a := (y2 - (2 * y1) + y0) / 2
	b := y1 - y0 - a
	c := y0
	return a, b, c
}

func SolveQuadratic[N Number](a, b, c, x N) N {
	y := a*x*x + b*x + c
	return y
}
