package day22

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
)

type Virus struct {
	state    map[grid.Point]bool
	current  grid.Point
	facing   grid.Direction
	infected int
}

func NewVirus(cluster string) *Virus {
	lines := strings.Split(cluster, "\n")
	middle := float64(len(lines[0]) / 2)
	state := make(map[grid.Point]bool)
	for l := 0; l < len(lines); l++ {
		line := lines[l]
		for x, char := range line {
			y := len(lines) - 1 - l
			state[grid.NewPoint(float64(x), float64(y))] = char == '#'
		}
	}
	return &Virus{
		state:   state,
		current: grid.NewPoint(middle, middle),
		facing:  grid.Up,
	}
}

func (this *Virus) Current() grid.Point {
	return this.current
}

func (this *Virus) Move() {
	this.facing = turn[this.state[this.current]][this.facing]
	this.state[this.current] = !this.state[this.current]
	this.infected += infect[this.state[this.current]]
	this.current = this.current.Move(this.facing)
}

var turn = map[bool]map[grid.Direction]grid.Direction{
	true:  grid.Clockwise,
	false: grid.CounterClockwise,
}
var infect = map[bool]int{
	false: 0,
	true:  1,
}

func (this *Virus) Infected() int {
	return this.infected
}
