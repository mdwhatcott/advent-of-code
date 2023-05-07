package advent

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() interface{} {
	return findMaximumAmplification(phaseCombinations, util.InputInts(","))
}

func Part2() interface{} {
	return findMaximumAmplification(phaseCombinations2, util.InputInts(","))
}

func findMaximumAmplification(combos [][]int, program []int) (max int) {
	for _, combo := range combos {
		answer := amplify(program, combo...)
		if answer > max {
			max = answer
		}
	}
	return max
}
