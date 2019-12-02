package advent

import (
	"advent/lib/util"
	_ "advent/lib/util"
)

func Part1() interface{} {
	program := util.InputInts(",")
	RunProgram(program)
	return program[0]
}

func Part2() interface{} {
	return nil
}
