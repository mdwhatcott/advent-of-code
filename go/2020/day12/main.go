package advent

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

type Turtle struct {
	intgrid.Point
	intgrid.Direction
}

func Part1() interface{} {
	turtle := &Turtle{
		Point:     intgrid.NewPoint(0, 0),
		Direction: intgrid.Right,
	}

	for _, line := range util.InputLines() {
		direction, distance := line[0], parse.Int(line[1:])
		switch direction {
		case 'F':
			for ; distance > 0; distance-- {
				turtle.Point = turtle.Point.Move(turtle.Direction)
			}
		case 'L':
			for degrees := distance; degrees > 0; degrees -= 90 {
				turtle.Direction = turtle.TurnLeft()
			}
		case 'R':
			for degrees := distance; degrees > 0; degrees -= 90 {
				turtle.Direction = turtle.TurnRight()
			}
		case 'N':
			for ; distance > 0; distance-- {
				turtle.Point = turtle.Point.Move(intgrid.Up)
			}
		case 'S':
			for ; distance > 0; distance-- {
				turtle.Point = turtle.Point.Move(intgrid.Down)
			}
		case 'E':
			for ; distance > 0; distance-- {
				turtle.Point = turtle.Point.Move(intgrid.Right)
			}
		case 'W':
			for ; distance > 0; distance-- {
				turtle.Point = turtle.Point.Move(intgrid.Left)
			}
		}
	}
	return intgrid.ManhattanDistance(intgrid.Origin, turtle.Point)
}

func Part2() interface{} {
	return intgrid.ManhattanDistance(intgrid.Origin, part2(util.InputLines()))
}

func part2(lines []string) intgrid.Point {
	ship := &Turtle{
		Point:     intgrid.NewPoint(0, 0),
		Direction: intgrid.Right,
	}
	waypoint := intgrid.NewPoint(10, 1)

	for _, line := range lines {
		//log.Println("before:", ship.Point, "instruction:", line, "waypoint:", waypoint)
		direction, distance := line[0], parse.Int(line[1:])
		switch direction {
		case 'F':
			for ; distance > 0; distance-- {
				ship.Point = intgrid.NewPoint(ship.X()+waypoint.X(), ship.Y()+waypoint.Y())
			}
		case 'L':
			for ; distance > 0; distance -= 90 {
				waypoint = RotateLeft(waypoint)
			}
		case 'R':
			for ; distance > 0; distance -= 90 {
				waypoint = RotateRight(waypoint)
			}
		case 'N':
			for ; distance > 0; distance-- {
				waypoint = waypoint.Move(intgrid.Up)
			}
		case 'S':
			for ; distance > 0; distance-- {
				waypoint = waypoint.Move(intgrid.Down)
			}
		case 'E':
			for ; distance > 0; distance-- {
				waypoint = waypoint.Move(intgrid.Right)
			}
		case 'W':
			for ; distance > 0; distance-- {
				waypoint = waypoint.Move(intgrid.Left)
			}
		}
	}
	return ship.Point
}

func RotateRight(p intgrid.Point) intgrid.Point {
	return intgrid.NewPoint(p.Y(), -p.X())
}
func RotateLeft(p intgrid.Point) intgrid.Point {
	return intgrid.NewPoint(-p.Y(), p.X())
}
