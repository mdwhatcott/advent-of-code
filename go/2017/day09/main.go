package day09

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() int {
	return groupScore(util.InputString())
}

func Part2() int {
	return garbageScore(util.InputString())
}
