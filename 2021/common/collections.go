package common

func AllValues[K, V comparable](c map[K]V, eq any) bool {
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
