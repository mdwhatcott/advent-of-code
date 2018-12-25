package day07

import (
	"advent/lib/util"
)

func Part1() interface{} {
	return NewTopologicalSort(util.InputString()).Sort()
}

func Part2() interface{} {
	sorter := NewConcurrentTopologicalSort(util.InputString(), 5, 60)
	sorter.Sort()
	return sorter.DurationSeconds()
}
