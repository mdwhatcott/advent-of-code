package util

import (
	"strconv"
	"strings"
)

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

// See: https://en.wikipedia.org/wiki/Hamming_weight
func BinaryHammingWeight(value int) (count int) {
	for count = 0; value > 0; count++ {
		value &= value - 1
	}
	return count
}

func EncodeBinary(value byte) string {
	var buffer = make([]string, 8)
	for i := 0; i < 8; i++ {
		buffer[7-i] = strconv.Itoa((int(value) >> uint(i)) & 1)
	}
	return strings.Join(buffer, "")
}
