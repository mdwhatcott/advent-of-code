package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestReindeerFixture(t *testing.T) {
	gunit.Run(new(ReindeerFixture), t)
}

type ReindeerFixture struct {
	*gunit.Fixture

	simulator *Simulator
	comet     Reindeer
	dancer    Reindeer
}

func (this *ReindeerFixture) TestParseReindeerLine() {
	reindeer := ParseReindeer("Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.")
	this.So(reindeer, should.Resemble, Reindeer{
		Name:     "Comet",
		Velocity: 14,
		Sustain:  10,
		Rest:     127,
	})
}

// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
// Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
func (this *ReindeerFixture) Setup() {
	this.comet = Reindeer{
		Name:     "Comet",
		Velocity: 14,
		Sustain:  10,
		Rest:     127,
	}
	this.dancer = Reindeer{
		Name:     "Dancer",
		Velocity: 16,
		Sustain:  11,
		Rest:     162,
	}
	this.simulator = NewSimulator()
	this.simulator.Register(this.comet)
	this.simulator.Register(this.dancer)
}

func (this *ReindeerFixture) TestAfter1Second() {
	this.simulator.Tick(1)
	this.So(this.simulator.distances[this.comet], should.Equal, 14)
	this.So(this.simulator.distances[this.dancer], should.Equal, 16)
	this.So(this.simulator.points[this.comet], should.Equal, 0)
	this.So(this.simulator.points[this.dancer], should.Equal, 1)
}

func (this *ReindeerFixture) TestAfter10Seconds() {
	this.simulator.Tick(10)
	this.So(this.simulator.distances[this.comet], should.Equal, 140)
	this.So(this.simulator.distances[this.dancer], should.Equal, 160)
}

func (this *ReindeerFixture) TestAfter11Seconds_CometIsResting() {
	this.simulator.Tick(11)
	this.So(this.simulator.distances[this.comet], should.Equal, 140)
	this.So(this.simulator.distances[this.dancer], should.Equal, 176)
}

func (this *ReindeerFixture) TestAfter12Seconds_BothAreResting() {
	this.simulator.Tick(12)
	this.So(this.simulator.distances[this.comet], should.Equal, 140)
	this.So(this.simulator.distances[this.dancer], should.Equal, 176)
}

func (this *ReindeerFixture) TestAfter1000Seconds() {
	this.simulator.Tick(1000)
	this.So(this.simulator.distances[this.comet], should.Equal, 1120)
	this.So(this.simulator.distances[this.dancer], should.Equal, 1056)
	this.So(this.simulator.points[this.dancer], should.Equal, 689)
	this.So(this.simulator.points[this.comet], should.Equal, 312)
}
