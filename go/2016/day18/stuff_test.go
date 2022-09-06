package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestTilesFixture(t *testing.T) {
	should.Run(&TilesFixture{T: should.New(t)}, should.Options.UnitTests())
}

type TilesFixture struct {
	*should.T

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
