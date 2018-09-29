package main

import "fmt"

func DistanceToDestination(queue *LocationQueue, maze, x, y int) int {
	origin := &Location{X: 1, Y: 1}
	queue.Enqueue(origin)

	for queue.Length() > 0 {
		current := queue.Dequeue()
		if current.IsDestination(x, y) {
			return current.Distance
		}
		for _, adjacent := range current.Adjacent() {
			if adjacent.IsHallway(maze) {
				queue.Enqueue(adjacent)
			}
		}
	}
	panic("Didn't find the destination.")
}

/**************************************************************************/

type LocationQueue struct {
	close     int
	seen      map[string]struct{}
	locations []*Location
}

func NewLocationQueue() *LocationQueue {
	return &LocationQueue{seen: make(map[string]struct{})}
}
func (this *LocationQueue) Length() int {
	return len(this.locations)
}

func (this *LocationQueue) Enqueue(location *Location) {
	if _, seen := this.seen[location.String()]; seen {
		return
	}
	this.locations = append(this.locations, location)
	this.seen[location.String()] = struct{}{}

	if location.Distance <= 50 {
		this.close++
	}
}

func (this *LocationQueue) Dequeue() *Location {
	l := this.locations[0]
	this.locations[0] = nil
	this.locations = this.locations[1:]
	return l
}

/**************************************************************************/

type Location struct {
	X, Y     int
	Previous *Location
	Distance int
}

func NewLocation(x, y, distance int, previous *Location) *Location {
	return &Location{X: x, Y: y, Distance: distance, Previous: previous}
}

func (this *Location) String() string {
	return fmt.Sprintf("%d,%d", this.X, this.Y)
}

func (this *Location) IsHallway(seed int) bool {
	return isHallway(seed, this.X, this.Y)
}
func (this *Location) IsDestination(x, y int) bool {
	return this.X == x && this.Y == y
}
func (this *Location) Adjacent() (a []*Location) {
	if this.Previous == nil {
		return append(a,
			NewLocation(this.X+1, this.Y, this.Distance+1, this),
			NewLocation(this.X-1, this.Y, this.Distance+1, this),
			NewLocation(this.X, this.Y+1, this.Distance+1, this),
			NewLocation(this.X, this.Y-1, this.Distance+1, this))
	}
	if this.X > 0 {
		a = append(a, NewLocation(this.X-1, this.Y, this.Distance+1, this))
	}
	if this.Y > 0 {
		a = append(a, NewLocation(this.X, this.Y-1, this.Distance+1, this))
	}
	a = append(a, NewLocation(this.X+1, this.Y, this.Distance+1, this))
	a = append(a, NewLocation(this.X, this.Y+1, this.Distance+1, this))
	return a
}

/**************************************************************************/

func isHallway(seed, x, y int) bool {
	return bits(sum(x, y)+seed)%2 == 0
}
func sum(x, y int) int {
	return x*x + 3*x + 2*x*y + y + y*y
}

// See: https://en.wikipedia.org/wiki/Hamming_weight
func bits(value int) int {
	var count int
	for count = 0; value > 0; count++ {
		value &= value - 1
	}
	return count
}
