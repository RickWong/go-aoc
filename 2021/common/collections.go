package common

func AllValues[K Hashable, V comparable](c map[K]V, eq any) bool {
	for _, v := range c {
		if v != eq {
			return false
		}
	}
	return true
}

// Reset resets the slice to be empty but keeps the capacity.
func Reset[T any](c *[]T) {
	capacity := cap(*c)
	*c = (*c)[:0:capacity]
}

func Map[T, R any](collection []T, fn func(a T) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v)
	}
	return m
}

func Map2[T, R any](collection []T, fn func(a T, i int) R) []R {
	m := make([]R, len(collection))
	for i, v := range collection {
		m[i] = fn(v, i)
	}
	return m
}
