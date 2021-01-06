package day23

import (
	"math"

	"advent/lib/util"
)

func Part1() int {
	interpreter := NewInterpreter(util.InputLines())
	interpreter.Run()
	return interpreter.mul
}

// Total rip-off: https://www.reddit.com/r/adventofcode/comments/7lms6p/2017_day_23_solutions/drnmlbk/
func Part2() (h int) {
	c := 125100 // a const given my input
	b := 108100 // a const given my input

	for ; b < c+1; b += 17 {
		if !prime(b) {
			h += 1
		}
	}
	return h
}

func prime(i int) bool {
	for j := 2; j < int(math.Sqrt(float64(i)))+1; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
