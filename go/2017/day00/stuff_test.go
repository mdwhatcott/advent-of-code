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

func (this *Stuff) Setup() {
}

func (this *Stuff) Test() {
	this.So(true, should.BeTrue)
}
