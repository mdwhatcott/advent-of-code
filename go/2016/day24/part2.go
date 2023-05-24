package main

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
	"github.com/mdwhatcott/astar"
)

type ReturnTurtle struct {
	*Turtle

	start grid.Point
}

func ReturnTo(turtle *Turtle) *ReturnTurtle {
	return &ReturnTurtle{Turtle: turtle, start: turtle.point}
}

func (this ReturnTurtle) EstimatedDistanceToTarget() float64 {
	return this.Turtle.EstimatedDistanceToTarget() + grid.CityBlockDistance(this.point, this.start)
}

func (this ReturnTurtle) AdjacentPositions() (neighbors []astar.Turtle) {
	for _, neighbor := range this.Turtle.AdjacentPositions() {
		neighbors = append(neighbors, &ReturnTurtle{Turtle: neighbor.(*Turtle), start: this.start})
	}
	return neighbors
}
