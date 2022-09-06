package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPasswordRotationFixture(t *testing.T) {
	should.Run(&PasswordRotationFixture{T: should.New(t)}, should.Options.UnitTests())
}

type PasswordRotationFixture struct {
	*should.T
}

func (this *PasswordRotationFixture) TestIncrementAlpha() {
	this.So(Increment("a"), should.Equal, "b")
	this.So(Increment("aa"), should.Equal, "ab")

	this.So(Increment("h"), should.Equal, "j") // skip 'i'
	this.So(Increment("k"), should.Equal, "m") // skip 'l'
	this.So(Increment("n"), should.Equal, "p") // skip 'o'

	this.So(Increment("az"), should.Equal, "ba")
	this.So(Increment("azz"), should.Equal, "baa")
	this.So(Increment("bcefbzz"), should.Equal, "bcefcaa")

	this.So(Increment("abcdefgh"), should.Equal, "abcdffaa")
	this.So(Increment("ghijklmn"), should.Equal, "ghjaabcc")
}

func (this *PasswordRotationFixture) TestIsValid() {
	this.So(IsValid([]byte("hijklmmn")), should.BeFalse) // contains disallowed 'i' and 'j'
	this.So(IsValid([]byte("abbceffg")), should.BeFalse) // lacks 3-long increasing straight
	this.So(IsValid([]byte("abbcegjk")), should.BeFalse) // lacks two pairs

	this.So(IsValid([]byte("abcdffaa")), should.BeTrue)
	this.So(IsValid([]byte("ghjaabcc")), should.BeTrue)
}
