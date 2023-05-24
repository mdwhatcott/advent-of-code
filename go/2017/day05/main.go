package day05

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() int {
	return NewProgram(parse.Ints(util.InputLines())).Execute()
}

func Part2() int {
	return NewProgram(parse.Ints(util.InputLines())).Part2().Execute()
}
