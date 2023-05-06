package advent

import (
	"fmt"

	"github.com/mdwhatcott/go-collections/set"

	"advent/2019/intcode"
	"advent/lib/intgrid"
	"advent/lib/maths"
	"advent/lib/util"
)

func Part1() interface{} {
	graph := set.New[intgrid.Point](0)
	x := 0
	y := 0
	intcode.RunProgram(util.InputInts(","), nil, func(c int) {
		s := string(rune(c))
		if s == "#" {
			graph.Add(intgrid.NewPoint(x, y))
		}
		if s == "\n" {
			y++
			x = 0
		} else {
			x++
		}
		fmt.Print(s)
	})

	var alignments []int
	for p := range graph {
		if isCrossSection(graph, p) {
			alignments = append(alignments, p.X()*p.Y())
		}
	}
	return maths.Sum(alignments...)
}
func isCrossSection(graph set.Set[intgrid.Point], p intgrid.Point) bool {
	for _, neighbor := range intgrid.Neighbors4 {
		if !graph.Contains(p.Move(neighbor)) {
			return false
		}
	}
	return true
}

func Part2() interface{} {
	return nil
}
