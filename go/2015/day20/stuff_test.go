package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuff(t *testing.T) {
	should.Run(&Stuff{T: should.New(t)}, should.Options.UnitTests())
}

type Stuff struct {
	*should.T
}

func (this *Stuff) Test() {
	this.So(presents(1), should.Equal, 10)
	this.So(presents(2), should.Equal, 30)
	this.So(presents(3), should.Equal, 40)
	this.So(presents(4), should.Equal, 70)
	this.So(presents(9), should.Equal, 130)
}
