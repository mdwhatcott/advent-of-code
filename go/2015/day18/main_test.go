package main

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test(t *testing.T) {
	assert.So(t, part1(), should.Equal, 814)
	assert.So(t, part2(), should.Equal, 924)
}

func part1() int {
	input := util.InputString()
	input = strings.Replace(input, ".", "-", -1)
	input = strings.Replace(input, "#", "x", -1)
	return runSimulation(NewGrid(input))
}

func runSimulation(grid *Grid) int {
	for x := 0; x < 100; x++ {
		grid.Scan()
	}

	return grid.CountAlive()
}

func part2() int {
	input := util.InputString()
	input = strings.Replace(input, ".", "-", -1)
	input = strings.Replace(input, "#", "x", -1)
	grid := NewGrid(input)
	grid.LockCornerLightsOn()
	return runSimulation(grid)
}
