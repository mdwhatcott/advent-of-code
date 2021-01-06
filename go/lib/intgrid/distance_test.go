package intgrid

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestDistanceFixture(t *testing.T) {
	gunit.Run(new(DistanceFixture), t)
}

type DistanceFixture struct {
	*gunit.Fixture
}

func (this *DistanceFixture) TestCityBlockDistance() {
	center := NewPoint(0, 0)
	this.So(CityBlockDistance(center, NewPoint(3, 4)), should.Equal, 7)
	this.So(CityBlockDistance(center, NewPoint(-3, 4)), should.Equal, 7)
	this.So(CityBlockDistance(center, NewPoint(3, -4)), should.Equal, 7)
	this.So(CityBlockDistance(center, NewPoint(-3, -4)), should.Equal, 7)
}

func (this *DistanceFixture) TestEuclideanDistance() {
	center := NewPoint(0, 0)
	this.So(EuclideanDistance(center, NewPoint(3, 4)), should.Equal, 5)
	this.So(EuclideanDistance(center, NewPoint(-3, 4)), should.Equal, 5)
	this.So(EuclideanDistance(center, NewPoint(3, -4)), should.Equal, 5)
	this.So(EuclideanDistance(center, NewPoint(-3, -4)), should.Equal, 5)
}
