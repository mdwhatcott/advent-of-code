package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestScramblerFixture(t *testing.T) {
	should.Run(&ScramblerFixture{T: should.New(t)}, should.Options.UnitTests())
}

type ScramblerFixture struct {
	*should.T
	scrambler *Scrambler
}

func (this *ScramblerFixture) Setup() {

}

func (this *ScramblerFixture) TestPart1Example() {
	this.scrambler = NewScrambler([]string{
		"swap position 4 with position 0",
		"swap letter d with letter b",
		"reverse positions 0 through 4",
		"rotate left 1 step",
		"move position 1 to position 4",
		"move position 3 to position 0",
		"rotate based on position of letter b",
		"rotate based on position of letter d",
	})

	this.So(this.scrambler.Process("abcde"), should.Equal, "decab")
}

func (this *ScramblerFixture) TestSwapPositions() {
	result := SwapPositions("abcde", 4, 0)
	this.So(result, should.Equal, "ebcda")
}

func (this *ScramblerFixture) TestSwapLetters() {
	result := SwapLetters("ebcda", "d", "b")
	this.So(result, should.Equal, "edcba")
}

func (this *ScramblerFixture) TestReverse() {
	this.So(ReverseRange("edcba", 0, 4), should.Equal, "abcde")
	this.So(ReverseRange("edcba", 1, 3), should.Equal, "ebcda")
}

func (this *ScramblerFixture) TestRotate() {
	this.So(Rotate("abcde", 0), should.Equal, "abcde")
	this.So(Rotate("abcde", 5), should.Equal, "abcde")
	this.So(Rotate("abcde", -5), should.Equal, "abcde")
	this.So(Rotate("abcde", 10), should.Equal, "abcde")
	this.So(Rotate("abcde", -10), should.Equal, "abcde")

	this.So(Rotate("abcde", 1), should.Equal, "eabcd")
	this.So(Rotate("abcde", -1), should.Equal, "bcdea")
}

func (this *ScramblerFixture) TestMove() {
	this.So(Move("bcdea", 1, 4), should.Equal, "bdeac")
	this.So(Move("bdeac", 3, 0), should.Equal, "abdec")
}

func (this *ScramblerFixture) TestRotateByChar() {
	this.So(RotateByChar("abdec", "b"), should.Equal, "ecabd")
	this.So(RotateByChar("ecabd", "d"), should.Equal, "decab")

	//                    original                        rotated
	this.So(RotateByChar("llllllla", "a"), should.Equal, "alllllll")
	this.So(RotateByChar("alllllll", "a"), should.Equal, "lallllll")
	this.So(RotateByChar("llllalll", "a"), should.Equal, "llalllll")
	this.So(RotateByChar("lallllll", "a"), should.Equal, "lllallll")
	this.So(RotateByChar("lllllall", "a"), should.Equal, "llllalll")
	this.So(RotateByChar("llalllll", "a"), should.Equal, "lllllall")
	this.So(RotateByChar("llllllal", "a"), should.Equal, "llllllal")
	this.So(RotateByChar("lllallll", "a"), should.Equal, "llllllla")
}
