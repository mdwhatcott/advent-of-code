package main

import (
	"fmt"
	"strings"

	"advent/lib/util"
	"github.com/mdwhatcott/golife/life"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 814))
	fmt.Println(assert.So(part2(), should.Equal, 924))
}

func part1() int {
	input := util.InputString()
	input = strings.Replace(input, ".", "-", -1)
	input = strings.Replace(input, "#", "x", -1)
	return runSimulation(life.New(input))
}

func runSimulation(grid *life.Grid) int {
	for x := 0; x < 100; x++ {
		grid.Scan()
	}

	return grid.CountAlive()
}

func part2() int {
	input := util.InputString()
	input = strings.Replace(input, ".", "-", -1)
	input = strings.Replace(input, "#", "x", -1)
	grid := life.New(input)
	grid.LockCornerLightsOn()
	return runSimulation(grid)
}
