package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestLookSayFixture(t *testing.T) {
	gunit.Run(new(LookSayFixture), t)
}

type LookSayFixture struct {
	*gunit.Fixture
}

func (this *LookSayFixture) Test() {
	this.So(LookSay("1"), should.Equal, "11")
	this.So(LookSay("11"), should.Equal, "21")
	this.So(LookSay("21"), should.Equal, "1211")
	this.So(LookSay("211"), should.Equal, "1221")
	this.So(LookSay("1211"), should.Equal, "111221")
	this.So(LookSay("111221"), should.Equal, "312211")
}
