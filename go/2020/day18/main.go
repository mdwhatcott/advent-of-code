package advent

import "advent/lib/util"

var (
	part1Precedence = map[rune]int{'(': 1, '+': 0, '*': 0}
	part2Precedence = map[rune]int{'(': 2, '+': 1, '*': 0}
)

func Part1() int {
	return solve(part1Precedence, util.InputLines())
}

func Part2() int {
	return solve(part2Precedence, util.InputLines())
}

func solve(precedence map[rune]int, lines []string) (sum int) {
	for _, line := range lines {
		sum += EvalPostfix(string(ParseShuntingYard(precedence, line)))
	}
	return sum
}
