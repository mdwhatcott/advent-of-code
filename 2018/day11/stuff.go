package day11

import (
	"fmt"
	"strings"
)

type Grid struct {
	serial int
	grid   map[string]int // map[x,y]power
}

func NewGrid(serial int) *Grid {
	return &Grid{
		serial: serial,
		grid:   make(map[string]int),
	}
}

func (this *Grid) String() string {
	b := new(strings.Builder)
	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			fmt.Fprintf(b, "%3d", this.PowerAt(x, y))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (this *Grid) MaxPowerXY(size int) (xy string) {
	var max, maxX, maxY int
	for y := 0; y < 300-2; y++ {
		for x := 0; x < 300-2; x++ {
			if power := this.PowerAtSquare(x, y, size); power > max {
				max, maxX, maxY = power, x, y
			}
		}
	}
	return fmt.Sprintf("%d,%d", maxX, maxY)
}

func (this *Grid) PowerAt(x, y int) (power int) {
	xy := fmt.Sprintf("%d,%d", x, y)
	if cached, found := this.grid[xy]; found {
		return cached
	}
	rack := x + 10
	power = rack * y
	power += this.serial
	power *= rack
	power = hundreds(power)
	power -= 5
	this.grid[xy] = power
	return power
}

func (this *Grid) MaxPowerXYSize() string {
	var max, maxSize, maxX, maxY int
	for size := 300; size >= 1; size-- {
		for y := 0; y < 300-size+1; y++ {
			for x := 0; x < 300-size+1; x++ {
				if power := this.PowerAtSquare(x, y, size); power > max {
					max, maxX, maxY = power, x, y
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d,%d", maxX, maxY, maxSize)
}

func (this *Grid) PowerAtSquare(x, y, size int) (power int) {
	for yy := y; yy < y+size; yy++ {
		for xx := x; xx < x+size; xx++ {
			power += this.PowerAt(xx, yy)
			//fmt.Println(x, y, xx, yy, power)
		}
	}
	return power
}

func hundreds(value int) (end int) {
	for value >= 1000 {
		value -= 1000
	}
	for value >= 100 {
		end++
		value -= 100
	}
	return end
}
