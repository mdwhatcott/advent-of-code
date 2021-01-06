// http://adventofcode.com/2016/day/1
package main

import (
	"fmt"

	"advent/lib/util"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	turtle := NewTurtle()
	turtle.FollowAll(util.InputString())

	fmt.Println()

	fmt.Println("Ending position:",
		assert.So(turtle.Position(), should.Equal, "-126,165").Fatal())
	fmt.Println("Distance to ending position:",
		assert.So(turtle.TaxiDistanceToEndingLocation(), should.Equal, 291).Fatal())

	fmt.Println()

	fmt.Println("Position first visited twice:",
		assert.So(turtle.PositionFirstVisitedTwice(), should.Equal, "16,143").Fatal())
	fmt.Println("Distance to position first visited twice:",
		assert.So(turtle.TaxiDistanceToLocationFirstVisitedTwice(), should.Equal, 159).Fatal())
}
