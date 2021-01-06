package day11

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"

	"advent/lib/grid"
)

func TestGrid2Fixture(t *testing.T) {
	gunit.Run(new(Grid2Fixture), t)
}

type Grid2Fixture struct {
	*gunit.Fixture
}

func (this *Grid2Fixture) TestExampleSerialNumber_18() {
	field := InitializePowerGrid(300, 18)
	table := NewSummedAreaTable(field)
	area := table.SummedArea(grid.NewPoint(33, 45), 3)
	this.So(area, should.Equal, 29)
}

func (this *Grid2Fixture) TestExampleSerialNumber_42() {
	field := InitializePowerGrid(300, 42)
	table := NewSummedAreaTable(field)
	area := table.SummedArea(grid.NewPoint(21, 61), 3)
	this.So(area, should.Equal, 30)
}
