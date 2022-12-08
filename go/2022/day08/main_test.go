package day08

import (
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/util"
)

var sampleLines = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestDay08(t *testing.T) {
	should.So(t, CountVisibleFromEdge(sampleLines), should.Equal, 21)
	should.So(t, CountVisibleFromEdge(util.InputLines()), should.Equal, 1835)

	should.So(t, HighestScenicScore(sampleLines), should.Equal, 8)
	should.So(t, HighestScenicScore(util.InputLines()), should.Equal, 0)
}
func CountVisibleFromEdge(lines []string) (result int) {
	grid := makeGrid(lines)
	visible := set.New[intgrid.Point](0)
	// looking from top
	for x := 0; x < len(lines); x++ {
		at := -1
		for y := 0; y < len(lines); y++ {
			p := intgrid.NewPoint(x, y)
			compare := grid[p]
			if compare > at {
				visible.Add(p)
				at = compare
			}
		}
	}
	// looking from bottom
	for x := 0; x < len(lines); x++ {
		at := -1
		for y := len(lines) - 1; y >= 0; y-- {
			p := intgrid.NewPoint(x, y)
			compare := grid[p]
			if compare > at {
				visible.Add(p)
				at = compare
			}
		}
	}
	// looking from left
	for y := 0; y < len(lines); y++ {
		at := -1
		for x := 0; x < len(lines); x++ {
			p := intgrid.NewPoint(x, y)
			compare := grid[p]
			if compare > at {
				visible.Add(p)
				at = compare
			}
		}
	}
	// looking from right
	for y := 0; y < len(lines); y++ {
		at := -1
		for x := len(lines) - 1; x >= 0; x-- {
			p := intgrid.NewPoint(x, y)
			compare := grid[p]
			if compare > at {
				visible.Add(p)
				at = compare
			}
		}
	}
	return visible.Len()
}
func HighestScenicScore(lines []string) (max int) {
	grid := makeGrid(lines)
	for p, h := range grid {
		score := look(grid, p, h, intgrid.Up) *
			look(grid, p, h, intgrid.Down) *
			look(grid, p, h, intgrid.Left) *
			look(grid, p, h, intgrid.Right)
		if score > max {
			max = score
		}
	}
	return max
}
func look(grid map[intgrid.Point]int, p intgrid.Point, h int, direction intgrid.Direction) (score int) {
	for {
		p = p.Move(direction)
		if !contains(grid, p) {
			break
		}
		h2 := grid[p]
		score++
		if h2 >= h {
			break
		}
	}
	return score
}
func makeGrid(lines []string) map[intgrid.Point]int {
	grid := make(map[intgrid.Point]int)
	for y, line := range lines {
		for x, char := range line {
			grid[intgrid.NewPoint(x, y)] = int(char - '0')
		}
	}
	return grid
}
func contains(grid map[intgrid.Point]int, p intgrid.Point) bool {
	_, ok := grid[p]
	return ok
}
