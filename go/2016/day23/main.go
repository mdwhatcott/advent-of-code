package main

import (
	"fmt"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/2016/util/assembunny"
	"advent/lib/util"
)

func main() {
	fmt.Println(`Part 1 - Password for Safe (register 'a' starts at 7):`)
	assert.So(t, nil, part1(), should.Equal, 12762)

	fmt.Println("WARNING: Part 2 takes about five minutes to execute 3500000000+ instructions.")
	fmt.Println("FUTURE:  It could be optimized (https://www.reddit.com/r/adventofcode/comments/5jvbzt/2016_day_23_solutions/)")

	fmt.Println(`Part 2 - Password for Safe (register 'a' starts at 12, with MULTIPLICATION):`)
	assert.So(t, nil, part2(), should.Equal, 479009322)
}

func part1() int {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.Set("a", 7)
	interpreter.ExecuteProgram()
	return interpreter.Get("a")
}

func part2() int {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.Set("a", 12)
	interpreter.ExecuteProgram()
	return interpreter.Get("a")
}
