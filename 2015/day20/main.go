package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	max := 0
	startingHouse := 660000 // found by trial and error
	for house := startingHouse; ; house++ {
		count := presents(house)
		if count > max {
			fmt.Println("Progress:", house, count)
			max = count
		}
		if count >= util.InputInt() {
			fmt.Println("Part 1:", house, count) // 665280
			break
		}
	}
}

func presents(house int) (presents int) {
	for elf := 1; elf <= house; elf++ {
		if house%elf == 0 {
			presents += elf * 10
		}
	}
	return presents
}
