package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	safe := 0
	row := ParseRow(util.InputString())
	safe += row.Safe()

	for x := 0; x < 39; x++ {
		row = row.Next()
		safe += row.Safe()
	}

	fmt.Println("Part 1:", safe)

	for x := 0; x < 400000-40; x++ {
		row = row.Next()
		safe += row.Safe()
	}

	fmt.Println("Part 2:", safe)
}
