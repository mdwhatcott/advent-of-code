package day09

import "advent/lib/util"

func Part1() interface{} {
	return MarbleHighScore(Parse(util.InputString()))
}

func Part2() interface{} {
	return MarbleHighScore(Parse10X(util.InputString()))
}
