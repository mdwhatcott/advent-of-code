package day06

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
