package day05

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

const toy = "dabAcCaCBAcCcaDA"
const toy2 = "yZzY" + "dabAcCaCBAcCcaDA"

func (this *StuffFixture) TestReaction() {
	this.So(react(toy2), should.Equal, "dabCBAcaDA")
}

func (this *StuffFixture) TestAggressiveReaction() {
	this.So(reactAggressive(toy), should.Equal, "daDA")
}
