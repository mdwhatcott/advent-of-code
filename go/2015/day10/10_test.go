package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestLookSayFixture(t *testing.T) {
	suite.Run(&LookSayFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type LookSayFixture struct {
	*suite.T
}

func (this *LookSayFixture) Test() {
	this.So(LookSay("1"), should.Equal, "11")
	this.So(LookSay("11"), should.Equal, "21")
	this.So(LookSay("21"), should.Equal, "1211")
	this.So(LookSay("211"), should.Equal, "1221")
	this.So(LookSay("1211"), should.Equal, "111221")
	this.So(LookSay("111221"), should.Equal, "312211")
}
