package day05

import "advent/lib/util"

func Part1() interface{} {
	return len(react(util.InputString()))
}

func Part2() interface{} {
	return len(reactAggressive(util.InputString()))
}
