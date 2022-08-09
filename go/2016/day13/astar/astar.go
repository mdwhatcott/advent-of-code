package main

import (
	"fmt"
	"log"

	"advent/lib/astar"
	"advent/lib/grid"
)

func main() {
	start := Turtle{
		seed:     1358,
		position: grid.NewPoint(1, 1),
		goal:     grid.NewPoint(31, 39),
	}
	if path, found := astar.SearchFrom(start); found {
		fmt.Println("Distance:", len(path))
	} else {
		log.Fatal("Not found: ", len(path))
	}
}

type Turtle struct {
	seed     int
	position grid.Point
	goal     grid.Point
}

func (this Turtle) AdjacentPositions() (n []astar.Turtle) {
	for _, point := range this.position.Neighbors4() {
		if this.inOpenHallway() {
			n = append(n, Turtle{seed: this.seed, position: point, goal: this.goal})
		}
	}
	return n
}

func (this Turtle) Hash() string {
	return fmt.Sprintf("%#v", this)
}

func (this Turtle) EstimatedDistanceToTarget() float64 {
	return grid.CityBlockDistance(this.position, this.goal)
}

func (this Turtle) StepCost() float64 {
	return 1.0
}

func (this Turtle) inOpenHallway() bool {
	return IsHallway(this.seed, int(this.position.X()), int(this.position.Y()))
}

func IsHallway(seed, x, y int) bool {
	return bits(sum(x, y)+seed)%2 == 0
}
func sum(x, y int) int {
	return x*x + 3*x + 2*x*y + y + y*y
}

// See: https://en.wikipedia.org/wiki/Hamming_weight
func bits(value int) int {
	var count int
	for count = 0; value > 0; count++ {
		value &= value - 1
	}
	return count
}
