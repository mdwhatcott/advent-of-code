package main

import (
	"fmt"

	"advent/lib/intgrid"
	"advent/lib/util"
)

func main() {
	target := intgrid.NewPoint(31, 39)
	distance, near := BreadthFirstSearch(util.InputInt(), target)
	fmt.Println("Distance to (31,39):", distance)
	fmt.Println("Queue within 50 moves:", near)
}
