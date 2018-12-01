package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) Test() {
	this.So(presents(1), should.Equal, 10)
	this.So(presents(2), should.Equal, 30)
	this.So(presents(3), should.Equal, 40)
	this.So(presents(4), should.Equal, 70)
	this.So(presents(9), should.Equal, 130)
}

