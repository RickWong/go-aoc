package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemo(t *testing.T) {
	t.Skip()

	type T1 struct {
		A int
	}

	type T2 struct {
		A [2]float32
		B float64
	}

	type T3 struct {
		A []bool
		B any
		C *string
	}

	c1 := make(map[T1]int)
	c2 := make(map[T2]float64)
	c3 := make(map[string]bool)
	f1 := Memo(func(i T1) int { return i.A }, &c1)
	f2 := Memo(func(i T2) float64 { return float64(i.A[0]) + i.B }, &c2)
	f3 := HashedMemo(func(i T3) bool { return i.C == nil }, nil, &c3)

	assert.Equal(t, 1, f1(T1{1}))
	assert.Equal(t, 1, f1(T1{1}))
	assert.Equal(t, 2, f1(T1{2}))
	assert.Len(t, c1, 2)
	assert.Equal(t, 3.0, f2(T2{[2]float32{1, 2}, 2}))
	assert.Equal(t, 3.0, f2(T2{[2]float32{1, 2}, 2}))
	assert.Equal(t, 4.0, f2(T2{[2]float32{1, 2}, 3}))
	assert.Len(t, c2, 2)
	assert.False(t, f3(T3{[]bool{true}, 1, Ptr("hallo")}))
	assert.False(t, f3(T3{[]bool{true}, 1, Ptr("hallo")}))
	assert.True(t, f3(T3{[]bool{true}, 1, nil}))
	assert.Len(t, c3, 2)
}

func TestFibionacci(t *testing.T) {
	t.Skip()

	cache := make(map[int]int)

	var fibionacci func(n int) int
	calls := 0
	fibionacci = Memo(func(n int) int {
		calls++
		if n < 2 {
			return n
		}
		return fibionacci(n-1) + fibionacci(n-2)
	}, &cache)

	assert.Equal(t, 1, fibionacci(1))
	assert.Equal(t, 1, calls)
	calls = 0
	assert.Equal(t, 1, fibionacci(2))
	assert.Equal(t, 2, calls)
	calls = 0
	assert.Equal(t, 2, fibionacci(3))
	assert.Equal(t, 1, calls)
	calls = 0
	assert.Equal(t, 3, fibionacci(4))
	assert.Equal(t, 1, calls)
	calls = 0
	assert.Equal(t, 5, fibionacci(5))
	assert.Equal(t, 1, calls)
	calls = 0
	assert.Equal(t, 832040, fibionacci(30))
	assert.Equal(t, 25, calls)
	calls = 0
	assert.Len(t, cache, 31)
}

func BenchmarkFibionacci(b *testing.B) {
	var fibionacci func(n int) int
	fibionacci = HashedMemo(func(n int) int {
		if n < 2 {
			return n
		}
		return fibionacci(n-1) + fibionacci(n-2)
	}, nil, nil)

	for n := 0; n < b.N; n++ {
		fibionacci(n)
	}
}
