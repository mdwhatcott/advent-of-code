package day19

import (
	"strings"
	"unicode"
)

type Turtle struct {
	maze      []string
	position  Point
	direction Direction
	sequence  string
}

func NewTurtle(maze []string) *Turtle {
	return &Turtle{
		maze:      maze,
		position:  Point{X: strings.Index(maze[0], "|"), Y: 0},
		direction: South,
	}
}

func (this *Turtle) Move() {
	this.position = this.position.Move(this.direction)
}

func (this *Turtle) Orient() (eof bool) {
	spot := this.spot(this.position)
	if spot == '+' {
		this.turn()
	} else if unicode.IsLetter(spot) {
		this.sequence += string(spot)
	}
	return !unicode.IsSpace(spot)
}

func (this *Turtle) spot(position Point) rune {
	x := position.X
	y := position.Y
	if y < 0 || y >= len(this.maze) || x < 0 || x >= len(this.maze[y]) {
		return 0
	}
	return rune(this.maze[y][x])
}

func (this *Turtle) turn() {
	left := Left[this.direction]
	right := Right[this.direction]

	if leftSpot := this.peek(left); leftSpot != 0 && leftSpot != ' ' {
		this.direction = left
	} else {
		this.direction = right
	}
}

func (this *Turtle) peek(direction Direction) rune {
	next := this.position.Move(direction)
	return this.spot(next)
}

type Point struct{ X, Y int }

func (this Point) Move(direction Direction) Point {
	return Point{
		X: this.X + direction.dx,
		Y: this.Y + direction.dy,
	}
}

type Direction struct{ dx, dy int }

func (this Direction) String() string {
	switch this {
	case North:
		return "N"
	case South:
		return "S"
	case East:
		return "E"
	case West:
		return "W"
	}
	return "?"
}

var (
	North = Direction{0, -1}
	South = Direction{0, 1}
	East  = Direction{1, 0}
	West  = Direction{-1, 0}

	Right = map[Direction]Direction{
		North: East,
		East:  South,
		South: West,
		West:  North,
	}
	Left = map[Direction]Direction{
		North: West,
		West:  South,
		South: East,
		East:  North,
	}
)
