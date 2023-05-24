package day21

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

var realRules = RegisterEnhancementRules(util.InputLines()...)

func Part1() int {
	return CountFractalPixels(realRules, 5)
}

func Part2() int {
	return CountFractalPixels(realRules, 18)
}
