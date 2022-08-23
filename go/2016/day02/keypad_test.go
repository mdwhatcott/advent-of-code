package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestNineKeyFixture(t *testing.T) {
	suite.Run(&NineKeyFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type NineKeyFixture struct {
	*suite.T

	keypad *Keypad
}

func (this *NineKeyFixture) Setup() {
	this.keypad = NewKeypad(nineKeyDirections)
}

func (this *NineKeyFixture) TestFindDigit() {
	this.So(this.keypad.FindDigit("ULL"), should.Equal, "1")
	this.So(this.keypad.FindDigit("RRDDD"), should.Equal, "9")
	this.So(this.keypad.FindDigit("LURDL"), should.Equal, "8")
	this.So(this.keypad.FindDigit("UUUUD"), should.Equal, "5")
}

func (this *NineKeyFixture) TestFullCodeDerivation() {
	this.So(this.keypad.DeriveCode(`ULL
RRDDD
LURDL
UUUUD`), should.Equal, "1985")
}

func TestThirteenKeyFixture(t *testing.T) {
	suite.Run(&ThirteenKeyFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type ThirteenKeyFixture struct {
	*suite.T

	keypad *Keypad
}

func (this *ThirteenKeyFixture) Setup() {
	this.keypad = NewKeypad(thirteenKeyDirections)
}

func (this *ThirteenKeyFixture) TestFindDigit() {
	this.So(this.keypad.FindDigit("ULL"), should.Equal, "5")
	this.So(this.keypad.FindDigit("RRDDD"), should.Equal, "D")
	this.So(this.keypad.FindDigit("LURDL"), should.Equal, "B")
	this.So(this.keypad.FindDigit("UUUUD"), should.Equal, "3")
}

func (this *ThirteenKeyFixture) TestFullCodeDerivation() {
	this.So(this.keypad.DeriveCode(`ULL
RRDDD
LURDL
UUUUD`), should.Equal, "5DB3")
}
