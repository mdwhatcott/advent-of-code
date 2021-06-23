package advent

import "advent/lib/util"

var part1Precedence = map[rune]int{'(': 1}

func Part1() interface{} {
	sum := 0
	for _, line := range util.InputLines() {
		sum += EvalPostfix(string(ParseShuntingYard(part1Precedence, line)))
	}
	return sum
}

func Part2() interface{} {
	return nil
}
