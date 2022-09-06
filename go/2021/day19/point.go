package day19

import (
	"fmt"

	"advent/lib/util"
)

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

func Diff(p1, p2 Point) Direction {
	return NewDirection(
		p1.X()-p2.X(),
		p1.Y()-p2.Y(),
		p1.Z()-p2.Z())
}

func Manhattan(p1, p2 Point) int {
	return 0 +
		util.Abs(p1.X()-p2.X()) +
		util.Abs(p1.Y()-p2.Y()) +
		util.Abs(p1.Z()-p2.Z())

}
