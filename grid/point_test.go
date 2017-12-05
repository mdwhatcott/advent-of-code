package grid

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestPointFixture(t *testing.T) {
	gunit.Run(new(PointFixture), t)
}

type PointFixture struct {
	*gunit.Fixture
}

func (this *PointFixture) TestXY() {
	this.So(NewPoint(1, 2).X(), should.Equal, 1)
	this.So(NewPoint(1, 2).Y(), should.Equal, 2)
}

func (this *PointFixture) TestString() {
	this.So(NewPoint(1, 2).String(), should.Equal, "(1, 2)")
	this.So(NewPoint(-1, -2).String(), should.Equal, "(-1, -2)")
}

func (this *PointFixture) TestOffset() {
	this.So(Point{0, 0}.Offset(42, -123), should.Resemble, NewPoint(42, -123))
}

func (this *PointFixture) TestNeighbors4() {
	center := NewPoint(0, 0)
	eight := center.Neighbors8()

	this.So(eight[:4], should.Resemble, center.Neighbors4())

	this.So(eight[0], should.Resemble, Point{x: 0 + 1, y: 0 + 0}) // right
	this.So(eight[1], should.Resemble, Point{x: 0 - 1, y: 0 + 0}) // left
	this.So(eight[2], should.Resemble, Point{x: 0 + 0, y: 0 + 1}) // top
	this.So(eight[3], should.Resemble, Point{x: 0 + 0, y: 0 - 1}) // bottom

	this.So(eight[4], should.Resemble, Point{x: 0 + 1, y: 0 + 1}) // top-right
	this.So(eight[5], should.Resemble, Point{x: 0 - 1, y: 0 + 1}) // top-left
	this.So(eight[6], should.Resemble, Point{x: 0 + 1, y: 0 - 1}) // bottom-right
	this.So(eight[7], should.Resemble, Point{x: 0 - 1, y: 0 - 1}) // bottom-left
}
