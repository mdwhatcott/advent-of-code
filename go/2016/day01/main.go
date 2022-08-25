// http://adventofcode.com/2016/day/1
package main

import (
	"fmt"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func main() {
	turtle := NewTurtle()
	turtle.FollowAll(util.InputString())

	fmt.Println()

	fmt.Println("Ending position:")
	assert.So(t, nil, turtle.Position(), should.Equal, "-126,165")

	fmt.Println("Distance to ending position:")
	assert.So(t, nil, turtle.TaxiDistanceToEndingLocation(), should.Equal, 291)

	fmt.Println()
	fmt.Println("Position first visited twice:")
	assert.So(t, nil, turtle.PositionFirstVisitedTwice(), should.Equal, "16,143")

	fmt.Println("Distance to position first visited twice:")
	assert.So(t, nil, turtle.TaxiDistanceToLocationFirstVisitedTwice(), should.Equal, 159)

}
