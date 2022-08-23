package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestDiscs(t *testing.T) {
	suite.Run(&Discs{T: suite.New(t)}, suite.Options.UnitTests())
}

type Discs struct {
	*suite.T
}

func (this *Discs) Setup() {
}

func (this *Discs) TestParse() {
	this.So(ParseDisc("Disc #1 has 13 positions; at time=0, it is at position 1."), should.Equal, Disc{
		Delay:     1,
		Positions: 13,
		Start:     1,
	})

	this.So(ParseDisc("Disc #2 has 19 positions; at time=0, it is at position 10."), should.Equal, Disc{
		Delay:     2,
		Positions: 19,
		Start:     10,
	})
}
