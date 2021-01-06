package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestAuntSueFixture(t *testing.T) {
	gunit.Run(new(AuntSueFixture), t)
}

type AuntSueFixture struct {
	*gunit.Fixture
	search AuntSue
}

func (this *AuntSueFixture) Setup() {
	this.search = AuntSue{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

func (this *AuntSueFixture) TestAunt() {
	allConflict := AuntSue{
		"children": 1,
		"cars":     8,
		"vizslas":  7,
	}
	oneConflict := AuntSue{
		"children": 3,
		"cats":     8,
	}
	zeroConflict := AuntSue{
		"permufes": 1,
		"cars":     2,
	}

	this.So(allConflict.Matches(this.search), should.BeFalse)
	this.So(oneConflict.Matches(this.search), should.BeFalse)
	this.So(zeroConflict.Matches(this.search), should.BeTrue)
}
