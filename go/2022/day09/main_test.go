package day09

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/maths"
	"advent/lib/util"
)

var (
	inputLines   = util.InputLines()
	sampleLines  = []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}
	sample2Lines = []string{"R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"}
)

func TestDay09(t *testing.T) {
	should.So(t, Simulate(sampleLines, 2), should.Equal, 13)
	should.So(t, Simulate(inputLines, 2), should.Equal, 5907)

	should.So(t, Simulate(sampleLines, 10), should.Equal, 1)
	should.So(t, Simulate(sample2Lines, 10), should.Equal, 36)
	should.So(t, Simulate(inputLines, 10), should.Equal, 2303)
}

var directions = map[string]intgrid.Direction{
	"L": intgrid.Left,
	"R": intgrid.Right,
	"U": intgrid.Up,
	"D": intgrid.Down,
}

func Simulate(moves []string, knots int) int {
	visited := set.From[intgrid.Point]()
	chain := make([]intgrid.Point, knots)
	for _, move := range moves {
		fields := strings.Fields(move)
		direction := directions[fields[0]]
		steps := util.ParseInt(fields[1])

		for ; steps > 0; steps-- {
			chain[0] = chain[0].Move(direction)
			drag(chain)
			visited.Add(chain[len(chain)-1])
		}
	}
	return visited.Len()
}
func drag(chain []intgrid.Point) {
	for x := 1; x < len(chain); x++ {
		tail, head := chain[x], chain[x-1]
		if !stretchedTooFar(tail, head) {
			return
		}
		chain[x] = tail.Move(follow(tail, head))
	}
}
func stretchedTooFar(tail, head intgrid.Point) bool {
	return intgrid.ManhattanDistance(tail, head) > 1 &&
		(maths.Abs(tail.X()-head.X()) > 1 ||
			maths.Abs(tail.Y()-head.Y()) > 1)
}
func follow(from, to intgrid.Point) intgrid.Direction {
	return intgrid.NewDirection(
		zeroOrOne(to.X()-from.X()),
		zeroOrOne(to.Y()-from.Y()))
}
func zeroOrOne(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}
