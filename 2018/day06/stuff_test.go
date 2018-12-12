package day06

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

const toy = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

func (this *StuffFixture) TestToy() {
	largest := calculateLargestFiniteArea(parsePoints(toy))
	this.So(largest, should.Equal, 17)
}
