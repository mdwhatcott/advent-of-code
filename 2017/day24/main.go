package day24

import "advent/lib/util"

func Part1() int {
	return FindStrongestBridge(util.InputLines())
}

func Part2() int {
	return FindStrongestLongestBridge(util.InputLines())
}
