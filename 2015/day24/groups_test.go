package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSleighFixture(t *testing.T) {
	gunit.Run(new(SleighFixture), t)
}

type SleighFixture struct {
	*gunit.Fixture
}

func (this *SleighFixture) TestEqualityOfIntSlices() {
	this.So(areEqual([]int{1}, []int{}), should.BeFalse)
	this.So(areEqual([]int{1}, []int{2}), should.BeFalse)
	this.So(areEqual([]int{1, 2, 3}, []int{1, 2, 3}), should.BeTrue)
	this.So(areEqual([]int{1, 2, 3}, []int{3, 2, 1}), should.BeTrue)
}

func (this *SleighFixture) TestGroupingUnbalanced() {
	this.So(new(Sleigh).IsBalanced(), should.BeTrue)
	this.So((&Sleigh{
		A: []int{1},
		B: []int{1},
		C: []int{2},
	}).IsBalanced(), should.BeFalse)
	this.So((&Sleigh{
		A: []int{1},
		B: []int{1},
		C: []int{1},
	}).IsBalanced(), should.BeTrue)
}

func (this *SleighFixture) TestQuantumEntanglement() {
	sleigh := new(Sleigh)
	sleigh.A = []int{3, 4, 5}
	expected := QuantumEntanglement(sleigh.A...)
	this.So(sleigh.QuantumEntanglement(), should.Equal, expected)
}

func (this *SleighFixture) TestComfort() {
	this.So((&Sleigh{A: []int{1}, B: []int{1, 1}, C: []int{1, 1}}).IsComfortable(), should.BeTrue)
	this.So((&Sleigh{A: []int{1}, B: []int{1}, C: []int{1}}).IsComfortable(), should.BeTrue)
	this.So((&Sleigh{A: []int{1, 1}, B: []int{1}, C: []int{1}}).IsComfortable(), should.BeFalse)
}

func (this *SleighFixture) TestEnumerateSleighConfigurations() {
}

func (this *SleighFixture) TestSumsOfTwenty() {
	sums := SumsOfTheThirdPart(1, 2, 3, 4, 5, 7, 8, 9, 10, 11)
	this.So(sums, should.HaveLength, 1)
}
