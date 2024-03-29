package day15

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	"github.com/mdwhatcott/astar"
)

func Part1(lines []string) int {
	grid := make(Grid)
	var max intgrid.Point
	for y, line := range lines {
		for x, char := range line {
			point := intgrid.NewPoint(x, y)
			grid[point] = float64(int(char - '0'))
			max = point
		}
	}
	return distance(grid, max)
}

func Part2(lines []string) int {
	grid := make(Grid)
	var max intgrid.Point
	for gridRow := 0; gridRow < 5; gridRow++ {
		for gridCol := 0; gridCol < 5; gridCol++ {
			for y, line := range lines {
				for x, char := range line {
					point := intgrid.NewPoint((gridCol*len(lines))+x, (gridRow*len(lines))+y)
					value := int(char - '0')
					for gr := gridRow; gr > 0; gr-- {
						value++
						if value == 10 {
							value = 1
						}
					}
					for gc := gridCol; gc > 0; gc-- {
						value++
						if value == 10 {
							value = 1
						}
					}
					grid[point] = float64(value)
					max = point
				}
			}
		}
	}
	return distance(grid, max)
}

func distance(grid Grid, max intgrid.Point) int {
	start := intgrid.NewPoint(0, 0)
	path, found := astar.SearchFrom(NewTurtle(grid, start, max, 0))
	if !found {
		panic("NOPE")
	}
	last := path[len(path)-1]
	return int(last.(*Turtle).distance)
}

type Grid map[intgrid.Point]float64

func (this Grid) contains(p intgrid.Point) bool {
	_, ok := this[p]
	return ok
}

type Turtle struct {
	grid     Grid
	point    intgrid.Point
	target   intgrid.Point
	distance float64
}

func NewTurtle(grid Grid, point, target intgrid.Point, distance float64) astar.Turtle {
	return &Turtle{
		grid:     grid,
		point:    point,
		target:   target,
		distance: distance,
	}
}

func (this *Turtle) EstimatedDistanceToTarget() float64 {
	return float64(intgrid.CityBlockDistance(this.point, this.target))
}

func (this *Turtle) StepCost() float64 {
	return this.grid[this.point]
}

func (this *Turtle) AdjacentPositions() (results []astar.Turtle) {
	for _, direction := range intgrid.Neighbors4 {
		neighbor := this.point.Move(direction)
		if this.grid.contains(neighbor) {
			distance := this.distance + this.grid[neighbor]
			results = append(results, NewTurtle(this.grid, neighbor, this.target, distance))
		}
	}
	return results
}

func (this *Turtle) Hash() string {
	return this.point.String()
}
