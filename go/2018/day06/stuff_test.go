package day06

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuffFixture(t *testing.T) {
	should.Run(&StuffFixture{T: should.New(t)}, should.Options.UnitTests())
}

type StuffFixture struct {
	*should.T
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
