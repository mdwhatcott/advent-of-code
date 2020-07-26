package advent

import (
	"fmt"
	"strings"

	"advent/2019/intcode"
	"advent/lib/grid"
	"advent/lib/util"
)

func Part1() interface{} {
	machine := intcode.NewHarness(util.InputInts(","))
	machine.Run()
	outputs := machine.Outputs()
	pixels := make(map[grid.Point]int)
	for x := 0; x < len(outputs); x += 3 {
		pixels[grid.NewPoint(float64(outputs[x]), float64(outputs[x+1]))] = outputs[x+2]
	}
	return strings.Count(Render(pixels), "-")
}

func Part2() interface{} {
	console := NewGameConsole(util.InputInts(","))
	console.InsertQuarters(2)
	return console.Play()
}

func Render(points map[grid.Point]int) string {
	builder := new(strings.Builder)
	minX, maxX, minY, maxY := drawBoundingBox(points)
	width := maxX - minX
	height := maxY - minY

	screen := make([][]string, int(height+1))
	for y := 0; y < len(screen); y++ {
		screen[y] = make([]string, int(width+1))
	}
	for point, pixel := range points {
		y := int(point.Y() - minY)
		x := int(point.X() - minX)
		screen[y][x] = draw[pixel]
	}

	for _, line := range screen {
		for _, char := range line {
			fmt.Fprint(builder, char)
		}
		fmt.Fprintln(builder)
	}
	return builder.String()
}

func drawBoundingBox(points map[grid.Point]int) (minX, maxX, minY, maxY float64) {
	for point := range points {
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
