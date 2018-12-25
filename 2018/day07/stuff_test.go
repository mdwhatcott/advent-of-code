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

func (this *StuffFixture) TestPart1Toy() {
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

func (this *StuffFixture) TestPart2Toy() {
	sorter := NewConcurrentTopologicalSort(toy, 2, 0)
	order := sorter.Sort()
	duration := sorter.DurationSeconds()
	this.So(order, should.Equal, "CABFDE")
	this.So(duration, should.Equal, 15)
	for _, worker := range sorter.workers {
		this.Println("Worker:", worker.history)
	}
}

func (this *StuffFixture) TestPart2() {
	sorter := NewConcurrentTopologicalSort(util.InputString(), 5, 60)
	order := sorter.Sort()
	this.So(order, should.Equal, "JKXDEPTFABUHOQSVYZMLNCIGRW")
	duration := sorter.DurationSeconds()
	this.So(duration, should.BeGreaterThan, 1044)
	//for _, worker := range sorter.workers {
	//	this.Println("Worker:", worker.history)
	//}
}

func (this *StuffFixture) FocusTestFogleman() {
	order1, seconds1 := Do(util.InputScanner(), 1)
	this.So(order1, should.Equal, "JDEKPFABTUHOQSXVYMLZCNIGRW")
	this.So(seconds1, should.Equal, 1911)

	order5, seconds5 := Do(util.InputScanner(), 5)
	this.So(order5, should.Equal, "JKXDEPTFABUHOQSVYZMLNCIGRW")
	this.So(seconds5, should.Equal, 1048)
}