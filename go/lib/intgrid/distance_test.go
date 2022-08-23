package intgrid

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestDistanceFixture(t *testing.T) {
	suite.Run(&DistanceFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type DistanceFixture struct {
	*suite.T
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
