package advent

import (
	"advent/2020/day17/part1"
	"advent/2020/day17/part2"
	"advent/lib/util"
)

func Part1() interface{} {
	world := part1.ParseInitialWorld(util.InputString())
	world.Boot()
	return len(world)
}

func Part2() interface{} {
	world := part2.ParseInitialWorld(util.InputString())
	world.Boot()
	return len(world)
}
