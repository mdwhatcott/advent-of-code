package day12

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

func (this *StuffFixture) TestSingleRule() {
	rule := ParseRule("..... => #")
	this.So(rule.Transform("....."), should.Equal, "#")
	this.So(rule.Transform("#...."), should.Equal, ".")
}
