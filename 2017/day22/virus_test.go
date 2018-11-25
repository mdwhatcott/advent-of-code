package day22

import (
	"testing"

	"advent/lib/grid"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) Setup() {
}

func (this *Stuff) TestTheWormStartsInTheMiddleAndFacesUp() {
	virus := NewVirus("..#\n#..\n...")
	this.So(virus.Current(), should.Resemble, grid.NewPoint(1, 1))
	this.So(virus.facing, should.Resemble, grid.Up)
}

func (this *Stuff) TestWhenCurrentNodeAlreadyInfected_TurnRightAndLeavesClean() {
	virus := NewVirus("...\n.#.\n...")
	this.So(virus.Current(), should.Resemble, grid.NewPoint(1, 1))
	virus.Move()
	this.So(virus.Current(), should.Resemble, grid.NewPoint(2, 1))
	this.So(virus.state[grid.NewPoint(1, 1)], should.BeFalse)
}

func (this *Stuff) TestWhenCurrentNodeClean_TurnLeftAndLeavesInfected() {
	virus := NewVirus("..#\n...\n...")
	this.So(virus.Current(), should.Resemble, grid.NewPoint(1, 1))
	virus.Move()
	this.So(virus.Current(), should.Resemble, grid.NewPoint(0, 1))
	this.So(virus.state[grid.NewPoint(1, 1)], should.BeTrue)
}

func (this *Stuff) TestExample1() {
	virus := NewVirus("..#\n#..\n...")
	for x := 0; x < 7; x++ {
		virus.Move()
	}
	this.So(virus.Infected(), should.Equal, 5)
}

func (this *Stuff) TestExample2() {
	virus := NewVirus("..#\n#..\n...")
	for x := 0; x < 70; x++ {
		virus.Move()
	}
	this.So(virus.Infected(), should.Equal, 41)
}

func (this *Stuff) TestPart2Example1() {
	virus := NewVirus2("..#\n#..\n...")
	for x := 0; x < 100; x++ {
		virus.Move()
	}
	this.So(virus.Infected(), should.Equal, 26)
}

func (this *Stuff) TestPart2Example2() {
	virus := NewVirus2("..#\n#..\n...")
	for x := 0; x < 10000000; x++ {
		virus.Move()
	}
	this.So(virus.Infected(), should.Equal, 2511944)
}