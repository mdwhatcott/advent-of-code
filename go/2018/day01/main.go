package day01

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() int {
	freq := 0
	for _, n := range util.InputInts("\n") {
		freq += n
	}
	return freq
}

func Part2() interface{} {
	freq := 0
	values := map[int]int{0: 1}
	for {
		for _, n := range util.InputInts("\n") {
			freq += n
			values[freq] = values[freq] + 1
			if values[freq] > 1 {
				return freq
			}
		}
	}
}
