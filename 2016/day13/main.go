package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	queue := NewLocationQueue()
	fmt.Println("Distance to (31,39):", BreadthFirstSearch(queue, util.InputInt(), 31, 39))
	fmt.Println("Locations within 50 moves:", queue.close)
}
