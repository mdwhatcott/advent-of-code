package day05

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

const toy = "dabAcCaCBAcCcaDA"
const toy2 = "yZzY" + "dabAcCaCBAcCcaDA"

func (this *StuffFixture) TestReaction() {
	this.So(react(toy2), should.Equal, "dabCBAcaDA")
}

func (this *StuffFixture) TestAggressiveReaction() {
	this.So(reactAggressive(toy), should.Equal, "daDA")
}
