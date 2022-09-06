package grid

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestDistanceFixture(t *testing.T) {
	should.Run(&DistanceFixture{T: should.New(t)}, should.Options.UnitTests())
}

type DistanceFixture struct {
	*should.T
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
