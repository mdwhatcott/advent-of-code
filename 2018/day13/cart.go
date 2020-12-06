package day13

type Cart struct {
	Point
	Direction
	Intersections int
}

func (this *Cart) Travel() {
	this.Point = this.Move(this.Direction)
}

func (this *Cart) Orient(track rune) {
	switch track {
	case '+':
		this.turnIntersection()
	case '/', '\\':
		this.Direction = this.Direction.TurnSlash(track)
	}
}
func (this *Cart) turnIntersection() {
	switch this.Intersections % 3 {
	case 0:
		this.Direction = this.Direction.TurnLeft()
	case 1:
		this.Direction = this.Direction // Straight
	case 2:
		this.Direction = this.Direction.TurnRight()
	}
	this.Intersections++
}

type Carts []*Cart

func (this Carts) Len() int {
	return len(this)
}
func (this Carts) Less(i, j int) bool {
	return this[i].y < this[j].y || (this[i].y == this[j].y && this[i].x < this[j].x)
}
func (this Carts) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
