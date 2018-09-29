package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	fmt.Println("Warning, part 2 takes forever and a day...")
	input := util.ParseInt(util.InputString())
	fmt.Println("Part 1 - Winning Elf:", WinningElf(input))
	fmt.Println("Part 2 - Winning Elf:", WinningElf2(input))
}
