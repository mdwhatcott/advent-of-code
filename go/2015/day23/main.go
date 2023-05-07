package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	interpreter := NewInterpreter(util.InputLines())
	interpreter.ExecuteProgram()
	fmt.Println("Part 1, Register B should be 255:", interpreter.registers["b"])
}

func part2() {
	interpreter := NewInterpreter(util.InputLines())
	interpreter.registers["a"] = 1
	interpreter.ExecuteProgram()
	fmt.Println("Part 2, Register B should be 334:", interpreter.registers["b"])
}
