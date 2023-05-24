package advent

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
	"github.com/mdwhatcott/advent-of-code/go/2020/day17/part1"
	"github.com/mdwhatcott/advent-of-code/go/2020/day17/part2"
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
