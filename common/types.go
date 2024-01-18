package common

import "strconv"

type Hashable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Atof(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func Hexi(s string) int {
	v, _ := strconv.ParseInt(s, 16, 64)
	return int(v)
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

func Deref[T any](v *T) T {
	return *v
}
