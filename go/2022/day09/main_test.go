package day09

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/util"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
	sample2Lines = []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
)

func TestDay09(t *testing.T) {
	should.So(t, Part1(sampleLines), should.Equal, 13)
	should.So(t, Part1(inputLines), should.Equal, 5907)

	// Part 1, but with code for Part 2 (which should work...)
	should.So(t, Part2(sampleLines, 2), should.Equal, 13)
	should.So(t, Part2(inputLines, 2), should.Equal, 5907)

	should.So(t, Part2(sampleLines, 10), should.Equal, 1)
	should.So(t, Part2(sample2Lines, 10), should.Equal, 36)
	should.So(t, Part2(inputLines, 10), should.Equal, 2303)
}

var directions = map[string]intgrid.Direction{
	"L": intgrid.Left,
	"R": intgrid.Right,
	"U": intgrid.Up,
	"D": intgrid.Down,
}

func Part2(lines []string, knots int) int {
	visited := set.From[intgrid.Point]()
	chain := make([]intgrid.Point, knots)
	for _, line := range lines {
		fields := strings.Fields(line)
		direction := directions[fields[0]]
		steps := util.ParseInt(fields[1])

		for ; steps > 0; steps-- {
			newChain := make([]intgrid.Point, knots)
			newChain[0] = chain[0].Move(direction)
			for x := 1; x < len(chain); x++ {
				tail := chain[x]
				newHead := newChain[x-1]
				if tooFar(tail, newHead) {
					newChain[x] = tail.Move(follow(tail, newHead))
				} else {
					newChain[x] = chain[x]
				}
			}
			chain = newChain
			visited.Add(chain[len(chain)-1])
		}
	}
	return visited.Len()
}

func Part1(lines []string) int {
	head := intgrid.NewPoint(0, 0)
	tail := intgrid.NewPoint(0, 0)
	visited := set.From[intgrid.Point](tail)
	for _, line := range lines {
		fields := strings.Fields(line)
		direction := directions[fields[0]]
		distance := util.ParseInt(fields[1])
		for ; distance > 0; distance-- {
			newHead := head.Move(direction)
			if tooFar(tail, newHead) {
				tail = tail.Move(follow(tail, newHead))
			}
			head = newHead
			visited.Add(tail)
		}
	}
	return visited.Len()
}

func tooFar(tail, head intgrid.Point) bool {
	if tail == head {
		return false
	}
	distance := intgrid.ManhattanDistance(tail, head)
	if distance == 1 {
		return false
	}
	if distance > 2 {
		return true
	}
	return util.Abs(tail.X()-head.X()) != 1 || util.Abs(tail.Y()-head.Y()) != 1
}

func one(n int) int {
	if n == 0 {
		return 0
	}
	if n > 0 {
		return 1
	}
	return -1
}
func follow(from, to intgrid.Point) intgrid.Direction {
	if from.X() == to.X() || from.Y() == to.Y() {
		return intgrid.NewDirection(one(to.X()-from.X()), one(to.Y()-from.Y()))
	}
	dist := intgrid.ManhattanDistance(from, to)
	for _, diag := range diags {
		if intgrid.ManhattanDistance(from.Move(diag), to) < dist {
			return diag
		}
	}
	panic("NOPE" + from.String() + to.String())
}

var diags = []intgrid.Direction{
	intgrid.NewDirection(1, 1),
	intgrid.NewDirection(-1, 1),
	intgrid.NewDirection(-1, -1),
	intgrid.NewDirection(1, -1),
}
