package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestConwayFixture(t *testing.T) {
	gunit.Run(new(ConwayFixture), t)
}

type ConwayFixture struct {
	*gunit.Fixture
	world World
}

func (this *ConwayFixture) Setup() {
	this.world = make(World)
	this.world.Set(Point(1, 0, 0), true)
	this.world.Set(Point(2, 1, 0), true)
	this.world.Set(Point(0, 2, 0), true)
	this.world.Set(Point(1, 2, 0), true)
	this.world.Set(Point(2, 2, 0), true)
	this.AssertActive(5)
}
func (this *ConwayFixture) AssertActive(expected int) {
	this.So(this.world, should.HaveLength, expected)
}
func (this *ConwayFixture) TestParseInitialWorld() {
	world := ParseInitialWorld(".#.\n..#\n###")
	this.So(world, should.Resemble, this.world)
}
func (this *ConwayFixture) TestNeighbors3d() {
	unique := make(map[P]struct{})
	for _, p := range Point(0, 0, 0).Neighbors3d() {
		unique[p] = struct{}{}
	}
	this.So(unique, should.HaveLength, 26)
}
func (this *ConwayFixture) TestCountActiveNeighbors() {
	neighbors := this.world.countActiveNeighbors(Point(1, 1, 0))
	this.So(neighbors, should.Equal, 5)
}
func (this *ConwayFixture) TestOneCycle() {
	this.world.Cycle()
	this.AssertActive(11)
}
func (this *ConwayFixture) TestSixCycles() {
	for x := 0; x < 6; x++ {
		this.world.Cycle()
	}
	this.AssertActive(112)
}
