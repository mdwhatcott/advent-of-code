// http://adventofcode.com/2016/day/1
package main

import (
	"fmt"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func main() {
	turtle := NewTurtle()
	turtle.FollowAll(util.InputString())

	fmt.Println()

	fmt.Println("Ending position:")
	should.So(nil, turtle.Position(), should.Equal, "-126,165")

	fmt.Println("Distance to ending position:")
	should.So(nil, turtle.TaxiDistanceToEndingLocation(), should.Equal, 291)

	fmt.Println()
	fmt.Println("Position first visited twice:")
	should.So(nil, turtle.PositionFirstVisitedTwice(), should.Equal, "16,143")

	fmt.Println("Distance to position first visited twice:")
	should.So(nil, turtle.TaxiDistanceToLocationFirstVisitedTwice(), should.Equal, 159)

}
