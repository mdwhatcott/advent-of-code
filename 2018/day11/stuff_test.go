package day11

import (
	"testing"

	"advent/lib/util"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
}

func (this *StuffFixture) TestHundredsDigit() {
	this.So(hundreds(0), should.Equal, 0)
	this.So(hundreds(10), should.Equal, 0)
	this.So(hundreds(100), should.Equal, 1)
	this.So(hundreds(200), should.Equal, 2)
	this.So(hundreds(1000), should.Equal, 0)
	this.So(hundreds(1200), should.Equal, 2)
}

func (this *StuffFixture) TestPowerAt() {
	this.So(NewGrid(8).PowerAt(3, 5), should.Equal, 4)
	this.So(NewGrid(57).PowerAt(122, 79), should.Equal, -5)
	this.So(NewGrid(39).PowerAt(217,196), should.Equal, 0)
	this.So(NewGrid(71).PowerAt(101,153), should.Equal, 4)
}

func (this *StuffFixture) TestPowerSquare() {
	this.So(NewGrid(18).PowerAtSquare(33, 45, 3), should.Equal, 29)
	this.So(NewGrid(42).PowerAtSquare(21, 61, 3), should.Equal, 30)
}

func (this *StuffFixture) TestMaxPowerXY() {
	this.So(NewGrid(18).MaxPowerXY(3), should.Equal, "33,45")
	this.So(NewGrid(42).MaxPowerXY(3), should.Equal, "21,61")
}

func (this *StuffFixture) TestMaxPowerXYZ() {
	this.Print(NewGrid(util.InputInt()))
	//this.So(NewGrid(18).MaxPowerXYSize(), should.Equal, "90,269,16")
	//this.So(NewGrid(42).MaxPowerXYSize(), should.Equal, "232,251,12")
}
