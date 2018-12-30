package day08

import "advent/lib/util"

func Part1() interface{} {
	return SumMetadata(ParseTree(LoadInputs(util.InputInts(" "))))
}

func Part2() interface{} {
	return RootValue(ParseTree(LoadInputs(util.InputInts(" "))))
}
