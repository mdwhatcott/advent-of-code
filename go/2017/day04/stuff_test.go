package day04

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestStuff(t *testing.T) {
	suite.Run(&Stuff{T: suite.New(t)}, suite.Options.UnitTests())
}

type Stuff struct {
	*suite.T
}

func (this *Stuff) TestValidPassphrase() {
	this.So(Valid("aa bb cc"), should.BeTrue)
	this.So(Valid("aa bb cc aa"), should.BeFalse)
}

func (this *Stuff) TestValidPassphrase2() {
	this.So(Valid("abcde fghij"), should.BeTrue)
	this.So(Valid("abcde xyz ecdab"), should.BeFalse)
	this.So(Valid("a ab abc abd abf abj"), should.BeTrue)
	this.So(Valid("iiii oiii ooii oooi oooo"), should.BeTrue)
	this.So(Valid("oiii ioii iioi iiio"), should.BeFalse)
}
