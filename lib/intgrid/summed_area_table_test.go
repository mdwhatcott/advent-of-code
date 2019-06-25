package intgrid

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSummedAreaTableFixture(t *testing.T) {
	gunit.Run(new(SummedAreaTableFixture), t)
}

type SummedAreaTableFixture struct {
	*gunit.Fixture

	table  [][]int
	summed SummedAreaTable
}

func (this *SummedAreaTableFixture) Setup() {
	this.table = [][]int{
		{31, 2, 4, 33, 5, 36},
		{12, 26, 9, 10, 29, 25},
		{13, 17, 21, 22, 20, 18},
		{24, 23, 15, 16, 14, 19},
		{30, 8, 28, 27, 11, 7},
		{1, 35, 34, 3, 32, 6},
	}
	this.summed = NewSummedAreaTable(this.table)
}

func (this *SummedAreaTableFixture) TestInitialization() {
	expected := SummedAreaTable{
		{31, 33, 37, 70, 75, 111},
		{43, 71, 84, 127, 161, 222},
		{56, 101, 135, 200, 254, 333},
		{80, 148, 197, 278, 346, 444},
		{110, 186, 263, 371, 450, 555},
		{111, 222, 333, 444, 555, 666},
	}
	this.So(this.summed, should.Resemble, expected)
}

func (this *SummedAreaTableFixture) TestLookupQuadrant() {
	quadrantSum := this.summed.SumQuadrant(NewPoint(2, 3), NewPoint(4, 4))
	this.So(quadrantSum, should.Resemble, 111)
}
