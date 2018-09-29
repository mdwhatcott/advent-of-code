package day05

import "advent/lib/util"

func Part1() int {
	return NewProgram(util.ParseInts(util.InputLines())).Execute()
}

func Part2() int {
	return NewProgram(util.ParseInts(util.InputLines())).Part2().Execute()
}
