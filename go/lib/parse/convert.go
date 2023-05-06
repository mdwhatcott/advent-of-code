package parse

import (
	"strconv"
)

func Float64(value string) float64 {
	parsed, _ := strconv.ParseFloat(value, 64)
	return parsed
}

func Int(value string) int {
	parsed, _ := strconv.Atoi(value)
	return parsed
}

func Ints(values []string) (ints []int) {
	for _, value := range values {
		ints = append(ints, Int(value))
	}
	return ints
}
