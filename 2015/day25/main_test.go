package main

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

func (this *Stuff) TestTotalCalculations() {
	this.So(totalCalculations(1, 1), should.Equal, 0)
	this.So(totalCalculations(1, 2), should.Equal, 1)
	this.So(totalCalculations(2, 1), should.Equal, 2)
	this.So(totalCalculations(1, 3), should.Equal, 3)
	this.So(totalCalculations(2, 2), should.Equal, 4)
	this.So(totalCalculations(3, 1), should.Equal, 5)
	this.So(totalCalculations(4, 3), should.Equal, 18)
	this.So(totalCalculations(6, 1), should.Equal, 20)
}
func (this *Stuff) TestCalculateLocation() {
	this.So(calculateLocation(1, 1), should.Equal, 20151125)
	this.So(calculateLocation(1, 2), should.Equal, 31916031)
	this.So(calculateLocation(2, 1), should.Equal, 18749137)
	this.So(calculateLocation(1, 3), should.Equal, 16080970)
	this.So(calculateLocation(2, 2), should.Equal, 21629792)
	this.So(calculateLocation(3, 1), should.Equal, 17289845)
}
