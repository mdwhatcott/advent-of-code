package main

import (
	"fmt"

	"advent/lib/astar"
	"advent/lib/util"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println("Part 1: Bot travelled:", assert.So(part1(), should.Equal, 460).Fatal())
	fmt.Println("Part 2: To return, bot travelled:", assert.So(part2(), should.Equal, 668).Fatal())
}

func part1() int {
	if path, found := astar.SearchFrom(StartPoint(util.InputLines())); found {
		return len(path) - 1
	}
	return -1
}

func part2() int {
	if path, found := astar.SearchFrom(ReturnTo(StartPoint(util.InputLines()))); found {
		return len(path) - 1
	}
	return -1
}
