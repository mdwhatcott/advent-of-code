package advent

import (
	year2018day10 "advent/2018/day10"
	"advent/lib/util"
)

func Part1() interface{} {
	robot := NewRobot(0, util.InputInts(","))
	for robot.Move() {
	}
	return len(robot.Hull)
}

func Part2() interface{} {
	robot := NewRobot(1, util.InputInts(","))
	for robot.Move() {
	}
	year2018day10.Render(robot.HullSlice())
	return "ABCLFUHJ"
}

/*

The Y coordinates are inverted, so it came out upside down:

 x  x xxx   xx  xxxx x     xx  x  x  xx
 x  x x  x x  x x    x    x  x x  x x  x
 xxxx x  x x    x    x    x  x x  x    x
 x  x xxx  x    x    xxx  x  x xxxx    x
 x  x x  x x  x x    x    x  x x  x    x
  xx  xxx   xx  x    xxxx x  x x  x   xx


Here it is, right side up:

  xx  xxx   xx  x    xxxx x  x x  x   xx
 x  x x  x x  x x    x    x  x x  x    x
 x  x xxx  x    x    xxx  x  x xxxx    x
 xxxx x  x x    x    x    x  x x  x    x
 x  x x  x x  x x    x    x  x x  x x  x
 x  x xxx   xx  xxxx x     xx  x  x  xx

Or, in actual ASCII text: ABCLFUHJ
 */
