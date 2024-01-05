package common

import (
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

// In is faster than strings.ContainsAny().
func In(haystack string, chars string) bool {
	for _, c := range chars {
		if strings.IndexByte(haystack, byte(c)) >= 0 {
			return true
		}
	}
	return false
}

func Join[I constraints.Integer, F constraints.Float](collection any, separator string) string {
	switch collection.(type) {
	case []string:
		ss := collection.([]string)
		return strings.Join(ss, separator)
	case []I:
		ss := Map(collection.([]int), strconv.Itoa)
		return strings.Join(ss, separator)
	case []F:
		ss := Map(collection.([]float64),
			func(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) },
		)
		return strings.Join(ss, separator)
	default:
		panic("unsupported type")
	}
}
