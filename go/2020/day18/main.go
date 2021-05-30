package advent

import (
	"strings"

	"advent/2020/day18/part1"
	"advent/lib/util"
)

func Part1() interface{} {
	input := util.InputString()
	lines := strings.Split(input, "\n")
	input = "(" + strings.Join(lines, ") + (") + ")"
	return part1.Calculate(input)
}

func Part2() interface{} {
	return nil
}
