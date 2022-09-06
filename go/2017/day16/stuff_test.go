package day16

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuff(t *testing.T) {
	should.Run(&Stuff{T: should.New(t)}, should.Options.UnitTests())
}

type Stuff struct {
	*should.T
}

func (this *Stuff) TestSpin() {
	this.So(spin("abcde", 1), should.Equal, "eabcd")
	this.So(spin("abcde", 2), should.Equal, "deabc")
	this.So(spin("abcde", 3), should.Equal, "cdeab")
	this.So(spin("abcde", 4), should.Equal, "bcdea")
	this.So(spin("abcde", 5), should.Equal, "abcde")
}

func (this *Stuff) TestExchange() {
	this.So(exchange("abcde", 3, 4), should.Equal, "abced")
	this.So(exchange("abcde", 0, 4), should.Equal, "ebcda")
}

func (this *Stuff) TestPartner() {
	this.So(partner("abcde", "a", "c"), should.Equal, "cbade")
}

func (this *Stuff) TestInterpreter() {
	interpreter := NewInterpreter("abcde")
	this.So(interpreter.Dance("s1"), should.Equal, "eabcd")
	this.So(interpreter.Dance("x3/4"), should.Equal, "eabdc")
	this.So(interpreter.Dance("pe/b"), should.Equal, "baedc")
	this.So(interpreter.state, should.Equal, "baedc")
}
