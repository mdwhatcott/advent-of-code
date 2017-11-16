package util

import "strconv"

func ParseInt(value string) int {
	parsed, _ := strconv.Atoi(value)
	return parsed
}
