package main

import "bytes"

type Grid struct {
	relations map[*cell][]*cell
	cells     [][]*cell
}

func NewGrid(grid string) *Grid {
	self := new(Grid)
	self.cells = initialize(grid)
	self.relations = formRelationships(self.cells)
	return self
}
func initialize(grid string) [][]*cell {
	var rows [][]*cell
	var row []*cell

	for _, c := range grid {
		if c == '\n' && len(row) > 0 {
			rows = append(rows, row)
			row = []*cell{}
		} else if c == '-' {
			row = append(row, newDeadCell())
		} else if c == 'x' {
			row = append(row, newLiveCell())
		}
	}
	if len(row) > 0 {
		rows = append(rows, row)
	}
	return rows
}
func formRelationships(grid [][]*cell) map[*cell][]*cell {
	relations := map[*cell][]*cell{}

	for y, row := range grid {
		for x, cell := range row {
			relations[cell] = neighbors(grid, x, y)
		}
	}
	return relations
}
func neighbors(grid [][]*cell, x, y int) []*cell {
	var yes []*cell

	for _, candidate := range adjoining(x, y) {
		if candidate.isOnGrid(grid) {
			yes = append(yes, grid[candidate.y][candidate.x])
		}
	}
	return yes
}
func adjoining(x, y int) []point {
	return []point{
		{x - 1, y - 1}, // upper left
		{x, y - 1},     // upper
		{x + 1, y - 1}, // upper right
		{x - 1, y},     // left
		{x + 1, y},     // right
		{x - 1, y + 1}, // lower left
		{x, y + 1},     // lower
		{x + 1, y + 1}, // lower right
	}
}

func (self *Grid) LockCornerLightsOn() {
	self.cells[0][0].Lock(true)
	self.cells[len(self.cells)-1][0].Lock(true)
	self.cells[0][len(self.cells[0])-1].Lock(true)
	self.cells[len(self.cells)-1][len(self.cells[0])-1].Lock(true)
}

func (self *Grid) Scan() {
	for cell, neighbors := range self.relations {
		cell.scan(neighbors)
	}

	for cell := range self.relations {
		cell.update()
	}
}

func (self *Grid) String() string {
	builder := bytes.NewBufferString("\n")
	for _, row := range self.cells {
		for _, cell := range row {
			if cell.isAlive() {
				builder.WriteString("x")
			} else {
				builder.WriteString("-")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (self *Grid) CountAlive() (count int) {
	for _, row := range self.cells {
		for _, cell := range row {
			if cell.alive {
				count++
			}
		}
	}
	return count
}

type point struct {
	x int
	y int
}

func (self point) isOnGrid(grid [][]*cell) bool {
	return self.x >= 0 &&
		self.y >= 0 &&
		self.x < len(grid[0]) &&
		self.y < len(grid)
}

///////////////////////////////////////////////////////////////

type cell struct {
	updater func()
	alive   bool
	locked  bool
}

func newDeadCell() *cell {
	return new(cell)
}

func newLiveCell() *cell {
	cell := newDeadCell()
	cell.revive()
	return cell
}

func (self *cell) Lock(alive bool) {
	self.locked = true
	self.alive = alive
}

func (self *cell) isAlive() bool {
	return self.alive
}

func (self *cell) revive() {
	if !self.locked {
		self.alive = true
	}
}

func (self *cell) kill() {
	if !self.locked {
		self.alive = false
	}
}

func (self *cell) scan(neighbors []*cell) {
	alive := self.scanForLifeSigns(neighbors)
	self.decideFate(alive)
}
func (self *cell) scanForLifeSigns(neighbors []*cell) int {
	alive := 0
	for _, neighbor := range neighbors {
		if neighbor.isAlive() {
			alive++
		}
	}
	return alive
}
func (self *cell) decideFate(alive int) {
	if self.isAlive() {
		self.handleLiving(alive)
	} else {
		self.handleDead(alive)
	}
}
func (self *cell) handleLiving(alive int) {
	if alive < 2 || alive > 3 {
		self.updater = self.kill
	} else {
		self.updater = self.revive
	}
}
func (self *cell) handleDead(alive int) {
	if alive == 3 {
		self.updater = self.revive
	} else {
		self.updater = self.kill
	}
}

func (self *cell) update() {
	self.updater()
}
