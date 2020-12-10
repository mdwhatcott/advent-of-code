package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestFixture(t *testing.T) {
	gunit.Run(new(Fixture), t)
}

type Fixture struct {
	*gunit.Fixture
}

func (this *Fixture) Setup() {
}

func (this *Fixture) Test() {
	numbers := xRange(1, 26)

	this.So(canSum(append(numbers, 26)...), should.BeTrue)
	this.So(canSum(append(numbers, 49)...), should.BeTrue)
	this.So(canSum(append(numbers, 100)...), should.BeFalse)
	this.So(canSum(append(numbers, 50)...), should.BeFalse)
}

func (this *Fixture) TestLargerExample() {
	this.So(canSum(95, 102, 117, 150, 182, 127), should.BeFalse)
}

func xRange(min, max int) (result []int) {
	for x := min; x < max; x++ {
		result = append(result, x)
	}
	return result
}

var largerExample = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}
