package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/go/lib/astar"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"

	"github.com/mdwhatcott/testing/should"
)

func main() {
	fmt.Println("Part 1: Bot travelled:")
	should.So(nil, part1(), should.Equal, 460)

	fmt.Println("Part 2: To return, bot travelled:")
	should.So(nil, part2(), should.Equal, 668)
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
