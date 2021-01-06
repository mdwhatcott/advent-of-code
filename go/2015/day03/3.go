package advent

import "strconv"

type SantaTurtle struct {
	visited map[string]int
	x, y    int
}

func NewSantaTurtle(visits map[string]int) *SantaTurtle {
	turtle := &SantaTurtle{
		visited: visits,
	}
	turtle.recordVisit()
	return turtle
}

func (this *SantaTurtle) Move(direction rune) {
	switch direction {
	case '<':
		this.x--
	case '>':
		this.x++
	case 'v':
		this.y--
	case '^':
		this.y++
	}
	this.recordVisit()
}

func (this *SantaTurtle) recordVisit() {
	this.visited[strconv.Itoa(this.x)+"|"+strconv.Itoa(this.y)]++
}
