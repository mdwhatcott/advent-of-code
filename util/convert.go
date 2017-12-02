package util

import "strconv"

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
