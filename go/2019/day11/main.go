package advent

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/go/lib/grid"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
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
	Render(robot.HullSlice())
	return "ABCLFUHJ"
}

func Render(points []grid.Point) {
	minX, maxX, minY, maxY := drawBoundingBox(points)
	width := maxX - minX
	height := maxY - minY
	if width > 300 || height > 300 {
		return
	}

	fmt.Println(minX, maxX, minY, maxY)
	screen := make([][]string, int(height)+20)
	for y := 0; y < len(screen); y++ {
		screen[y] = make([]string, int(width+20))
		for x := 0; x < len(screen[y]); x++ {
			screen[y][x] = " "
		}
	}
	for _, point := range points {
		y := int(point.Y() - minY)
		x := int(point.X() - minX)
		screen[y][x] = "x"
	}

	for _, line := range screen {
		for _, char := range line {
			fmt.Print(char)
		}
		fmt.Println()
	}
}
func drawBoundingBox(points []grid.Point) (minX, maxX, minY, maxY float64) {
	for _, point := range points {
		x := point.X()
		y := point.Y()
		if x > maxX {
			maxX = x
		} else if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		} else if y < minY {
			minY = y
		}
	}
	return minX, maxX, minY, maxY
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
