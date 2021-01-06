package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestJSONAbacusFixture(t *testing.T) {
	gunit.Run(new(JSONAbacusFixture), t)
}

type JSONAbacusFixture struct {
	*gunit.Fixture
}

func (this *JSONAbacusFixture) TestPart1() {
	this.So(part1(`[]`), should.Equal, 0)
	this.So(part1(`{}`), should.Equal, 0)

	this.So(part1(`[1,2,3]`), should.Equal, 6)
	this.So(part1(`{"a":2,"b":4}`), should.Equal, 6)

	this.So(part1(`[[[3]]]`), should.Equal, 3)
	this.So(part1(`{"a":{"b":4},"c":-1}`), should.Equal, 3)

	this.So(part1(`{"a":[-1,1]}`), should.Equal, 0)
	this.So(part1(`[-1,{"a":1}]`), should.Equal, 0)
}

func (this *JSONAbacusFixture) TestPart2() {
	this.So(part2(`[1,2,3]`), should.Equal, 6)
	this.So(part2(`[1,{"c":"red","b":2},3]`), should.Equal, 4)
	this.So(part2(`{"d":"red","e":[1,2,3,4],"f":5}`), should.Equal, 0)
	this.So(part2(`[1,"red",5]`), should.Equal, 6)
}
