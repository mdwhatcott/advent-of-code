package day15

import (
	"testing"
)

func TestStuffFixture(t *testing.T) {
	should.Run(&StuffFixture{T: should.New(t)}, should.Options.UnitTests())
}

type StuffFixture struct {
	*should.T
}

func (this *StuffFixture) Setup() {
}

func (this *StuffFixture) Test() {
}
