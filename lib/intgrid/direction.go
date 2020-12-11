package intgrid

type Direction struct{ dx, dy int }

func NewDirection(dx, dy int) Direction {
	return Direction{dx: dx, dy: dy}
}

func (this Direction) Dx() int { return this.dx }
func (this Direction) Dy() int { return this.dy }

func (this Direction) TurnRight() Direction { return Clockwise[this] }
func (this Direction) TurnLeft() Direction  { return CounterClockwise[this] }

var (
	Static Direction

	Right = NewDirection(1, 0)
	Left  = NewDirection(-1, 0)
	Up    = NewDirection(0, 1)
	Down  = NewDirection(0, -1)

	TopRight    = NewDirection(1, 1)
	TopLeft     = NewDirection(-1, 1)
	BottomRight = NewDirection(1, -1)
	BottomLeft  = NewDirection(-1, -1)

	Neighbors4 = []Direction{Right, Left, Up, Down}
	Diagonals4 = []Direction{TopRight, TopLeft, BottomRight, BottomLeft}
	Neighbors8 = append(Neighbors4, Diagonals4...)

	Clockwise = map[Direction]Direction{
		Down:  Left,
		Left:  Up,
		Up:    Right,
		Right: Down,
	}
	CounterClockwise = map[Direction]Direction{
		Down:  Right,
		Right: Up,
		Up:    Left,
		Left:  Down,
	}
)
