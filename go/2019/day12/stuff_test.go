package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestStuffFixture(t *testing.T) {
	suite.Run(&StuffFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type StuffFixture struct {
	*suite.T
}

func (this *StuffFixture) TestExample1() {
	this.So(CalculateCombinedEnergy(10, e1a, e1b, e1c, e1d), should.Equal, 179)
}
func (this *StuffFixture) TestExample2() {
	this.So(CalculateCombinedEnergy(100, e2a, e2b, e2c, e2d), should.Equal, 1940)
}

func (this *StuffFixture) TestCalculatePeriods_Example1() {
	x, y, z := CalculatePeriods(e1a, e1b, e1c, e1d)
	this.So(x, should.Equal, 18)
	this.So(y, should.Equal, 28)
	this.So(z, should.Equal, 44)
	this.So(x*y*z, should.Equal, 22176)

	this.So(CalculatePeriodIntersection(x, y, z), should.Equal, 2772)
}

func (this *StuffFixture) TestCalculatePeriods_Example2() {
	x, y, z := CalculatePeriods(e2a, e2b, e2c, e2d)
	this.So(x, should.Equal, 2028)
	this.So(y, should.Equal, 5898)
	this.So(z, should.Equal, 4702)
	this.So(x*y*z, should.Equal, 56241299088)

	this.So(CalculatePeriodIntersection(x, y, z), should.Equal, 4686774924)
}
