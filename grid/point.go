package grid

import "fmt"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (this Point) Y() float64 { return this.y }
func (this Point) X() float64 { return this.x }

func (this Point) Offset(x, y float64) Point {
	return NewPoint(this.x+x, this.y+y)
}

func (this Point) String() string {
	return fmt.Sprintf("(%v, %v)", this.x, this.y)
}

func (this Point) Neighbors4() (neighbors []Point) {
	neighbors = append(neighbors, NewPoint(this.x+1, this.y+0)) // right
	neighbors = append(neighbors, NewPoint(this.x-1, this.y+0)) // left
	neighbors = append(neighbors, NewPoint(this.x+0, this.y+1)) // top
	neighbors = append(neighbors, NewPoint(this.x+0, this.y-1)) // bottom

	return neighbors
}

func (this Point) Neighbors8() (neighbors []Point) {
	neighbors = this.Neighbors4()
	neighbors = append(neighbors, NewPoint(this.x+1, this.y+1)) // top-right
	neighbors = append(neighbors, NewPoint(this.x-1, this.y-1)) // bottom-left
	neighbors = append(neighbors, NewPoint(this.x+1, this.y-1)) // bottom-right
	neighbors = append(neighbors, NewPoint(this.x-1, this.y+1)) // top-right
	return neighbors
}
