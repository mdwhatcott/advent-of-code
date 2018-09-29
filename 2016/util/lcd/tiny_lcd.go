package lcd

import "strings"

type LCD struct {
	rows    int
	columns int
	grid    []bool
	changes map[int]bool
}

func NewLCD(columns, rows int) *LCD {
	this := &LCD{
		columns: columns,
		rows:    rows,
		grid:    []bool{},
		changes: make(map[int]bool),
	}
	for cell := 0; cell < columns*rows; cell++ {
		this.grid = append(this.grid, false)
	}
	return this
}

func (this *LCD) RectangleOn(columns, rows int) {
	defer this.apply()

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			this.changes[r*this.columns+c] = true
		}
	}
}

func (this *LCD) RotateColumn(column, rotation int) {
	defer this.apply()

	for r := 0; r < this.rows; r++ {
		currentIndex := r*this.columns + column
		nextIndex := (currentIndex + this.columns*rotation) % len(this.grid)
		this.changes[nextIndex] = this.grid[currentIndex]
	}
}

func (this *LCD) RotateRow(row, rotation int) {
	defer this.apply()

	start := row * this.columns
	end := start + this.columns

	for c := 0; c < this.columns; c++ {
		currentIndex := row*this.columns + c
		nextIndex := currentIndex + rotation
		if nextIndex >= end {
			nextIndex -= this.columns
		}
		this.changes[nextIndex] = this.grid[currentIndex]
	}
}

func (this *LCD) apply() {
	for i, value := range this.changes {
		this.grid[i] = value
	}
	this.changes = make(map[int]bool)
}

func (this *LCD) String() (result string) {
	for i := 0; i < len(this.grid); i++ {
		if this.grid[i] {
			result += "#"
		} else {
			result += " "
		}
		if (i+1)%this.columns == 0 {
			result += "\n"
		}
	}
	return strings.TrimRight(result, "\n")
}
