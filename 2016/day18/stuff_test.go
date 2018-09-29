package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestTilesFixture(t *testing.T) {
	gunit.Run(new(TilesFixture), t)
}

type TilesFixture struct {
	*gunit.Fixture

	row *Row
}

func (this *TilesFixture) Setup() {
	this.row = ParseRow("..^^.")
}

func (this *TilesFixture) TestRowChecks() {
	this.So(this.row.NextIsTrap(0), should.BeFalse)
	this.So(this.row.NextIsTrap(1), should.BeTrue)
	this.So(this.row.NextIsTrap(2), should.BeTrue)
	this.So(this.row.NextIsTrap(3), should.BeTrue)
	this.So(this.row.NextIsTrap(4), should.BeTrue)
}

func (this *TilesFixture) TestString() {
	this.So(this.row.String(), should.Equal, "..^^.")
}

func (this *TilesFixture) TestNextRow() {
	this.So(this.row.Next().String(), should.Equal, ".^^^^")
	this.So(this.row.Next().Next().String(), should.Equal, "^^..^")
}
