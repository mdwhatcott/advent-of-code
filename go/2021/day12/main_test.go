package day12

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func TestDay12Suite(t *testing.T) {
	should.Run(&Day12Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Day12Suite struct {
	*should.T
	lines util.Slice[string]
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
