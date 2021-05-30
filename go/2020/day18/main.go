package advent

import (
	"strings"

	"advent/2020/day18/part1"
	"advent/2020/day18/part2"
	"advent/lib/util"
)

func Part1() interface{} {
	input := util.InputString()
	lines := strings.Split(input, "\n")
	input = "(" + strings.Join(lines, ") + (") + ")"
	return part1.Calculate(input)
}

func Part2() interface{} {
	input := util.InputString()
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += part2.Calculate(line)
	}
	return sum
}
