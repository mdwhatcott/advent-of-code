package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	queue := NewLocationQueue()
	fmt.Println("Distance to (31,39):", DistanceToDestination(queue, util.ParseInt(util.InputString()), 31, 39))
	fmt.Println("Locations within 50 moves:", queue.close)
}
