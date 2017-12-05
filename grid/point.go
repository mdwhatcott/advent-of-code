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

//////////////////////////////////////////////////////////

type Direction struct{ dx, dy float64 }

func NewDirection(dx, dy float64) Direction {
	return Direction{dx: dx, dy: dy}
}

func (this Direction) Dx() float64 { return this.dx }
func (this Direction) Dy() float64 { return this.dy }

var (
	Right = Direction{dx: 1, dy: 0}
	Left  = Direction{dx: -1, dy: 0}
	Up    = Direction{dx: 0, dy: 1}
	Down  = Direction{dx: 0, dy: -1}

	TopRight    = Direction{1, 1}
	TopLeft     = Direction{-1, 1}
	BottomRight = Direction{1, -1}
	BottomLeft  = Direction{-1, -1}

	Neighbors4 = []Direction{Right, Left, Up, Down}
	Neighbors8 = append(Neighbors4, []Direction{TopRight, TopLeft, BottomRight, BottomLeft}...)
)
