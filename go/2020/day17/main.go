package advent

import (
	"advent/2020/day17/part1"
	"advent/lib/util"
)

func Part1() interface{} {
	world := part1.ParseInitialWorld(util.InputString())
	world.Boot()
	return len(world)
}

func Part2() interface{} {
	return nil
}
