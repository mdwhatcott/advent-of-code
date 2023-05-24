package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
)

type Turtle struct {
	x  int
	y  int
	dx int
	dy int

	x2 int
	y2 int

	twice   bool
	visited map[string]int
}

func NewTurtle() *Turtle {
	return &Turtle{dy: 1, visited: make(map[string]int)}
}

func (this *Turtle) FollowAll(instructions string) {
	for _, instruction := range strings.Split(instructions, ", ") {
		this.Follow(instruction)
	}
}

func (this *Turtle) Follow(instruction string) {
	this.turn(instruction[0])
	this.walk(parse.Int(instruction[1:]))
}

func (this *Turtle) turn(direction byte) {
	switch direction {
	case 'R':
		this.turnRight()
	case 'L':
		this.turnLeft()
	}
}

func (this *Turtle) turnRight() {
	if this.dy != 0 {
		this.dx = this.dy
		this.dy = 0
	} else if this.dx != 0 {
		this.dy = -this.dx
		this.dx = 0
	}
}

func (this *Turtle) turnLeft() {
	this.turnRight()
	this.turnRight()
	this.turnRight()
}

func (this *Turtle) walk(steps int) {
	for x := 0; x < steps; x++ {
		this.x += this.dx
		this.y += this.dy

		this.logCurrentPosition()
	}
}

func (this *Turtle) logCurrentPosition() {
	position := this.Position()
	this.visited[position]++
	if this.twice {
		return
	}
	if this.visited[position] != 2 {
		return
	}
	this.x2 = this.x
	this.y2 = this.y
	this.twice = true
}

func (this *Turtle) Position() string {
	return fmt.Sprintf("%d,%d", this.x, this.y)
}

func (this *Turtle) PositionFirstVisitedTwice() string {
	if !this.twice {
		return ""
	}
	return fmt.Sprintf("%d,%d", this.x2, this.y2)
}

func (this *Turtle) TaxiDistanceToEndingLocation() int {
	return abs(this.x) + abs(this.y)
}

func (this *Turtle) TaxiDistanceToLocationFirstVisitedTwice() int {
	return abs(this.x2) + abs(this.y2)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
