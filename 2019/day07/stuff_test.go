package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t, gunit.Options.AllSequential())
}

type StuffFixture struct {
	*gunit.Fixture
}

func (this *StuffFixture) Test1() {
	runner := NewIO(4, 3, 2, 1, 0)
	answer := runner.Run(3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0)
	this.So(answer, should.Equal, 43210)
}

func (this *StuffFixture) Test2() {
	runner := NewIO(0, 1, 2, 3, 4)
	answer := runner.Run(3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0)
	this.So(answer, should.Equal, 54321)
}

func (this *StuffFixture) Test3() {
	runner := NewIO(1, 0, 4, 3, 2)
	answer := runner.Run(3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0)
	this.So(answer, should.Equal, 65210)
}
