package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestTrigFixture(t *testing.T) {
	gunit.Run(new(TrigFixture), t)
}

type TrigFixture struct {
	*gunit.Fixture
}

func (this *TrigFixture) TestAngleFromOrigin() {
	this.So(NewAsteroid(0, -1).AngleFromOrigin(), should.Equal, 0)
	this.So(NewAsteroid(1, 0).AngleFromOrigin(), should.Equal, 90)
	this.So(NewAsteroid(0, 1).AngleFromOrigin(), should.Equal, 180)
	this.So(NewAsteroid(-1, 0).AngleFromOrigin(), should.Equal, 270)

	this.So(NewAsteroid(3, -3).AngleFromOrigin(), should.Equal, 0+45)
	this.So(NewAsteroid(3, 3).AngleFromOrigin(), should.Equal, 90+45)
	this.So(NewAsteroid(-3, 3).AngleFromOrigin(), should.Equal, 180+45)
	this.So(NewAsteroid(-3, -3).AngleFromOrigin(), should.Equal, 270+45)

	this.So(NewAsteroid(3, -4).AngleFromOrigin(), should.AlmostEqual, 36.87, .01)
	this.So(NewAsteroid(4, 3).AngleFromOrigin(), should.AlmostEqual, 90+36.87, .01)
	this.So(NewAsteroid(-3, 4).AngleFromOrigin(), should.AlmostEqual, 180+36.87, .01)
	this.So(NewAsteroid(-4, -3).AngleFromOrigin(), should.AlmostEqual, 270+36.87, .01)
}

func (this *TrigFixture) TestDistanceFromOrigin() {
	this.So(NewAsteroid(0, 1).DistanceFromOrigin(), should.Equal, 1)
	this.So(NewAsteroid(1, 0).DistanceFromOrigin(), should.Equal, 1)
	this.So(NewAsteroid(0, -1).DistanceFromOrigin(), should.Equal, 1)
	this.So(NewAsteroid(-1, 0).DistanceFromOrigin(), should.Equal, 1)

	this.So(NewAsteroid(3, 4).DistanceFromOrigin(), should.Equal, 5)
	this.So(NewAsteroid(-3, 4).DistanceFromOrigin(), should.Equal, 5)
	this.So(NewAsteroid(-3, -4).DistanceFromOrigin(), should.Equal, 5)
	this.So(NewAsteroid(3, -4).DistanceFromOrigin(), should.Equal, 5)
}
