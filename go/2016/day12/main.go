package main

import (
	"fmt"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/2016/util/assembunny"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	fmt.Println(`Part 1 - Value of 'a':`)
	should.So(nil, part1(), should.Equal, 318077)

	fmt.Println(`Part 2 - Value of 'a' (when 'c' starts at 1):`)
	should.So(nil, part2(), should.Equal, 9227731)
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
