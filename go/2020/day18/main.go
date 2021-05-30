package advent

import (
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	input := util.InputString()
	lines := strings.Split(input, "\n")
	input = "(" + strings.Join(lines, ") + (") + ")"
	return Calculate(input)
}

func Part2() interface{} {
	return nil
}
