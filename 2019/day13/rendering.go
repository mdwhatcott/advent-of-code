package advent

import (
	"fmt"
	"strings"

	"advent/lib/grid"
)

const (
	Empty  = 0
	Wall   = 1
	Block  = 2
	Paddle = 3
	Ball   = 4
)

var draw = map[int]string{
	Empty:  " ",
	Wall:   "#",
	Block:  "-",
	Paddle: "=",
	Ball:   "*",
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
