package advent

import (
	"log"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	var grid [][]string
	lines := util.InputLines()
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	field := NewField(grid, 4, false)

	for {
		field.Scan()
		if field.Update() == 0 {
			break
		}
	}
	return field.CountOccupied()
}

func Part2() interface{} {
	log.SetFlags(log.Lshortfile)
	var grid [][]string
	lines := util.InputLines()
	for _, line := range lines {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	field := NewField(grid, 5, true)

	for {
		field.Scan()
		if field.Update() == 0 {
			break
		}
	}
	return field.CountOccupied()
}

type Cell struct {
	intgrid.Point
	State string
	Next  string
}

type Field struct {
	width       int
	height      int
	cells       map[intgrid.Point]*Cell
	overcrowded int
	farsighted  bool
}

func NewField(field [][]string, overcrowded int, farsighted bool) *Field {
	cells := make(map[intgrid.Point]*Cell)
	for r, row := range field {
		for c, col := range row {
			cell := &Cell{Point: intgrid.NewPoint(c, r), State: col}
			cells[cell.Point] = cell
		}
	}

	return &Field{
		width:       len(field[0]),
		height:      len(field),
		cells:       cells,
		overcrowded: overcrowded,
		farsighted:  farsighted,
	}
}

func (this *Field) gatherNeighbors(c *Cell) (all []*Cell) {
	if this.farsighted {
		for _, direction := range intgrid.Neighbors8 {
			point := c.Point
			for {
				point = point.Move(direction)
				u := this.cells[point]
				if u == nil {
					break
				}
				if u.State != "." {
					all = append(all, u)
					break
				}
			}
		}
	} else {
		for _, d := range intgrid.Neighbors8 {
			u := c.Point.Move(d)
			n := this.cells[u]
			if n == nil {
				continue
			}
			all = append(all, n)
		}
	}
	return all
}

func (this *Field) Scan() {
	for _, c := range this.cells {
		occupied := 0
		for _, neighbor := range this.gatherNeighbors(c) {
			if neighbor.State == "#" {
				occupied++
			}

			if c.State == "L" && occupied == 0 {
				c.Next = "#"
			} else if c.State == "#" && occupied >= this.overcrowded {
				c.Next = "L"
			} else {
				c.Next = c.State
			}
		}
	}
}

func (this *Field) Update() (changed int) {
	for _, c := range this.cells {
		if c.Next != c.State {
			changed++
		}
		c.State = c.Next
	}
	return changed
}

func (this *Field) CountOccupied() (occupied int) {
	for _, c := range this.cells {
		if c.State == "#" {
			occupied++
		}
	}
	return occupied
}

func (this *Field) String() string {
	builder := new(strings.Builder)
	for y := 0; y < this.height; y++ {
		for x := 0; x < this.width; x++ {
			builder.WriteString(this.cells[intgrid.NewPoint(x, y)].State)
		}
		builder.WriteString("\n")
	}
	return "\n\n" + builder.String() + "\n\n"
}
