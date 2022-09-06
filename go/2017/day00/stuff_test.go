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

func (this *Stuff) Setup() {
}

func (this *Stuff) Test() {
	this.So(true, should.BeTrue)
}
