package main

import (
	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/set"

	"advent/lib/intgrid"
	"advent/lib/util"
)

func BreadthFirstSearch(maze int, target intgrid.Point) (distance, near int) {
	seen := set.New[intgrid.Point](0)
	q := queue.New[Step](0)
	q.Enqueue(Step{Point: intgrid.NewPoint(1, 1), Distance: 0})

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
		if neighbor.X() >= 0 && neighbor.Y() >= 0 && isHallway(maze, neighbor) {
			results = append(results, Step{Point: neighbor, Distance: distance})
		}
	}
	return results
}
func isHallway(maze int, point intgrid.Point) bool {
	x, y := point.X(), point.Y()
	sum := x*x + 3*x + 2*x*y + y + y*y + maze
	bits := util.BinaryHammingWeight(sum)
	return bits%2 == 0
}
