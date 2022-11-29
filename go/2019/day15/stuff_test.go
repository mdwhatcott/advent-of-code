package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
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
