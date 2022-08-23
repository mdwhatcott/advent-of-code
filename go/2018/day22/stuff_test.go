package day22

import (
	"testing"

	"github.com/mdwhatcott/testing/suite"
)

func TestStuffFixture(t *testing.T) {
	suite.Run(&StuffFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type StuffFixture struct {
	*suite.T
}

func (this *StuffFixture) Setup() {
}

func (this *StuffFixture) Test() {
}
