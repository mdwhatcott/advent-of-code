package day12

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/astar"
	"advent/lib/intgrid"
	"advent/lib/util"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}
)

func TestDay12(t *testing.T) {
	should.So(t, Part1(sampleLines), should.Equal, 31)
	should.So(t, Part1(inputLines), should.Equal, 408)

	should.So(t, Part2(sampleLines), should.Equal, 29)
	should.So(t, Part2(inputLines), should.Equal, 399)
}

func Part2(lines []string) (min int) {
	min = 0xFFFFFFFF
	graph, _, target := parseGraph(lines)
	for origin, char := range graph {
		if char == int('a') {
			path, found := astar.SearchFrom(&Step{Graph: graph, Target: target, At: origin})
			if found && len(path)-1 < min {
				min = len(path) - 1
			}
		}
	}
	return min
}
func Part1(lines []string) int {
	graph, origin, target := parseGraph(lines)
	path, _ := astar.SearchFrom(&Step{Graph: graph, Target: target, At: origin})
	return len(path) - 1
}
func parseGraph(lines []string) (graph map[intgrid.Point]int, origin, target intgrid.Point) {
	graph = make(map[intgrid.Point]int)
	for y, line := range lines {
		for x, char := range line {
			point := intgrid.NewPoint(x, y)
			if char == 'S' {
				origin = point
				char = 'a'
			} else if char == 'E' {
				target = point
				char = 'z'
			}
			graph[point] = int(char)
		}
	}
	return graph, origin, target
}

type Step struct {
	Graph  map[intgrid.Point]int
	Target intgrid.Point
	At     intgrid.Point
}

func (this *Step) AdjacentPositions() (result []astar.Turtle) {
	for _, direction := range intgrid.Neighbors4 {
		neighbor := this.At.Move(direction)
		_, ok := this.Graph[neighbor]
		if !ok {
			continue
		}
		if this.Graph[neighbor] <= this.Graph[this.At]+1 {
			result = append(result, &Step{Graph: this.Graph, Target: this.Target, At: neighbor})
		}
	}

	return result
}
func (this *Step) EstimatedDistanceToTarget() float64 {
	return float64(intgrid.ManhattanDistance(this.At, this.Target))
}
func (this *Step) StepCost() float64 { return 1 }
func (this *Step) Hash() string      { return this.At.String() }
