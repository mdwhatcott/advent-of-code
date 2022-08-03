package day12

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"

	"advent/lib/util"
)

func TestDay12Suite(t *testing.T) {
	suite.Run(&Day12Suite{T: suite.New(t)}, suite.Options.UnitTests())
}

type Day12Suite struct {
	*suite.T
	lines util.Strings
}

func (this *Day12Suite) Setup() {
	this.lines = util.InputLines()
}
func (this *Day12Suite) Test1() {
	this.So(Part1(this.lines), should.Equal, 4720)
}
func (this *Day12Suite) Test2() {
	this.So(Part2(this.lines), should.Equal, 147848)
}
