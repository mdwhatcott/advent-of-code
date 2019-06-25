package day11

import "advent/lib/util"

func Part1() interface{} {
	return NewGrid(util.InputInt()).MaxPowerXY(3)
}

func Part2() interface{} {
	return nil
}
