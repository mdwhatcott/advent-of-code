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
}

func (this *OpCodeFixture) TestRunProgramA() {
	program := []int{
		1, 9, 10, 3,
		2, 3, 11, 0,
		99,
		30, 40, 50,
	}

	RunProgram(program)

	this.So(program, should.Resemble, []int{
		3500, 9, 10, 70,
		2, 3, 11, 0,
		99,
		30, 40, 50,
	})
}
func (this *OpCodeFixture) TestRunProgramB() {
	program := []int{
		1, 0, 0, 0, 99,
	}

	RunProgram(program)

	this.So(program, should.Resemble, []int{
		2, 0, 0, 0, 99,
	})
}
func (this *OpCodeFixture) TestRunProgramC() {
	program := []int{
		2, 3, 0, 3, 99,
	}

	RunProgram(program)

	this.So(program, should.Resemble, []int{
		2, 3, 0, 6, 99,
	})
}
func (this *OpCodeFixture) TestRunProgramD() {
	program := []int{
		2, 4, 4, 5, 99, 0,
	}

	RunProgram(program)

	this.So(program, should.Resemble, []int{
		2, 4, 4, 5, 99, 9801,
	})
}
func (this *OpCodeFixture) TestRunProgramE() {
	program := []int{
		1, 1, 1, 4, 99, 5, 6, 0, 99,
	}

	RunProgram(program)

	this.So(program, should.Resemble, []int{
		30, 1, 1, 4, 2, 5, 6, 0, 99,
	})
}
