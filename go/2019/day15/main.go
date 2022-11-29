package advent

import (
	"log"

	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/set"

	"advent/2019/intcode"
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

type PathStep struct {
	Distance  int
	NowAt     intgrid.Point
	Direction intgrid.Direction
	LastAt    *PathStep
}
type PathFinder struct {
	seen     set.Set[intgrid.Point]
	frontier *queue.Queue[*PathStep]
	current  *PathStep
	target   *PathStep
}

func NewPathFinder() *PathFinder {
	s := queue.New[*PathStep](0)
	origin := &PathStep{NowAt: intgrid.Origin}
	for _, d := range intgrid.Neighbors4 {
		s.Enqueue(&PathStep{NowAt: intgrid.Origin.Move(d), Direction: d, LastAt: origin, Distance: 1})
	}
	return &PathFinder{
		seen:     set.From[intgrid.Point](intgrid.Origin),
		frontier: s,
	}
}

func (this *PathFinder) input() int {
	this.current = this.frontier.Dequeue()
	i := directionals[this.current.Direction]
	log.Println(this.current.Distance, "providing input:", i, this.current.NowAt)
	return i
}

func (this *PathFinder) output(report int) {
	this.seen.Add(this.current.NowAt)

	switch report {
	case 0:
		// hit a wall...hope there's something in frontier...
	case 1:
		for _, d := range intgrid.Neighbors4 {
			destination := this.current.NowAt.Move(d)
			if !this.seen.Contains(destination) {
				this.frontier.Enqueue(&PathStep{
					Direction: d,
					NowAt:     destination,
					LastAt:    this.current,
					Distance:  this.current.Distance + 1,
				})
			}
		}
	case 2:
		if this.target != nil {
			this.target = this.current
		}
	}
}

func Part1() (result int) {
	finder := NewPathFinder()
	intcode.RunProgram(util.InputInts(","), finder.input, finder.output)
	return finder.target.Distance
}

func Part2() interface{} {
	return nil
}
