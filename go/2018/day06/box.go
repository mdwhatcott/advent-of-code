package day06

import (
	"advent/lib/grid"
)

type BoundingBox struct {
	maxX float64
	maxY float64
	minX float64
	minY float64
}

func NewBoundingBox(points []grid.Point) *BoundingBox {
	box := &BoundingBox{
		maxX: -1000,
		maxY: -1000,
		minX: 1000,
		minY: 1000,
	}
	for _, point := range points {
		x := point.X()
		y := point.Y()
		if x > box.maxX {
			box.maxX = x
		}
		if x < box.minX {
			box.minX = x
		}
		if y > box.maxY {
			box.maxY = y
		}
		if y < box.minY {
			box.minY = y
		}
	}
	return box
}

func (this *BoundingBox) IsOnBoundary(point grid.Point) bool {
	x := point.X()
	y := point.Y()
	return x == this.maxX || x == this.maxY || y == this.minX || y == this.minY
}
