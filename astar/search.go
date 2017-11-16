package astar

import "github.com/jupp0r/go-priority-queue"

type astarSearch struct {
	frontier   pq.PriorityQueue
	trail      map[Turtle]Turtle
	distanceTo map[string]float64
}

func newAStarSearch() *astarSearch {
	return &astarSearch{
		frontier:   pq.New(),
		trail:      make(map[Turtle]Turtle),
		distanceTo: make(map[string]float64),
	}
}

func (this *astarSearch) Search(start Turtle) (path []Turtle, found bool) {
	this.frontier.Insert(start, start.EstimatedDistanceToTarget())
	this.trail[start] = nil
	this.distanceTo[start.Hash()] = 0

	for this.frontier.Len() > 0 {
		pop, _ := this.frontier.Pop()
		current := pop.(Turtle)
		currentID := current.Hash()

		if current.EstimatedDistanceToTarget() == 0 {
			return this.pathFromStart(current, nil), true
		}

		for _, adjacent := range current.AdjacentPositions() {
			newCost := this.distanceTo[currentID] + 1

			adjacentID := adjacent.Hash()
			if cost, contains := this.distanceTo[adjacentID]; !contains || newCost < cost {
				this.frontier.Insert(adjacent, newCost+adjacent.EstimatedDistanceToTarget())
				this.trail[adjacent] = current
				this.distanceTo[adjacentID] = newCost
			}
		}
	}

	return path, false
}

func (this *astarSearch) pathFromStart(current Turtle, path []Turtle) []Turtle {
	if previous := this.trail[current]; previous == nil {
		return path
	} else {
		return this.pathFromStart(previous, append([]Turtle{previous}, path...))
	}
}
