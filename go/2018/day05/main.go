package day05

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() interface{} {
	return len(react(util.InputString()))
}

func Part2() interface{} {
	return len(reactAggressive(util.InputString()))
}
