package day18

import "advent/lib/util"

func Part1() int {
	interpreter := NewInterpreter(util.InputLines())
	interpreter.Run()
	return interpreter.recovered
}

func Part2() int {
	return 0
}