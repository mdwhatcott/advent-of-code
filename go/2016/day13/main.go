package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func main() {
	origin := intgrid.NewPoint(1, 1)
	target := intgrid.NewPoint(31, 39)
	distance, near := BreadthFirstSearch(util.InputInt(), origin, target)
	fmt.Println("Distance to (31,39):", distance)
	fmt.Println("Queue within 50 moves:", near)
}
