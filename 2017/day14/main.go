package day14

import "advent/lib/util"

func Part1() int {
	return makeDisk(util.InputString()).CountUsedSectors()
}

func Part2() int {
	return makeDisk(util.InputString()).CountUsedRegions()
}

