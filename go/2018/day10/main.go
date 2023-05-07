package day10

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/grid"
	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

const ExpectedSecondsToReachMessage = 10813

func Part1() interface{} {
	input := util.InputString()
	input = util.RemoveAll(input, "position", "velocity", "=", "<", ">", ",")
	lines := strings.Split(input, "\n")
	var points []grid.Point
	var directions []grid.Direction
	for _, line := range lines {
		fields := strings.Fields(line)
		x, y, dx, dy := util.Slice[int](parse.Ints(fields)).Unpack4()
		points = append(points, grid.NewPoint(float64(x), float64(y)))
		directions = append(directions, grid.NewDirection(float64(dx), float64(dy)))
	}

	for reps := 1; ; reps++ {
		for x := 0; x < len(points); x++ {
			points[x] = points[x].Move(directions[x])
		}
		if reps == ExpectedSecondsToReachMessage { // magic number attained by trial and error
			Render(points)
			break
		}
	}
	return "ERCXLAJL"
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

func Part2() interface{} {
	return ExpectedSecondsToReachMessage
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
