package advent

import "advent/lib/grid"

type Pixel struct {
	grid.Point
	ID int
}

const (
	Empty  = 0
	Wall   = 1
	Block  = 2
	Paddle = 3
	Ball   = 4
)

var draw = map[int]string{
	Empty:  " ",
	Wall:   "#",
	Block:  "-",
	Paddle: "=",
	Ball:   "*",
}
