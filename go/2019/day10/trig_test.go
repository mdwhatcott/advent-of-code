package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestTrigFixture(t *testing.T) {
	suite.Run(&TrigFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type TrigFixture struct {
	*suite.T
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

	this.So(NewAsteroid(3, -4).AngleFromOrigin(), should.Equal, 36.86989764584402)
	this.So(NewAsteroid(4, 3).AngleFromOrigin(), should.Equal, 90+36.86989764584402)
	this.So(NewAsteroid(-3, 4).AngleFromOrigin(), should.Equal, 180+36.86989764584402)
	this.So(NewAsteroid(-4, -3).AngleFromOrigin(), should.Equal, 270+36.86989764584405)
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
