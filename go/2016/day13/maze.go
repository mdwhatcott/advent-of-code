package main

import (
	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/set"

	"advent/lib/intgrid"
)

func BreadthFirstSearch(maze int, origin, target intgrid.Point) (distance, near int) {
	frontier := queue.New[Step](0)
	frontier.Enqueue(Step{Point: origin, Distance: 0})
	seen := set.From[intgrid.Point]()
	short := set.From[intgrid.Point]()

	for !frontier.Empty() {
		current := frontier.Dequeue()
		if current.Point == target {
			return current.Distance, short.Len()
		}
		if current.Distance <= 50 {
			short.Add(current.Point)
		}
		for _, adjacent := range neighbors(maze, current) {
			if !seen.Contains(adjacent.Point) {
				frontier.Enqueue(adjacent)
				seen.Add(adjacent.Point)
			}
		}
	}
	panic("Didn't find the destination.")
}

type Step struct {
	intgrid.Point
	Distance int
}

func neighbors(maze int, l Step) (results []Step) {
	distance := l.Distance + 1
	for _, neighbor := range l.Point.Neighbors4() {
		if isHallway(maze, neighbor) {
			results = append(results, Step{Point: neighbor, Distance: distance})
		}
	}
	return results
}
func isHallway(maze int, point intgrid.Point) bool {
	x, y := point.X(), point.Y()
	if x < 0 || y < 0 {
		return false
	}
	sum := x*x + 3*x + 2*x*y + y + y*y + maze
	bits := binaryHammingWeight(sum)
	return bits%2 == 0
}
func binaryHammingWeight(value int) (count int) { // See: https://en.wikipedia.org/wiki/Hamming_weight
	for ; value > 0; count++ {
		value &= value - 1
	}
	return count
}
