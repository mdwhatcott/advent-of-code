package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code-2016/util/assembunny"
	"advent/lib/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(`Part 1 - Value of 'a':`,
		assert.So(part1(), should.Equal, 318077).Fatal())

	fmt.Println(`Part 2 - Value of 'a' (when 'c' starts at 1):`,
		assert.So(part2(), should.Equal, 9227731).Fatal())
}

func part1() int {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.ExecuteProgram()
	return interpreter.Get("a")
}

func part2() int {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.Set("c", 1)
	interpreter.ExecuteProgram()
	return interpreter.Get("a")
}
