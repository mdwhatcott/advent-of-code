package day24

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() int {
	return FindStrongestBridge(util.InputLines())
}

func Part2() int {
	return FindStrongestLongestBridge(util.InputLines())
}
