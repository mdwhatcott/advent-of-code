package intgrid

import "fmt"

var Origin Point

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (this Point) Y() int { return this.y }
func (this Point) X() int { return this.x }

func (this Point) Move(d Direction) Point {
	return NewPoint(this.x+d.dx, this.y+d.dy)
}

func (this Point) Offset(x, y int) Point {
	return NewPoint(this.x+x, this.y+y)
}

func (this Point) String() string {
	return fmt.Sprintf("(%v, %v)", this.x, this.y)
}

func (this Point) Neighbors4() (neighbors []Point) {
	for _, offset := range Neighbors4 {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}

func (this Point) Neighbors8() (neighbors []Point) {
	for _, offset := range Neighbors8 {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}
