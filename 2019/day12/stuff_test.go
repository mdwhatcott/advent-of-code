package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
}

func (this *StuffFixture) TestExample1() {
	this.So(CalculateCombinedEnergy(10, e1a, e1b, e1c, e1d), should.Equal, 179)
}
func (this *StuffFixture) TestExample2() {
	this.So(CalculateCombinedEnergy(100, e2a, e2b, e2c, e2d), should.Equal, 1940)
}
