package day05

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

const toy = "dabAcCaCBAcCcaDA"
const toy2 = "yZzY" + "dabAcCaCBAcCcaDA"

func (this *StuffFixture) TestReaction() {
	this.So(react(toy2), should.Equal, "dabCBAcaDA")
}

func (this *StuffFixture) TestAggressiveReaction() {
	this.So(reactAggressive(toy), should.Equal, "daDA")
}
