package day11

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/grid"
)

func TestSummedAreaTableFixture(t *testing.T) {
	should.Run(&SummedAreaTableFixture{T: should.New(t)}, should.Options.UnitTests())
}

type SummedAreaTableFixture struct {
	*should.T
	source map[grid.Point]int
	table  SummedAreaTable
}

func (this *SummedAreaTableFixture) Setup() {
	this.source = map[grid.Point]int{
		grid.NewPoint(0, 0): 1, grid.NewPoint(1, 0): 1, grid.NewPoint(2, 0): 1,
		grid.NewPoint(0, 1): 1, grid.NewPoint(1, 1): 1, grid.NewPoint(2, 1): 1,
		grid.NewPoint(0, 2): 1, grid.NewPoint(1, 2): 1, grid.NewPoint(2, 2): 1,
	}
	this.table = NewSummedAreaTable(this.source)
}

func (this *SummedAreaTableFixture) TestCalculateAllSums() {
	this.So(this.table, should.Equal, SummedAreaTable{
		grid.NewPoint(0, 0): 1, grid.NewPoint(1, 0): 2, grid.NewPoint(2, 0): 3,
		grid.NewPoint(0, 1): 2, grid.NewPoint(1, 1): 4, grid.NewPoint(2, 1): 6,
		grid.NewPoint(0, 2): 3, grid.NewPoint(1, 2): 6, grid.NewPoint(2, 2): 9,
	})
}

func (this *SummedAreaTableFixture) TestCalculateSummedArea() {
	this.So(this.table.SummedArea(grid.NewPoint(0, 0), 2), should.Equal, 4)
	this.So(this.table.SummedArea(grid.NewPoint(1, 1), 2), should.Equal, 4)
}
