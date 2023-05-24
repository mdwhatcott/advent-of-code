package advent

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
	"github.com/mdwhatcott/advent-of-code-intcode"
)

const (
	Black = 0
	White = 1

	Left  = 0
	Right = 1
)

type Robot struct {
	Position  grid.Point
	Direction grid.Direction
	Hull      map[grid.Point]int

	input     chan int
	output    chan int
	processor *intcode.Interpreter
}

func (this *Robot) in() int   { return <-this.input }
func (this *Robot) out(i int) { this.output <- i }
func (this *Robot) start() {
	this.processor.RunProgram()
	close(this.input)
	close(this.output)
}

func NewRobot(startingColor int, program []int) *Robot {
	start := grid.NewPoint(0, 0)
	robot := &Robot{
		Position:  start,
		Direction: grid.NewDirection(0, 1),
		Hull:      map[grid.Point]int{start: startingColor},
		input:     make(chan int),
		output:    make(chan int),
	}
	robot.processor = intcode.NewIntCodeInterpreter(program, robot.in, robot.out)
	go robot.start()
	return robot
}

func (this *Robot) Move() (active bool) {
	defer func() { // HACK: recover the send on closed channel
		if recover() != nil {
			active = false
		}
	}()

	this.input <- this.Hull[this.Position]

	this.Hull[this.Position] = <-this.output

	if <-this.output == Left {
		this.Direction = this.Direction.TurnLeft()
	} else {
		this.Direction = this.Direction.TurnRight()
	}

	this.Position = this.Position.Move(this.Direction)

	return true
}

func (this *Robot) HullSlice() (painted []grid.Point) {
	for position, color := range this.Hull {
		if color == 1 {
			painted = append(painted, position)
		}
	}
	return painted
}
