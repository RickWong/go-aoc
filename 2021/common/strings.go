package common

import "strings"

// In is faster than strings.ContainsAny().
func In(haystack string, chars string) bool {
	for _, c := range chars {
		if strings.IndexByte(haystack, byte(c)) >= 0 {
			return true
		}
	}
	return false
}
