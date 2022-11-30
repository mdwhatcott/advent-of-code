package advent

import (
	"log"

	"github.com/mdwhatcott/go-collections/set"

	"advent/2019/intcode"
	"advent/lib/astar"
	"advent/lib/intgrid"
	"advent/lib/util"
)

func init() {
	log.SetFlags(0)
}

var directionals = map[intgrid.Direction]int{
	intgrid.Up:    1,
	intgrid.Down:  2,
	intgrid.Left:  3,
	intgrid.Right: 4,
}

type Explorer struct {
	maze               set.Set[intgrid.Point]
	at                 intgrid.Point
	facing             intgrid.Direction
	target             intgrid.Point
	currentLength      int
	currentLengthCount int
}

func NewExplorer() *Explorer {
	return &Explorer{
		maze:   set.New[intgrid.Point](0),
		at:     intgrid.Origin,
		facing: intgrid.Up,
	}
}

func (this *Explorer) input() int {
	return directionals[this.facing]
}

const (
	ReportHitWall = 0
	ReportOxygen  = 2
)

func (this *Explorer) output(report int) {
	// If we hit a wall, turn right, otherwise move forward and turn left.
	// Credit for this strategy: https://www.reddit.com/r/adventofcode/comments/eaurfo/comment/fbbh2d3/

	if report == ReportHitWall {
		this.facing = intgrid.Clockwise[this.facing]
		return
	}

	log.Println("at:", this.at, "facing:", this.facing, len(this.maze))

	this.at = this.at.Move(this.facing)
	this.facing = intgrid.CounterClockwise[this.facing]
	this.maze.Add(this.at)

	if this.exploredAll() {
		log.Panicln("exploration complete") // HACK
	}

	if report == ReportOxygen {
		log.Println("TARGET ACQUIRED!!!", this.at)
		this.target = this.at
	}
}

func (this *Explorer) exploredAll() bool {
	if this.target == intgrid.Origin {
		return false
	}
	if len(this.maze) == this.currentLength {
		this.currentLengthCount++
	} else {
		this.currentLength = len(this.maze)
		this.currentLengthCount = 0
	}

	return this.currentLengthCount >= 1000
}

func (this *Explorer) CalculatePathDistanceToTarget() int {
	path, found := astar.NewGridTurtle(this.maze, intgrid.Origin, this.target).Search()
	if !found {
		return -1
	}
	return len(path) - 1
}

func (this *Explorer) CalculatePathDistanceToPointFurthestFromOxygen() (result int) {
	return result
}

func Part1() (result int) {
	finder := NewExplorer()
	defer func() {
		r := recover()
		if r != nil {
			log.Println("Pathfinding to target...")
			result = finder.CalculatePathDistanceToTarget()
			log.Println("Shortest path to target:", result)
		}
	}()
	intcode.RunProgram(util.InputInts(","), finder.input, finder.output)
	return -1 // shouldn't happen
}

func Part2() (result int) {
	finder := NewExplorer()
	defer func() {
		r := recover()
		if r != nil {
			log.Println("Pathfinding to target...")
			result = finder.CalculatePathDistanceToPointFurthestFromOxygen()
			log.Println("Shortest path to target:", result)
		}
	}()
	intcode.RunProgram(util.InputInts(","), finder.input, finder.output)
	return -1 // shouldn't happen
}
