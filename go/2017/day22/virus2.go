package day22

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
)

type State int

const (
	Clean State = iota
	Weakened
	Infected
	Flagged
)

var States = map[State]State{
	Clean:    Weakened,
	Weakened: Infected,
	Infected: Flagged,
	Flagged:  Clean,
}

type Virus2 struct {
	state    map[grid.Point]State
	current  grid.Point
	facing   grid.Direction
	infected int
}

func NewVirus2(cluster string) *Virus2 {
	lines := strings.Split(cluster, "\n")
	middle := float64(len(lines[0]) / 2)
	state := make(map[grid.Point]State)
	for l := 0; l < len(lines); l++ {
		line := lines[l]
		for x, char := range line {
			y := len(lines) - 1 - l
			var s State
			if char == '#' {
				s = Infected
			}
			state[grid.NewPoint(float64(x), float64(y))] = s
		}
	}
	return &Virus2{
		state:   state,
		current: grid.NewPoint(middle, middle),
		facing:  grid.Up,
	}
}

func (this *Virus2) Current() grid.Point {
	return this.current
}

func (this *Virus2) Move() {
	current := this.state[this.current]
	switch current {
	case Clean:
		this.facing = grid.CounterClockwise[this.facing]
	case Infected:
		this.facing = grid.Clockwise[this.facing]
	case Flagged:
		this.facing = grid.Clockwise[grid.Clockwise[this.facing]]
	}
	this.state[this.current] = States[current]
	if this.state[this.current] == Infected {
		this.infected++
	}
	this.current = this.current.Move(this.facing)
}

func (this *Virus2) Infected() int {
	return this.infected
}
