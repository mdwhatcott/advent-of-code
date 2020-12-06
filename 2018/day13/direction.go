package day13

type Direction struct{ dx, dy int }

func NewDirection(dx, dy int) Direction {
	return Direction{dx: dx, dy: dy}
}

func (this Direction) TurnSlash(slash rune) Direction { return slashTurns[slash][this] }
func (this Direction) TurnRight() Direction           { return Clockwise[this] }
func (this Direction) TurnLeft() Direction            { return CounterClockwise[this] }

var (
	Right = NewDirection(1, 0)
	Left  = NewDirection(-1, 0)
	Up    = NewDirection(0, -1)
	Down  = NewDirection(0, 1)

	Clockwise        = map[Direction]Direction{Down: Left, Left: Up, Up: Right, Right: Down}
	CounterClockwise = map[Direction]Direction{Down: Right, Right: Up, Up: Left, Left: Down}
)

var cartDirections = map[rune]Direction{
	'^': Up,
	'v': Down,
	'<': Left,
	'>': Right,
}

var slashTurns = map[rune]map[Direction]Direction{
	'/': {
		Up:    Right,
		Down:  Left,
		Right: Up,
		Left:  Down,
	},
	'\\': {
		Up:    Left,
		Down:  Right,
		Right: Down,
		Left:  Up,
	},
}

var cartPositions = map[Direction]byte{
	Up:    '^',
	Down:  'v',
	Left:  '<',
	Right: '>',
}
