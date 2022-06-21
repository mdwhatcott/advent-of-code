package day12

import (
	"testing"

	"github.com/mdwhatcott/testing/suite"
)

func TestDay12Suite(t *testing.T) {
	suite.Run(&Day12Suite{T: suite.New(t)}, suite.Options.UnitTests())
}

type Day12Suite struct {
	*suite.T
}

func (this *Day12Suite) Setup() {
}

func (this *Day12Suite) Test() {
}
