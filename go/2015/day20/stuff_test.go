package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestStuff(t *testing.T) {
	suite.Run(&Stuff{T: suite.New(t)}, suite.Options.UnitTests())
}

type Stuff struct {
	*suite.T
}

func (this *Stuff) Test() {
	this.So(presents(1), should.Equal, 10)
	this.So(presents(2), should.Equal, 30)
	this.So(presents(3), should.Equal, 40)
	this.So(presents(4), should.Equal, 70)
	this.So(presents(9), should.Equal, 130)
}
