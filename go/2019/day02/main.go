package advent

import (
	"advent/2019/intcode"
	"advent/lib/util"
)

func Part1() interface{} {
	program := util.InputInts(",")
	return RunTweakedProgram(program, 12, 2)[0]
}

func Part2() interface{} {
	original := util.InputInts(",")
	program := make([]int, len(original))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			copy(program, original)
			tweaked := RunTweakedProgram(program, x, y)
			if tweaked[0] == 19690720 {
				return 100*x + y
			}
		}
	}
	panic("solution unknown")
}

func RunTweakedProgram(program []int, noun, verb int) []int {
	program[1] = noun
	program[2] = verb
	return RunProgram(program)
}
func RunProgram(program []int) []int {
	return intcode.RunProgram(program, nil, nil)
}
