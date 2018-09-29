package day01

import "advent/lib/util"

func Part1() (sum int) {
	return sumOfIdenticalNeighborDigits(util.InputCharacters())
}

func Part2() (sum int) {
	return sumOfIdenticalOppositeDigits(util.InputCharacters())
}
