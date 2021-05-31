package part2

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
	this.world.Set(Point(0, 1, 0, 0), true)
	this.world.Set(Point(0, 2, 1, 0), true)
	this.world.Set(Point(0, 0, 2, 0), true)
	this.world.Set(Point(0, 1, 2, 0), true)
	this.world.Set(Point(0, 2, 2, 0), true)
	this.AssertActive(5)
}
func (this *ConwayFixture) AssertActive(expected int) {
	this.So(this.world, should.HaveLength, expected)
}
func (this *ConwayFixture) TestParseInitialWorld() {
	world := ParseInitialWorld(".#.\n..#\n###")
	this.So(world, should.Resemble, this.world)
}
func (this *ConwayFixture) TestNeighbors4d() {
	unique := make(map[P]struct{})
	for _, p := range Point(0, 0, 0, 0).Neighbors4d() {
		unique[p] = struct{}{}
	}
	this.So(unique, should.HaveLength, 80)
}
func (this *ConwayFixture) TestCountActiveNeighbors() {
	neighbors := this.world.countActiveNeighbors(Point(0, 1, 1, 0))
	this.So(neighbors, should.Equal, 5)
}
func (this *ConwayFixture) TestOneCycle() {
	this.world.Cycle()
	this.AssertActive(29)
}
func (this *ConwayFixture) TestSixCycles() {
	this.world.Boot()
	this.AssertActive(848)
}