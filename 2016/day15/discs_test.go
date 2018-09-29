package main

import (
	"testing"
	"github.com/smartystreets/gunit"
	"github.com/smartystreets/assertions/should"
)

func TestDiscs(t *testing.T) {
    gunit.Run(new(Discs), t)
}

type Discs struct {
    *gunit.Fixture
}

func (this *Discs) Setup() {
}

func (this *Discs) TestParse() {
	this.So(ParseDisc("Disc #1 has 13 positions; at time=0, it is at position 1."), should.Resemble, Disc{
		Delay: 1,
		Positions: 13,
		Start: 1,
	})

	this.So(ParseDisc("Disc #2 has 19 positions; at time=0, it is at position 10."), should.Resemble, Disc{
		Delay: 2,
		Positions: 19,
		Start: 10,
	})
}
