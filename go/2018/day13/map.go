package day13

import (
	"bytes"
	"sort"
	"strings"
)

type Map struct {
	Signals chan Point
	lines   []string
	carts   Carts
}

func NewMap(input string) (result *Map) {
	result = &Map{Signals: make(chan Point)}

	for y, line := range strings.Split(input, "\n") {
		var current string
		for x, char := range line {

			switch char {
			case ' ', '/', '\\', '+', '|', '-':
				current += string(char)
			case 'v', '^':
				current += "|"
			case '<', '>':
				current += "-"
			}

			direction, ok := cartDirections[char]
			if ok {
				cart := &Cart{Point: NewPoint(x, y), Direction: direction}
				result.carts = append(result.carts, cart)
			}
		}
		result.lines = append(result.lines, current)
	}

	return result
}

func (this *Map) String() string {
	lines := bytes.Split([]byte(strings.Join(this.lines, "\n")), []byte("\n"))
	for _, cart := range this.carts {
		switch lines[cart.y][cart.x] {
		case 'v', '^', '<', '>':
			lines[cart.y][cart.x] = 'X'
		default:
			lines[cart.y][cart.x] = cartPositions[cart.Direction]
		}
	}
	return string(bytes.Join(lines, []byte("\n")))
}

func (this *Map) At(point Point) rune {
	return rune(this.lines[point.y][point.x])
}

func (this *Map) Tick() bool {
	sort.Sort(this.carts)

	for c, cart := range this.carts {
		cart.Travel()
		cart.Orient(this.At(cart.Point))

		collision := this.CheckCollision(c)
		if collision != Origin {
			this.Signals <- collision
		}
	}

	if len(this.carts) == 1 {
		this.Signals <- this.carts[0].Point
		close(this.Signals)
		return false
	}
	return true
}

func (this *Map) CheckCollision(current int) Point {
	compare := this.carts[current]
	for c, cart := range this.carts {
		if c == current {
			continue
		}
		if compare.Point == cart.Point {

		}
	}
	return Point{}
}

func (this *Map) LastCart() Point {
	if len(this.carts) > 1 {
		return Origin
	}
	if len(this.carts) == 0 {
		panic("ran out of carts")
	}
	return this.carts[0].Point
}
