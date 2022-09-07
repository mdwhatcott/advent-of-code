package main

import (
	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/set"

	"advent/lib/intgrid"
	"advent/lib/util"
)

func BreadthFirstSearch(maze int, origin, target intgrid.Point) (distance, near int) {
	q := queue.New[Step](0)
	q.Enqueue(Step{Point: origin, Distance: 0})
	seen := set.From[intgrid.Point]()

	for !q.Empty() {
		current := q.Dequeue()
		if current.Point == target {
			return current.Distance, near
		}
		for _, adjacent := range neighbors(maze, current) {
			if !seen.Contains(adjacent.Point) {
				q.Enqueue(adjacent)
				seen.Add(adjacent.Point)
				if adjacent.Distance <= 50 {
					near++
				}
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
	bits := util.BinaryHammingWeight(sum)
	return bits%2 == 0
}
