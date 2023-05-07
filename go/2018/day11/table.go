package day11

import "github.com/mdwhatcott/advent-of-code/go/lib/grid"

type SummedAreaTable map[grid.Point]int

func NewSummedAreaTable(source map[grid.Point]int) SummedAreaTable {
	summed := make(SummedAreaTable, len(source))
	for point := range source {
		summed[point] = summed.sum(source, point)
	}
	return summed
}

func (this SummedAreaTable) sum(source map[grid.Point]int, at grid.Point) int {
	value, found := this[at]
	if found {
		return value
	}
	if at.X() == 0 && at.Y() == 0 {
		this[at] = source[at]
		return this.sum(source, at)
	}
	if at.X() == 0 {
		this[at] = source[at] + this.sum(source, at.Offset(0, -1))
		return this.sum(source, at)
	}
	if at.Y() == 0 {
		this[at] = source[at] + this.sum(source, at.Offset(-1, 0))
		return this.sum(source, at)
	}
	this[at] = source[at] +
		this.sum(source, at.Offset(-1, 0)) +
		this.sum(source, at.Offset(0, -1)) -
		this.sum(source, at.Offset(-1, -1))
	return this.sum(source, at)
}

func (this SummedAreaTable) SummedArea(at grid.Point, size int) int {
	return this[at.Offset(-1, -1)] +
		this[at.Offset(float64(size-1), float64(size-1))] -
		this[at.Offset(float64(size-1), -1)] -
		this[at.Offset(-1, float64(size-1))]
}
