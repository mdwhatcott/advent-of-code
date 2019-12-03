package advent

import (
	"advent/lib/util"
	_ "advent/lib/util"
)

func Part1() interface{} {
	program := util.InputInts(",")
	program[1] = 12
	program[2] = 2
	RunProgram(program)
	return program[0]
}

func Part2() interface{} {
	original := util.InputInts(",")
	program := make([]int, len(original))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			copy(program, original)
			program[1] = x
			program[2] = y
			RunProgram(program)
			if program[0] == 19690720 {
				return 100*x + y
			}
		}
	}
	panic("solution unknown")
}
