package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	start := Parse(util.InputLines())
	start.Receive(start.start)
	fmt.Println("Part 2 - Sum of outputs 0-2:",
		start.outs[0].value*start.outs[1].value*start.outs[2].value)
}
