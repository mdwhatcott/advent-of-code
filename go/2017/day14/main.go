package day14

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() int {
	return makeDisk(util.InputString()).CountUsedSectors()
}

func Part2() int {
	return makeDisk(util.InputString()).CountUsedRegions()
}
