package day04

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
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
