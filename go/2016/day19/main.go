package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func main() {
	fmt.Println("Warning, part 2 takes forever and a day...")
	input := parse.Int(util.InputString())
	fmt.Println("Part 1 - Winning Elf:", WinningElf(input))
	fmt.Println("Part 2 - Winning Elf:", WinningElf2(input))
}
