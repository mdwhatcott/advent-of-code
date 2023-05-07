package day01

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() (sum int) {
	return sumOfIdenticalNeighborDigits(InputCharacters())
}

func Part2() (sum int) {
	return sumOfIdenticalOppositeDigits(InputCharacters())
}

func InputCharacters() (all []string) {
	for _, c := range util.InputString() {
		all = append(all, string(c))
	}
	return all
}
