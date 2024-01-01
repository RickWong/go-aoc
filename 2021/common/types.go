package common

import "strconv"

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Atof(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func NonZero(i int) int {
	if i != 0 {
		return 1
	}
	return 0
}

func Ptr[T any](v T) *T {
	return &v
}
