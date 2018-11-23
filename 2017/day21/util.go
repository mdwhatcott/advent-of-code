package day21

import (
	"math"
	"strings"
)

func sizeOf(grid string) int {
	return sqrt(len(strings.Replace(grid, "/", "", -1)))
}
func sqrt(v int) int {
	return int(math.Sqrt(float64(v)))
}
func contains(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}
