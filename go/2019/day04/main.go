package advent

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() (count int) {
	input := util.InputInts("-")
	low, high := input[0], input[1]
	for x := low; x < high; x++ {
		num := fmt.Sprint(x)
		if len(num) == 6 && hasPair(num) && !isDecreasing(num) {
			count++
		}
	}
	return count
}

func hasPair(num string) bool {
	for y := 0; y < len(num)-1; y++ {
		if num[y] == num[y+1] {
			return true
		}
	}
	return false
}

func isDecreasing(num string) bool {
	for y := 0; y < len(num)-1; y++ {
		if num[y+1] < num[y] {
			return true
		}
	}
	return false
}
func Part2() (count int) {
	input := util.InputInts("-")
	low, high := input[0], input[1]
	for x := low; x < high; x++ {
		num := fmt.Sprint(x)
		if len(num) == 6 && !isDecreasing(num) && hasPair2(num) {
			count++
		}
	}
	return count
}

// This really should be a simple regular expression...
func hasPair2(num string) bool {
	for test := 0; test < len(pairs); test++ {
		if strings.Contains(num, pairs[test]) && !strings.Contains(num, triplets[test]) {
			return true
		}
	}
	return false
}

var (
	pairs    = []string{"00", "11", "22", "33", "44", "55", "66", "77", "88", "99"}
	triplets = []string{"000", "111", "222", "333", "444", "555", "666", "777", "888", "999"}
)
