package util

import (
	"fmt"
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

func BinaryHammingWeight(value int) (count int) { // See: https://en.wikipedia.org/wiki/Hamming_weight
	for count = 0; value > 0; count++ {
		value &= value - 1
	}
	return count
}

func EncodeBinary(value byte) string {
	return fmt.Sprintf("%08b", value)
}
