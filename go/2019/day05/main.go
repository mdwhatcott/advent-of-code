package advent

import (
	"log"

	"advent/2019/intcode"
	"advent/lib/util"
)

func Part1() interface{} {
	var outputs []int
	intcode.RunProgram(
		util.InputInts(","),
		func() int { return 1 },
		func(out int) { outputs = append(outputs, out) },
	)

	for i, out := range outputs {
		if i < len(outputs)-1 {
			if out != 0 {
				log.Panic("FAILED:", outputs)
			}
		}
	}
	return outputs[len(outputs)-1]
}

func Part2() interface{} {
	var outputs []int
	intcode.RunProgram(
		util.InputInts(","),
		func() int { return 5 },
		func(out int) { outputs = append(outputs, out) },
	)

	return outputs[len(outputs)-1]
}
