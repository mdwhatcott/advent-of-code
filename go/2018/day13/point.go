package day13

import "fmt"

type Point struct{ x, y int }

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

var Origin = NewPoint(0, 0)
