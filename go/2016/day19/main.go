package main

import (
	"fmt"

	"advent/lib/parse"
	"advent/lib/util"
)

func main() {
	fmt.Println("Warning, part 2 takes forever and a day...")
	input := parse.Int(util.InputString())
	fmt.Println("Part 1 - Winning Elf:", WinningElf(input))
	fmt.Println("Part 2 - Winning Elf:", WinningElf2(input))
}
