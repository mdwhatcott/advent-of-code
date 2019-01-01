package day09

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

func (this *StuffFixture) TestExamples() {
	this.So(MarbleHighScore(Parse("9 players; last marble is worth 25 points")), should.Equal, 32)
	this.So(MarbleHighScore(Parse("10 players; last marble is worth 1618 points")), should.Equal, 8317)
	this.So(MarbleHighScore(Parse("13 players; last marble is worth 7999 points")), should.Equal, 146373)
	this.So(MarbleHighScore(Parse("17 players; last marble is worth 1104 points")), should.Equal, 2764)
	this.So(MarbleHighScore(Parse("21 players; last marble is worth 6111 points")), should.Equal, 54718)
	this.So(MarbleHighScore(Parse("30 players; last marble is worth 5807 points")), should.Equal, 37305)
}

