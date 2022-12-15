package day08

import (
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/util"
)

var sampleLines = []string{"30373", "25512", "65332", "33549", "35390"}

func TestDay08(t *testing.T) {
	should.So(t, CountVisibleFromEdge(sampleLines), should.Equal, 21)
	should.So(t, CountVisibleFromEdge(util.InputLines()), should.Equal, 1835)

	should.So(t, HighestScenicScore(sampleLines), should.Equal, 8)
	should.So(t, HighestScenicScore(util.InputLines()), should.Equal, 263670)
}
func CountVisibleFromEdge(lines []string) (result int) {
	grid := makeGrid(lines)
	visible := set.New[intgrid.Point](0)
	for x := 0; x < len(lines); x++ { // looking from top
		for at, y := -1, 0; y < len(lines); y++ {
			at = look(intgrid.NewPoint(x, y), grid, visible, at)
		}
	}
	for x := 0; x < len(lines); x++ { // looking from bottom
		for at, y := -1, len(lines)-1; y >= 0; y-- {
			at = look(intgrid.NewPoint(x, y), grid, visible, at)
		}
	}
	for y := 0; y < len(lines); y++ { // looking from left
		for at, x := -1, 0; x < len(lines); x++ {
			at = look(intgrid.NewPoint(x, y), grid, visible, at)
		}
	}
	for y := 0; y < len(lines); y++ { // looking from right
		for at, x := -1, len(lines)-1; x >= 0; x-- {
			at = look(intgrid.NewPoint(x, y), grid, visible, at)
		}
	}
	return visible.Len()
}
func look(p intgrid.Point, grid map[intgrid.Point]int, visible set.Set[intgrid.Point], at int) int {
	if grid[p] > at {
		visible.Add(p)
		at = grid[p]
	}
	return at
}
func HighestScenicScore(lines []string) (max int) {
	grid := makeGrid(lines)
	for p, h := range grid {
		score := lineOfSight(grid, p, h, intgrid.Up) *
			lineOfSight(grid, p, h, intgrid.Down) *
			lineOfSight(grid, p, h, intgrid.Left) *
			lineOfSight(grid, p, h, intgrid.Right)
		if score > max {
			max = score
		}
	}
	return max
}
func lineOfSight(grid map[intgrid.Point]int, p intgrid.Point, h int, direction intgrid.Direction) (score int) {
	for {
		p = p.Move(direction)
		if !contains(grid, p) {
			break
		}
		score++
		if h2 := grid[p]; h2 >= h {
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
