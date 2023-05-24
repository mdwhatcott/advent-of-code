package day09

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() interface{} {
	return MarbleHighScore(Parse(util.InputString()))
}

func Part2() interface{} {
	return MarbleHighScore(Parse10X(util.InputString()))
}
