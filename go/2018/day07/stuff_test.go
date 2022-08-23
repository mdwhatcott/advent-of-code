package day07

import (
	"testing"

	"github.com/mdwhatcott/testing/suite"

	"advent/lib/util"

	"github.com/mdwhatcott/testing/should"
)

func TestStuffFixture(t *testing.T) {
	suite.Run(&StuffFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type StuffFixture struct {
	*suite.T
}

const toy = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func (this *StuffFixture) FocusTestFogleman() {
	orderToy, _ := NewTopologicalSort(toy, 1, 1).Sort()
	this.So(orderToy, should.Equal, "CABDFE")

	order1, _ := NewTopologicalSort(util.InputString(), 1, 1).Sort()
	this.So(order1, should.Equal, "JDEKPFABTUHOQSXVYMLZCNIGRW")

	order5, seconds5 := NewTopologicalSort(util.InputString(), 5, 60).Sort()
	this.So(order5, should.Equal, "JKXDEPTFABUHOQSVYZMLNCIGRW")
	this.So(seconds5, should.Equal, 1048)
}
