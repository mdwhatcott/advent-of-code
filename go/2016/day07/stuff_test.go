package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestABBAFixture(t *testing.T) {
	suite.Run(&ABBAFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type ABBAFixture struct {
	*suite.T
}

func (this *ABBAFixture) TestHasABBA() {
	this.So(HasABBA("abba"), should.BeTrue)
	this.So(HasABBA("cabbad"), should.BeTrue)
	this.So(HasABBA("aaaaaaaacddc"), should.BeTrue)
	this.So(HasABBA("cddcaaaaaaaa"), should.BeTrue)
	this.So(HasABBA("aacddcaaaaaa"), should.BeTrue)
	this.So(HasABBA("acddcaaaaaa"), should.BeTrue)
}
func (this *ABBAFixture) TestIsABBA() {
	this.So(IsABBA("abba"), should.BeTrue)
	this.So(IsABBA("aaaa"), should.BeFalse)
	this.So(IsABBA("abcd"), should.BeFalse)
	this.So(IsABBA("abbd"), should.BeFalse)
	this.So(IsABBA("abbb"), should.BeFalse)
	this.So(IsABBA("aaab"), should.BeFalse)
	this.So(IsABBA("abab"), should.BeFalse)
}
func (this *ABBAFixture) TestTLSCompliance() {
	this.So(IsTLSCompliant("abba[mnop]qrst"), should.BeTrue)
	this.So(IsTLSCompliant("qrst[mnop]abba"), should.BeTrue)
	this.So(IsTLSCompliant("abcd[bddb]xyyx"), should.BeFalse)
	this.So(IsTLSCompliant("aaaa[qwer]tyui"), should.BeFalse)
	this.So(IsTLSCompliant("ioxxoj[asdfgh]zxcvbn"), should.BeTrue)
	this.So(IsTLSCompliant("asdf[asdf]asdf[oxxo]asdf"), should.BeFalse)
	this.So(IsTLSCompliant("asdf[asdf]asdf[asdf]asdfoxxoas"), should.BeTrue)
}

func (this *ABBAFixture) TestIsABA() {
	this.So(IsABA("aba"), should.BeTrue)
	this.So(IsABA("bab"), should.BeTrue)
	this.So(IsABA("abb"), should.BeFalse)
	this.So(IsABA("baa"), should.BeFalse)
	this.So(IsABA("bbb"), should.BeFalse)
	this.So(IsABA("aaa"), should.BeFalse)
}

func (this *ABBAFixture) TestHasABA() {
	this.So(ExtractABA("qqqqqq"+"aba"+"zzzzzzz"), should.Equal, []string{"aba"})
	this.So(ExtractABA("qqqqqqzzzzzzz"+"aba"), should.Equal, []string{"aba"})
	this.So(ExtractABA("aba"+"qqqqqqzzzzzzz"), should.Equal, []string{"aba"})
	this.So(ExtractABA("qqqqqqzzzzzzz"), should.BeEmpty)
	this.So(ExtractABA("zazbz"), should.Equal, []string{"zaz", "zbz"})
}

func (this *ABBAFixture) HasCorrespondingBAB() {
	this.So(HasCorrespondingBAB("qqqqqq"+"bab"+"zzzzzz", "aba"), should.BeTrue)
	this.So(HasCorrespondingBAB("qqqqqq"+"aba"+"zzzzzz", "aba"), should.BeFalse)
}

func (this *ABBAFixture) TestIsSSLCompliant() {
	this.So(IsSSLCompliant("aba[bab]xyz"), should.BeTrue)
	this.So(IsSSLCompliant("xyx[xyx]xyx"), should.BeFalse)
	this.So(IsSSLCompliant("aaa[kek]eke"), should.BeTrue)
	this.So(IsSSLCompliant("zazbz[bzb]cdb"), should.BeTrue)
}
