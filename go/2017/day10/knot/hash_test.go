package knot

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) TestKnotHash() {
	this.So(HashString(""), should.Equal, "a2582a3a0e66e6e86e3812dcb672a272")
	this.So(HashString("AoC 2017"), should.Equal, "33efeb34ea91902bb2f59c9920caa6cd")
	this.So(HashString("1,2,3"), should.Equal, "3efbe78a8d82f29979031a4aa0b16a9d")
	this.So(HashString("1,2,4"), should.Equal, "63960835bcdc130f0b66d7ff4f6a5a8e")
}
