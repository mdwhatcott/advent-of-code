package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	input := util.InputString()

	fmt.Println("Part 1 - Dragon Checksum (length: 272):     ", DragonChecksum(input, 272))
	fmt.Println("Part 2 - Dragon Checksum (length: 35651584):", DragonChecksum(input, 35651584))
}
