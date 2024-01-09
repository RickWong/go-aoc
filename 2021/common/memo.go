package common

import (
	"fmt"
)

type MemoOptions[Out any] struct {
	Cache map[string]Out
	Hash  func(data any) string
}

// Memo is a generic memoization function that accepts 1 argument.
func Memo[A, Out any](f func(a A) Out, opts *MemoOptions[Out]) func(a A) Out {
	if opts == nil {
		opts = &MemoOptions[Out]{}
	}
	if opts.Cache == nil {
		opts.Cache = make(map[string]Out)
	}
	return func(a A) Out {
		var key string
		if opts.Hash != nil {
			key = opts.Hash(a)
		} else {
			key = fmt.Sprint(a)
		}
		if v, ok := opts.Cache[key]; ok {
			return v
		}
		v := f(a)
		opts.Cache[key] = v
		return v
	}
}

// Memo2 is a generic memoization function that accepts 2 arguments.
func Memo2[A, B, Out any](f func(a A, b B) Out, opts *MemoOptions[Out]) func(a A, b B) Out {
	memoF := Memo(func(in []any) Out {
		a, _ := in[0].(A)
		b, _ := in[1].(B)
		return f(a, b)
	}, opts)
	return func(a A, b B) Out {
		return memoF([]any{a, b})
	}
}

// Memo3 is a generic memoization function that accepts 3 arguments.
func Memo3[A, B, C, Out any](f func(a A, b B, c C) Out, opts *MemoOptions[Out]) func(a A, b B, c C) Out {
	memoF := Memo(func(in []any) Out {
		a, _ := in[0].(A)
		b, _ := in[1].(B)
		c, _ := in[2].(C)
		return f(a, b, c)
	}, opts)
	return func(a A, b B, c C) Out {
		return memoF([]any{a, b, c})
	}
}

// Memo4 is a generic memoization function that accepts 4 arguments.
func Memo4[A, B, C, D, Out any](f func(a A, b B, c C, d D) Out, opts *MemoOptions[Out]) func(a A, b B, c C, d D) Out {
	memoF := Memo(func(in []any) Out {
		a, _ := in[0].(A)
		b, _ := in[1].(B)
		c, _ := in[2].(C)
		d, _ := in[3].(D)
		return f(a, b, c, d)
	}, opts)
	return func(a A, b B, c C, d D) Out {
		return memoF([]any{a, b, c, d})
	}
}
