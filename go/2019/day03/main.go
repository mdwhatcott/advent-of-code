package advent

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	input := util.InputLines()
	path1 := Follow(strings.Split(input[0], ","), grid.Origin)
	path2 := Follow(strings.Split(input[1], ","), grid.Origin)
	intersections := FindIntersections(path1, path2)
	closest := FindClosest(intersections, grid.Origin)
	return int(grid.ManhattanDistance(grid.Origin, closest))
}

func Follow(instructions []string, current grid.Point) (path map[grid.Point]int) {
	path = make(map[grid.Point]int)
	step := 0
	for _, instruction := range instructions {
		direction := instruction[0:1]
		steps := parse.Int(instruction[1:])
		for x := 0; x < steps; x++ {
			old := path[current]
			if old == 0 {
				path[current] = step
			}
			current = current.Move(directions[direction])
			step++
		}
	}
	return path
}

var directions = map[string]grid.Direction{
	"U": grid.Up,
	"D": grid.Down,
	"R": grid.Right,
	"L": grid.Left,
}

func FindIntersections(path1, path2 map[grid.Point]int) (intersections map[grid.Point]int) {
	intersections = make(map[grid.Point]int)
	for point, steps1 := range path1 {
		if steps2, found := path2[point]; found {
			intersections[point] = steps1 + steps2
		}
	}
	return intersections
}

func FindClosest(intersections map[grid.Point]int, origin grid.Point) grid.Point {
	min := 10000.0
	var minPoint grid.Point
	for point := range intersections {
		distance := grid.ManhattanDistance(origin, point)
		if distance == 0 {
			continue
		}
		if distance < min {
			min = distance
			minPoint = point
		}
	}
	return minPoint
}

func Part2() interface{} {
	input := util.InputLines()
	origin := grid.NewPoint(0, 0)
	path1 := Follow(strings.Split(input[0], ","), origin)
	path2 := Follow(strings.Split(input[1], ","), origin)
	intersections := FindIntersections(path1, path2)
	return FindShortest(intersections)
}

func FindShortest(intersections map[grid.Point]int) (shortest int) {
	shortest = 0xFFFFFFFF
	for _, steps := range intersections {
		if steps > 0 && steps < shortest {
			shortest = steps
		}
	}
	return shortest
}
