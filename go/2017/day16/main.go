package day16

import (
	"strings"

	"advent/lib/util"
)

const start = "abcdefghijklmnop"

var danceSteps = strings.Split(util.InputString(), ",")

func Part1() string {
	interpreter := NewInterpreter(start)
	interpreter.doCompleteDance()
	return interpreter.state
}

func Part2() string {
	interpreter := NewInterpreter(start)

	for x := 0; x < 1000000000%getDanceCycleLength(); x++ {
		interpreter.doCompleteDance()
	}

	return interpreter.state
}

func getDanceCycleLength() int {
	interpreter := NewInterpreter(start)
	cycle := make(map[string]int)

	for x := 0; x < 1000; x++ {
		interpreter.doCompleteDance()

		cycle[interpreter.state]++
		if cycle[interpreter.state] > 1 {
			return len(cycle)
		}
	}
	panic("no cycle?")
}
