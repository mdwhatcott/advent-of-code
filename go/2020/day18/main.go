package advent

import "advent/lib/util"

var (
	part1Precedence = map[rune]int{'(': 1, '+': 0, '*': 0}
	part2Precedence = map[rune]int{'(': 2, '+': 1, '*': 0}
)

func Part1() int {
	return solve(part1Precedence, util.InputLines(), 0)
}

func Part2() int {
	return solve(part2Precedence, util.InputLines(), 0)
}

func solve(precedence map[rune]int, lines []string, sum int) int {
	if len(lines) == 0 {
		return sum
	}
	sum += EvalPostfix(ParseShuntingYard(precedence, lines[0]))
	return solve(precedence, lines[1:], sum)
}
