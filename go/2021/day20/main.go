package day20

import (
	"strings"

	"advent/lib/intgrid"
)

type World map[intgrid.Point]bool

func parseAlgorithm(s string) (result []bool) {
	result = make([]bool, len(s))
	for i, c := range s {
		result[i] = c == '#'
	}
	return result
}
func parseImage(raw string) World {
	result := make(World)
	for y, line := range strings.Split(raw, "\n") {
		for x, char := range line {
			result[intgrid.NewPoint(x, y)] = char == '#'
		}
	}
	return result
}

func (this World) enhance(algorithm []bool, infinity bool) (next World) {
	next = make(World)
	minX, maxX, minY, maxY := this.bounds()
	for y := minY - 2; y < maxY+2; y++ {
		for x := minX - 2; x < maxX+2; x++ {
			p := intgrid.NewPoint(x, y)
			next[p] = algorithm[this.niner(p, infinity)]
		}
	}
	return next
}
func (this World) bounds() (minX, maxX, minY, maxY int) {
	for p := range this {
		x, y := p.X(), p.Y()
		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}
	}
	return minX, maxX, minY, maxY
}
func (this World) niner(point intgrid.Point, infinity bool) (n int) {
	// HACK: we pass in an alternating 'infinity' bool to account for
	// the fact that the upper left corner will, according to the algorithm
	// found in the real input, be 'on' every other round.
	for i, offset := range reverseOffsets {
		on, ok := this[point.Move(offset)]
		if ok && on || !ok && infinity /* <-- HACK */ {
			n |= 1 << i
		}
	}
	return n
}
func (this World) count() (result int) {
	for _, v := range this {
		if v {
			result++
		}
	}
	return result
}

var reverseOffsets = []intgrid.Direction{
	intgrid.NewDirection(1, 1),
	intgrid.NewDirection(0, 1),
	intgrid.NewDirection(-1, 1),
	intgrid.NewDirection(1, 0),
	intgrid.NewDirection(0, 0),
	intgrid.NewDirection(-1, 0),
	intgrid.NewDirection(1, -1),
	intgrid.NewDirection(0, -1),
	intgrid.NewDirection(-1, -1),
}
