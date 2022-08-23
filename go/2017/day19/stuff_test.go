package day19

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
	input := []string{
		//123456789012345
		"     |          ", // 0
		"     |  +--+    ", // 1
		"     A  |  C    ", // 2
		" F---|----E|--+ ", // 3
		"     |  |  |  D ", // 4
		"     +B-+  +--+ ", // 5
	}
	turtle := NewTurtle(input)
	for turtle.Orient() {
		turtle.Move()
	}
	this.So(turtle.sequence, should.Equal, "ABCDEF")
}
