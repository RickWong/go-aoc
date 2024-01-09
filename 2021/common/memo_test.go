package common

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMemo(t *testing.T) {
	f1 := Memo(func(a int) int { return a }, nil)
	f2 := Memo2(func(a float32, b float64) float64 { return float64(a) + b }, nil)
	f3 := Memo3(func(a string, b int, c any) any { return c }, nil)
	f4 := Memo4(func(a string, b int, c float32, d error) bool { return d.Error() != "" }, nil)

	assert.Equal(t, 1, f1(1))
	assert.Equal(t, 1, f1(1))
	assert.Equal(t, 2, f1(2))
	assert.Equal(t, 3.0, f2(1, 2))
	assert.Equal(t, 3.0, f2(1, 2))
	assert.Equal(t, 4.0, f2(1, 3))
	assert.Equal(t, t, f3("a", 1, t))
	assert.Equal(t, t, f3("a", 1, t))
	assert.Equal(t, nil, f3("a", 2, nil))
	assert.Equal(t, true, f4("a", 1, 2, errors.New("hallo")))
	assert.Equal(t, true, f4("a", 1, 2, errors.New("hallo")))
	assert.Equal(t, false, f4("a", 1, 2, errors.New("")))
}

func TestFibionacci(t *testing.T) {
	var fibionacci func(n int) int
	calls := 0
	fibionacci = Memo(func(n int) int {
		calls++
		if n < 2 {
			return n
		}
		return fibionacci(n-1) + fibionacci(n-2)
	}, &MemoOptions[int]{
		Cache: make(map[string]int),
		Hash: func(data any) string {
			return strconv.Itoa(data.(int))
		},
	})

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
}
