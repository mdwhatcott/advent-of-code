package day09

import "advent/lib/util"

func Part1() int {
	return groupScore(util.InputString())
}

func Part2() int {
	return garbageScore(util.InputString())
}
