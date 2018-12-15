package day07

import (
	"advent/lib/util"
)

func Part1() interface{} {
	return NewTopologicalSort(util.InputString()).Sort()
}

func Part2() interface{} {
	return nil
}
