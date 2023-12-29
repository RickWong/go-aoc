package common

func AllValues[K, V comparable](c map[K]V, eq any) bool {
	for _, v := range c {
		if v != eq {
			return false
		}
	}
	return true
}
