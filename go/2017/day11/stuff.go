package day11

import (
	"advent/lib/maths"
)

// https://www.redblobgames.com/grids/hexagons/
type Hex struct {
	X, Y, Z int
}

func (this Hex) Offset(offset Hex) Hex {
	return Hex{
		X: this.X + offset.X,
		Y: this.Y + offset.Y,
		Z: this.Z + offset.Z,
	}
}

func (this Hex) Neighbors() []Hex {
	return []Hex{
		this.Offset(NorthWest),
		this.Offset(North),
		this.Offset(NorthEast),
		this.Offset(SouthEast),
		this.Offset(South),
		this.Offset(SouthWest),
	}
}

func (this Hex) DistanceTo(other Hex) int {
	return maths.Max(Abs(this.X-other.X), Abs(this.Y-other.Y), Abs(this.Z-other.Z))
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var (
	North     = Hex{0, 1, -1}
	NorthEast = Hex{1, 0, -1}
	NorthWest = Hex{-1, 1, 0}
	South     = Hex{0, -1, 1}
	SouthEast = Hex{1, -1, 0}
	SouthWest = Hex{-1, 0, 1}
)
