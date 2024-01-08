package common

import (
	"fmt"
)

// Memo is a generic memoization function that accepts 1 argument.
func Memo[Input comparable, Output any](f func(a Input) Output) func(a Input) Output {
	cache := make(map[string]Output)
	return func(a Input) Output {
		key := fmt.Sprintf("%v", a)
		if v, ok := cache[key]; ok {
			return v
		}

		v := f(a)
		cache[key] = v
		return v
	}
}

// Memo2 is a generic memoization function that accepts 2 arguments.
func Memo2[Input1, Input2 comparable, Output any](f func(a Input1, b Input2) Output) func(a Input1, b Input2) Output {
	cache := make(map[string]Output)
	return func(a Input1, b Input2) Output {
		key := fmt.Sprintf("%v", []any{a, b})
		if v, ok := cache[key]; ok {
			return v
		}

		v := f(a, b)
		cache[key] = v
		return v
	}
}

// Memo3 is a generic memoization function that accepts 3 arguments.
func Memo3[Input1, Input2, Input3 comparable, Output any](f func(a Input1, b Input2, c Input3) Output) func(a Input1, b Input2, c Input3) Output {
	cache := make(map[string]Output)
	return func(a Input1, b Input2, c Input3) Output {
		key := fmt.Sprintf("%v", []any{a, b, c})
		if v, ok := cache[key]; ok {
			return v
		}

		v := f(a, b, c)
		cache[key] = v
		return v
	}
}

// Memo4 is a generic memoization function that accepts 4 arguments.
func Memo4[Input1, Input2, Input3, Input4 comparable, Output any](f func(a Input1, b Input2, c Input3, d Input4) Output) func(a Input1, b Input2, c Input3, d Input4) Output {
	cache := make(map[string]Output)
	return func(a Input1, b Input2, c Input3, d Input4) Output {
		key := fmt.Sprintf("%v", []any{a, b, c, d})
		if v, ok := cache[key]; ok {
			return v
		}

		v := f(a, b, c, d)
		cache[key] = v
		return v
	}
}
