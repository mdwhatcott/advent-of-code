package main

import (
	"fmt"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/2016/util/assembunny"
	"advent/lib/util"
)

func main() {
	fmt.Println(`Part 1 - Value of 'a':`)
	assert.So(t, nil, part1(), should.Equal, 318077)

	fmt.Println(`Part 2 - Value of 'a' (when 'c' starts at 1):`)
	assert.So(t, nil, part2(), should.Equal, 9227731)
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
