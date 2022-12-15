package day14

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/util"
)

var (
	sampleLines     = []string{"498,4 -> 498,6 -> 496,6", "503,4 -> 502,4 -> 502,9 -> 494,9"}
	inputLines      = util.InputLines()
	sampleWithFloor = append(sampleLines, "0,11 -> 1000,11")
	inputWithFloor  = append(util.InputLines(), "0,171 -> 1000,171")
)

func TestDay14(t *testing.T) {
	should.So(t, FillSand(sampleLines), should.Equal, 24)
	should.So(t, FillSand(inputLines), should.Equal, 745)

	should.So(t, FillSand(sampleWithFloor), should.Equal, 93)
	should.So(t, FillSand(inputWithFloor), should.Equal, 27551)
}

func FillSand(lines []string) int {
	floor := 0
	rocks := set.New[intgrid.Point](0)
	for _, line := range lines {
		var endpoints []intgrid.Point
		for _, endpoint := range strings.Split(line, " -> ") {
			xy := strings.Split(endpoint, ",")
			endpoints = append(endpoints, intgrid.NewPoint(util.ParseInt(xy[0]), util.ParseInt(xy[1])))
		}
		for x := 1; x < len(endpoints); x++ {
			from, to := endpoints[x-1], endpoints[x-0]
			direction := pointDirection(from, to)
			for from != to {
				if from.Y() > floor {
					floor = from.Y()
				}
				rocks.Add(from)
				from = from.Move(direction)
			}
			rocks.Add(to)
		}
	}

	sands := set.New[intgrid.Point](0)
	for {
		sand := intgrid.NewPoint(500, 0)
		final, isAtRest := fall(sand, floor, rocks, sands)
		if !isAtRest {
			return sands.Len()
		}
		if final == sand {
			sands.Add(final)
			return sands.Len()
		}
		sands.Add(final)
	}
}
func pointDirection(from, to intgrid.Point) intgrid.Direction {
	if from.X() < to.X() {
		return intgrid.Right
	} else if from.X() > to.X() {
		return intgrid.Left
	} else if from.Y() < to.Y() {
		return intgrid.Up
	} else {
		return intgrid.Down
	}
}
func fall(sand intgrid.Point, floor int, rocks, sands set.Set[intgrid.Point]) (final intgrid.Point, isAtRest bool) {
	for sand.Y() < floor {
		if tryFall(sand, intgrid.Up, rocks, sands) {
			sand = sand.Move(intgrid.Up)
		} else if tryFall(sand, intgrid.TopLeft, rocks, sands) {
			sand = sand.Move(intgrid.TopLeft)
		} else if tryFall(sand, intgrid.TopRight, rocks, sands) {
			sand = sand.Move(intgrid.TopRight)
		} else {
			return sand, true
		}
	}
	return sand, false
}
func tryFall(sand intgrid.Point, direction intgrid.Direction, rocks, sands set.Set[intgrid.Point]) bool {
	sand = sand.Move(direction)
	return !rocks.Contains(sand) && !sands.Contains(sand)
}
