package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	shortestPath, longestDistance := Navigate(util.InputString())
	fmt.Println("Part 1:", shortestPath)
	fmt.Println("Part 2:", longestDistance)
}
