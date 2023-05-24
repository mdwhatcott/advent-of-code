package day07

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() interface{} {
	order1, _ := NewTopologicalSort(util.InputString(), 1, 1).Sort()
	return order1
}

func Part2() interface{} {
	_, seconds5 := NewTopologicalSort(util.InputString(), 5, 60).Sort()
	return seconds5
}
