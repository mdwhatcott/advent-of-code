package day19

import "fmt"

var Origin Point

type Point struct {
	x int
	y int
	z int
}

func NewPoint(x, y, z int) Point {
	return Point{x: x, y: y, z: z}
}

func (this Point) Y() int { return this.y }
func (this Point) X() int { return this.x }
func (this Point) Z() int { return this.z }

func (this Point) Move(d Direction) Point {
	return NewPoint(this.x+d.dx, this.y+d.dy, this.z+d.dz)
}

func (this Point) Offset(x, y, z int) Point {
	return NewPoint(this.x+x, this.y+y, this.z+z)
}

func (this Point) String() string {
	return fmt.Sprintf("(%v, %v, %v)", this.x, this.y, this.z)
}
