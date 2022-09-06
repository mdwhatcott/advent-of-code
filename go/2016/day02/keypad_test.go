package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestNineKeyFixture(t *testing.T) {
	should.Run(&NineKeyFixture{T: should.New(t)}, should.Options.UnitTests())
}

type NineKeyFixture struct {
	*should.T

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
	should.Run(&ThirteenKeyFixture{T: should.New(t)}, should.Options.UnitTests())
}

type ThirteenKeyFixture struct {
	*should.T

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
