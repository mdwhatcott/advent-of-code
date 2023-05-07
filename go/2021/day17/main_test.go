package day17

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/intgrid"
)

func TestDay17Suite(t *testing.T) {
	should.Run(&Day17Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Day17Suite struct {
	*should.T
	start       string
	conversions []string
}

// sample: target area: x=20..30, y=-10..-5
// input:  target area: x=60..94, y=-171..-136
var (
	exampleTarget = intgrid.NewBoundingBox(intgrid.NewPoint(20, -5), intgrid.NewPoint(30, -10))
	realTarget    = intgrid.NewBoundingBox(intgrid.NewPoint(60, -136), intgrid.NewPoint(94, -171))
)

func (this *Day17Suite) TestGravity() {
	probe := NewProbe(intgrid.NewDirection(0, 0))

	probe.Advance()
	this.So(probe.Location, should.Equal, intgrid.NewPoint(0, 0))

	probe.Advance()
	this.So(probe.Location, should.Equal, intgrid.NewPoint(0, -1))

	probe.Advance()
	this.So(probe.Location, should.Equal, intgrid.NewPoint(0, -3))

	probe.Advance()
	this.So(probe.Location, should.Equal, intgrid.NewPoint(0, -6))
}
func (this *Day17Suite) TestPositiveDrag() {
	probe := NewProbe(intgrid.NewDirection(5, 0))

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5+4)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5+4+3)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5+4+3+2)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5+4+3+2+1)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5+4+3+2+1+0)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, 5+4+3+2+1+0+0)
}
func (this *Day17Suite) TestNegativeDrag() {
	probe := NewProbe(intgrid.NewDirection(-5, 0))

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5-4)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5-4-3)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5-4-3-2)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5-4-3-2-1)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5-4-3-2-1-0)

	probe.Advance()
	this.So(probe.Location.X(), should.Equal, -5-4-3-2-1-0-0)
}
func (this *Day17Suite) TestHighestElevation() {
	probe := NewProbe(intgrid.NewDirection(0, 10))
	probe.Advance()
	for probe.Location.Y() > 0 {
		probe.Advance()
	}
	this.So(probe.MaxY, should.Equal, 10+9+8+7+6+5+4+3+2+1)
}
func (this *Day17Suite) TestLaunchWithGoodAim() {
	probe := NewProbe(intgrid.NewDirection(7, 2))
	didHitTarget := Launch(probe, exampleTarget)
	this.So(didHitTarget, should.Equal, true)
	this.So(probe.MaxY, should.Equal, 3)
}
func (this *Day17Suite) TestLaunchWithBadAim() {
	probe := NewProbe(intgrid.NewDirection(17, -4))
	didHitTarget := Launch(probe, exampleTarget)
	this.So(didHitTarget, should.Equal, false)
	this.So(probe.MaxY, should.Equal, 0)
}
func (this *Day17Suite) TestSolutions() {
	part1HighestElevation, part2HitCount := SimulateLaunches(realTarget)
	this.So(part1HighestElevation, should.Equal, 14535)
	this.So(part2HitCount, should.Equal, 2270)
}
