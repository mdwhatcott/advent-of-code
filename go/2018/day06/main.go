package day06

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() int {
	return int(calculateLargestFiniteArea(parsePoints(util.InputString())))
}

func Part2() int {
	return calculateAreaWithinRadius(parsePoints(util.InputString()), 10000)
}
