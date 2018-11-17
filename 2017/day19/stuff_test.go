package day19

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
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
