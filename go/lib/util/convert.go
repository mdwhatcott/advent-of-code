package util

import (
	"strconv"
)

func ParseFloat(value string) float64 {
	parsed, _ := strconv.ParseFloat(value, 64)
	return parsed
}

func ParseInt(value string) int {
	parsed, _ := strconv.Atoi(value)
	return parsed
}

func ParseInts(values []string) (ints []int) {
	for _, value := range values {
		ints = append(ints, ParseInt(value))
	}
	return ints
}
