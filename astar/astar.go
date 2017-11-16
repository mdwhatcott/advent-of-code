package astar

func SearchFrom(start Turtle) (path []Turtle, found bool) {
	return newAStarSearch().Search(start)
}

type Turtle interface {
	EstimatedDistanceToTarget() float64
	AdjacentPositions() []Turtle
	Hash() string
}
