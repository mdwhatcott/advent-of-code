package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	shortestPath, longestDistance := Navigate(util.InputString())
	fmt.Println("Part 1:", shortestPath)
	fmt.Println("Part 2:", longestDistance)
}
