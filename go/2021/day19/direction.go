package day19

type Direction struct{ dx, dy, dz int }

func NewDirection(dx, dy, dz int) Direction {
	return Direction{dx: dx, dy: dy, dz: dz}
}

func (this Direction) Dx() int { return this.dx }
func (this Direction) Dy() int { return this.dy }
func (this Direction) Dz() int { return this.dz }

func (this Direction) Add(d Direction) Direction {
	return NewDirection(
		this.dx+d.dx,
		this.dy+d.dy,
		this.dz+d.dz,
	)
}
func (this Direction) Sub(d Direction) Direction {
	return NewDirection(
		this.dx-d.dx,
		this.dy-d.dy,
		this.dz-d.dz,
	)
}
func (this Direction) Backtrack() Direction {
	return NewDirection(-this.dx, -this.dy, -this.dz)
}
