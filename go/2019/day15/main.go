package advent

import (
	"log"

	"github.com/mdwhatcott/go-collections/queue"
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

	//log.Println("at:", this.at, "facing:", this.facing, len(this.maze))

	this.at = this.at.Move(this.facing)
	this.facing = intgrid.CounterClockwise[this.facing]
	this.maze.Add(this.at)

	if this.exploredAll() {
		log.Panicln("exploration complete, maze size:", len(this.maze)) // HACK
	}

	if report == ReportOxygen && this.target == intgrid.Origin {
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

func (this *Explorer) AStarDistanceToTarget() int {
	path, found := astar.NewGridTurtle(this.maze, intgrid.Origin, this.target).Search()
	if !found {
		return -1
	}
	return len(path) - 1
}

func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func (this *Explorer) BFSDistanceToPointFurthestFromOxygen() (result int) {
	type Step struct {
		at       intgrid.Point
		distance int
	}
	seen := make(map[intgrid.Point]int)
	frontier := queue.New[Step](0)
	frontier.Enqueue(Step{at: this.target})
	for !frontier.Empty() {
		step := frontier.Dequeue()
		if !Contains(seen, step.at) {
			seen[step.at] = step.distance
		}
		for _, d := range intgrid.Neighbors4 {
			move := step.at.Move(d)
			if this.maze.Contains(move) && !Contains(seen, move) {
				frontier.Enqueue(Step{at: move, distance: step.distance + 1})
			}
		}
	}
	for _, distance := range seen {
		if distance > result {
			result = distance
		}
	}
	return result
}

func Part1() (result int) {
	finder := NewExplorer()
	defer func() {
		r := recover()
		if r != nil {
			log.Println("Pathfinding to target...")
			result = finder.AStarDistanceToTarget()
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
			log.Println("Measuring distance to point furthest from target...")
			result = finder.BFSDistanceToPointFurthestFromOxygen()
			log.Println("Distance to point furthest from target:", result)
		}
	}()
	intcode.RunProgram(util.InputInts(","), finder.input, finder.output)
	return -1 // shouldn't happen
}
