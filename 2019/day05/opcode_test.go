package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestOpCodeFixture(t *testing.T) {
	gunit.Run(new(OpCodeFixture), t)
}

type OpCodeFixture struct {
	*gunit.Fixture
	i       int
	inputs  []int
	outputs []int
}

func (this *OpCodeFixture) input() int {
	this.i++
	return this.inputs[(this.i-1)%len(this.inputs)]
}
func (this *OpCodeFixture) output(value int) {
	this.outputs = append(this.outputs, value)
}
func (this *OpCodeFixture) run(program []int) {
	RunProgram(program, this.input, this.output)
}

func (this *OpCodeFixture) TestRunProgramA() {
	program := []int{
		1, 9, 10, 3,
		2, 3, 11, 0,
		99,
		30, 40, 50,
	}
	this.run(program)
	this.So(program, should.Resemble, []int{
		3500, 9, 10, 70,
		2, 3, 11, 0,
		99,
		30, 40, 50,
	})
}
func (this *OpCodeFixture) TestRunProgramB() {
	program := []int{1, 0, 0, 0, 99}
	this.run(program)
	this.So(program, should.Resemble, []int{2, 0, 0, 0, 99})
}
func (this *OpCodeFixture) TestRunProgramC() {
	program := []int{2, 3, 0, 3, 99}
	this.run(program)
	this.So(program, should.Resemble, []int{2, 3, 0, 6, 99})
}
func (this *OpCodeFixture) TestRunProgramD() {
	program := []int{2, 4, 4, 5, 99, 0}
	this.run(program)
	this.So(program, should.Resemble, []int{2, 4, 4, 5, 99, 9801})
}
func (this *OpCodeFixture) TestRunProgramE() {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	this.run(program)
	this.So(program, should.Resemble, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}

func (this *OpCodeFixture) SkipTestRunProgramInputOutput() {
	program := []int{3, 0, 4, 0, 99}
	this.run(program)
	this.So(program, should.Resemble, []int{3, 0, 4, 0, 99})
}

func (this *OpCodeFixture) SkipTestRunProgramF() {
	program := []int{1002, 4, 3, 4, 33}
	this.run(program)
	this.So(program, should.Resemble, []int{1002, 4, 3, 4, 99})
}
