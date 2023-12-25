package common

import "strconv"

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Ptr[T any](v T) *T {
	return &v
}
