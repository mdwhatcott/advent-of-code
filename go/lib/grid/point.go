package grid

import "fmt"

var Origin Point

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (this Point) Y() float64 { return this.y }
func (this Point) X() float64 { return this.x }

func (this Point) Move(d Direction) Point {
	return NewPoint(this.x+d.dx, this.y+d.dy)
}

func (this Point) Offset(x, y float64) Point {
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
