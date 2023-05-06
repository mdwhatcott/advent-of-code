package day05

import (
	"advent/lib/parse"
	"advent/lib/util"
)

func Part1() int {
	return NewProgram(parse.Ints(util.InputLines())).Execute()
}

func Part2() int {
	return NewProgram(parse.Ints(util.InputLines())).Part2().Execute()
}
