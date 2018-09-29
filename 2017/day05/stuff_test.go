package day05

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

	program *Program
}

func (this *Stuff) Setup() {
	this.program = NewProgram([]int{0, 3, 0, 1, -3})
}

func (this *Stuff) TestPart1Example() {
	this.So(this.program.Execute(), should.Equal, 5)
}

func (this *Stuff) TestPart2Example() {
	this.program.Part2()
	this.So(this.program.Execute(), should.Equal, 10)
}
