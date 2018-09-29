package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestDragonChecksumFixture(t *testing.T) {
	gunit.Run(new(DragonChecksumFixture), t)
}

type DragonChecksumFixture struct {
	*gunit.Fixture
}

func (this *DragonChecksumFixture) TestReverse() {
	this.So(Reverse(""), should.Equal, "")
	this.So(Reverse("1"), should.Equal, "1")
	this.So(Reverse("11"), should.Equal, "11")
	this.So(Reverse("10"), should.Equal, "01")
	this.So(Reverse("101"), should.Equal, "101")
	this.So(Reverse("1010"), should.Equal, "0101")
}

func (this *DragonChecksumFixture) TestInvert() {
	this.So(Invert("1"), should.Equal, "0")
	this.So(Invert("0"), should.Equal, "1")
	this.So(Invert("1011"), should.Equal, "0100")
	this.So(Invert("1101"), should.Equal, "0010")
}

func (this *DragonChecksumFixture) TestDragon() {
	this.So(Dragon("1"), should.Equal, "100")
	this.So(Dragon("0"), should.Equal, "001")
	this.So(Dragon("11111"), should.Equal, "11111000000")
	this.So(Dragon("111100001010"), should.Equal, "1111000010100101011110000")
}

func (this *DragonChecksumFixture) TestExpand() {
	this.So(ExpandDragon("111100001010", 15), should.Equal, "1111000010100101011110000"[:15])
}

func (this *DragonChecksumFixture) TestChecksum() {
	this.So(Checksum("110010110100"), should.Equal, "100")
	this.So(DragonChecksum("10000", 20), should.Equal, "01100")
}