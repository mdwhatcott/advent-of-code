package day05

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuff(t *testing.T) {
	should.Run(&Stuff{T: should.New(t)}, should.Options.UnitTests())
}

type Stuff struct {
	*should.T

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
