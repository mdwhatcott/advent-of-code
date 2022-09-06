package day15

import (
	"testing"
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
}
