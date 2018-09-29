package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	lines := util.InputLines()

	scrambler := NewScrambler(lines)
	fmt.Println("Part 1 - Scrambled password:", scrambler.Process("abcdefgh"))

	scrambler = NewUnscrambler(lines)
	fmt.Println("Part 2 - Unscrambled password:", scrambler.Process("fbgdceah"))
}
