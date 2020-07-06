package advent

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
}

func (this *StuffFixture) Setup() {
}

func (this *StuffFixture) TestScan() {
	field := scanField(exampleMap1)
	this.So(field, should.Resemble, AsteroidField{
		{X: 1, Y: 0},
		{X: 4, Y: 0},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 4, Y: 3},
		{X: 3, Y: 4},
		{X: 4, Y: 4},
	})
}

func (this *StuffFixture) TestSlope() {
	a := NewAsteroid(0, 0)
	b := NewAsteroid(1, 1)
	c := NewAsteroid(2, 2)
	d := NewAsteroid(6, 7)

	this.So(Slope(a, b), should.Equal, 1)
	this.So(Slope(a, c), should.Equal, 1)
	this.So(Slope(c, b), should.Equal, 1)
	this.So(Slope(a, d), should.Equal, 6.0/7.0)
}

func (this *StuffFixture) TestCountVisible() {
	field := scanField(exampleMap1)

	this.So(CountVisible(field, field[0]), should.Equal, 7)
}

func (this *StuffFixture) TestBestPlace() {
	field := scanField(exampleMap1)

	this.So(BestPlace(field), should.Resemble, NewAsteroid(3, 4))
}

var (
	exampleMap1 = strings.Split(strings.TrimSpace(""+
		".#..#\n"+
		".....\n"+
		"#####\n"+
		"....#\n"+
		"...##\n",
	), "\n")
)
