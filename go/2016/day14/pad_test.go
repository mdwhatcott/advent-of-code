package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestOneTimePadFixture(t *testing.T) {
	suite.Run(&OneTimePadFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type OneTimePadFixture struct {
	*suite.T
	generator *Generator
}

func (this *OneTimePadFixture) Setup() {
	this.generator = NewGenerator("abc")
}

func (this *OneTimePadFixture) TestRuns() {
	this.So(this.generator.firstRunOfThree("1234555678"), should.Equal, "5")
	this.So(this.generator.firstRunOfThree("1234555"), should.Equal, "5")
}
func (this *OneTimePadFixture) LongTestFind64thKey() {
	this.So(this.generator.IndexOfKey(64), should.Equal, 22728)
}
func (this *OneTimePadFixture) TestIsKey() {
	this.So(this.generator.IsKey(17), should.BeFalse)
	this.So(this.generator.IsKey(18), should.BeFalse)

	this.So(this.generator.IsKey(38), should.BeFalse)
	this.So(this.generator.IsKey(39), should.BeTrue)
	this.So(this.generator.IsKey(40), should.BeFalse)

	this.So(this.generator.IsKey(91), should.BeFalse)
	this.So(this.generator.IsKey(92), should.BeTrue)
	this.So(this.generator.IsKey(93), should.BeFalse)

	this.So(this.generator.IsKey(22728), should.BeTrue)
}

func (this *OneTimePadFixture) TestIsStretchedKey() {
	this.generator.stretch = true
	this.So(this.generator.IsKey(10), should.BeTrue)
}
func (this *OneTimePadFixture) LongTestFind64thStretchedKey() {
	this.generator.stretch = true
	this.So(this.generator.IndexOfKey(64), should.Equal, 22551)
}
