package astar

import (
	"github.com/mdwhatcott/go-collections/set"

	"advent/lib/intgrid"
)

type GridTurtle struct {
	grid     set.Set[intgrid.Point]
	from, to intgrid.Point
}

func NewGridTurtle(grid set.Set[intgrid.Point], from, to intgrid.Point) *GridTurtle {
	return &GridTurtle{grid: grid, from: from, to: to}
}
func (this *GridTurtle) Search() (path []Turtle, found bool) {
	return SearchFrom(this)
}
func (this *GridTurtle) EstimatedDistanceToTarget() float64 {
	return intgrid.CityBlockDistance(this.from, this.to)
}
func (this *GridTurtle) AdjacentPositions() (results []Turtle) {
	for _, d := range intgrid.Neighbors4 {
		at := this.from.Move(d)
		if this.grid.Contains(at) {
			results = append(results, NewGridTurtle(this.grid, at, this.to))
		}
	}
	return results
}

func (this *GridTurtle) StepCost() float64 { return 1 }
func (this *GridTurtle) Hash() string      { return this.from.String() }
func (this *GridTurtle) At() intgrid.Point { return this.from }
