package day07

import (
	"testing"

	"advent/lib/util"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
}

const toy = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func (this *StuffFixture) TestToy() {
	sorter := NewTopologicalSort(toy)
	order := sorter.Sort()
	this.So(order, should.Equal, "CABDFE")
}

func (this *StuffFixture) TestPart1() {
	sorter := NewTopologicalSort(util.InputString())
	order := sorter.Sort()
	this.So(order, should.NotEqual, "JADECGHIKFBLMNOPQRSTUVXYZW")
	this.So(order, should.NotEqual, "JDETKPFABUHOQSXVYMLZCNIGRW")
	this.So(order, should.Equal, "JDEKPFABTUHOQSXVYMLZCNIGRW")
}
