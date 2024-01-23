package common

import (
	jsoniter "github.com/json-iterator/go"
	"hash/fnv"
	"strconv"
)

// Memo is a generic memoization function that accepts 1 comparable argument.
func Memo[In comparable, Out any](
	f func(a In) Out,
	cache *map[In]Out,
) func(a In) Out {
	if cache == nil {
		cache = Ptr(make(map[In]Out))
	}

	return func(a In) Out {
		if v, ok := (*cache)[a]; ok {
			return v
		}
		v := f(a)
		(*cache)[a] = v
		return v
	}
}

// HashedMemo is a generic memoization function that hashes an incomparable argument.
func HashedMemo[In any, Out any](
	f func(a In) Out,
	hashFn *func(data any) string,
	cache *map[string]Out,
) func(a In) Out {
	if cache == nil {
		cache = Ptr(make(map[string]Out))
	}

	if hashFn == nil {
		hashFn = Ptr(FasterHash)
	}

	return func(a In) Out {
		k := (*hashFn)(a)
		if v, ok := (*cache)[k]; ok {
			return v
		}

		v := f(a)
		(*cache)[k] = v
		return v
	}
}

// FasterHash takes any object and returns a hash string as fast as possible.
func FasterHash(data any) string {
	bytes, _ := jsoniter.ConfigDefault.Marshal(data)
	hash := fnv.New64a()
	_, _ = hash.Write(bytes)
	res := strconv.Itoa(int(hash.Sum64()))
	return res
}
