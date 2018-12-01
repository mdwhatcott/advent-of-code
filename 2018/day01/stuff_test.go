package day01

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
    gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
    *gunit.Fixture

    stuff *Stuff
}

func (this *StuffFixture) Setup() {
	this.stuff = NewStuff("")
}

func (this *StuffFixture) Test() {
	this.stuff.Process()
	this.So(this.stuff.Answer(), should.NotBeNil)
}
