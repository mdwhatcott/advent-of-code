package day03

import (
	"advent/lib/grid"
	"advent/lib/util"
)

func Part1() int {
	input := util.InputInt()

	turtle := Start()
	for {
		turtle.Advance()
		if turtle.id == input {
			return int(grid.CityBlockDistance(turtle.at, grid.Origin))
		}
	}
}

func Part2() int {
	turtle := Start()
	for {
		turtle.Advance()
		written := turtle.Sum()
		if written > util.InputInt() {
			return written
		}
	}
	panic(":(")
}

type Turtle struct {
	at grid.Point
	id int

	distance  int
	traveled  int
	direction grid.Direction

	visited map[grid.Point]int
}

func Start() *Turtle {
	return &Turtle{
		at:        grid.Origin,
		direction: grid.Right,
		id:        1,
		distance:  2,
		traveled:  0,
		visited:   map[grid.Point]int{},
	}
}

func (this *Turtle) Sum() (sum int) {
	for _, d := range this.at.Neighbors8() {
		sum += this.visited[d]
	}
	if sum == 0 {
		sum++
	}
	return sum
}

func (this *Turtle) Advance() {
	at := this.at
	sum := this.Sum()
	this.visited[at] = sum

	if this.traveled == this.distance/2 {
		this.direction = grid.CounterClockwise[this.direction]
	} else if this.traveled == this.distance {
		this.traveled = 0
		this.distance += 2
		this.direction = grid.CounterClockwise[this.direction]
	}
	this.traveled++

	next := at.Move(this.direction)
	//log.Printf("Advancing from %v to %v (sum: %d)", at, next, sum)

	this.at = next
	this.id++
}
