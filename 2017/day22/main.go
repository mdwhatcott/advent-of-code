package day22

import "advent/lib/util"

func Part1() int {
	virus := NewVirus(util.InputString())
	for x := 0; x < 10000; x++ {
		virus.Move()
	}
	return virus.Infected()
}

func Part2() int {
	virus := NewVirus2(util.InputString())
	for x := 0; x < 10000000; x++ {
		virus.Move()
	}
	return virus.Infected()
}