package intgrid

import (
	"github.com/mdwhatcott/advent-of-code/go/lib/maths"
)

type BoundingBox struct {
	corner1 Point
	corner2 Point
}

func NewBoundingBox(corner1, corner2 Point) BoundingBox {
	return BoundingBox{
		corner1: corner1,
		corner2: corner2,
	}
}

func (this BoundingBox) MinY() int { return maths.Min(this.corner1.Y(), this.corner2.Y()) }
func (this BoundingBox) MaxY() int { return maths.Max(this.corner1.Y(), this.corner2.Y()) }
func (this BoundingBox) MinX() int { return maths.Min(this.corner1.X(), this.corner2.X()) }
func (this BoundingBox) MaxX() int { return maths.Max(this.corner1.X(), this.corner2.X()) }

func (this BoundingBox) Contains(point Point) bool {
	return point.X() >= this.MinX() &&
		point.X() <= this.MaxX() &&
		point.Y() >= this.MinY() &&
		point.Y() <= this.MaxY()
}
