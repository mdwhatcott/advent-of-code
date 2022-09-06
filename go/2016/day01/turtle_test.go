package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestTurtleFixture(t *testing.T) {
	should.Run(&TurtleFixture{T: should.New(t)}, should.Options.UnitTests())
}

type TurtleFixture struct {
	*should.T

	turtle *Turtle
}

func (this *TurtleFixture) Setup() {
	this.turtle = NewTurtle()
}

func (this *TurtleFixture) Test1() {
	this.turtle.Follow("R2")
	this.So(this.turtle.Position(), should.Equal, "2,0")
	this.So(this.turtle.TaxiDistanceToEndingLocation(), should.Equal, 2)

	this.turtle.Follow("L3")
	this.So(this.turtle.Position(), should.Equal, "2,3")
	this.So(this.turtle.TaxiDistanceToEndingLocation(), should.Equal, 5)
}

func (this *TurtleFixture) Test2() {
	this.turtle.Follow("R2")
	this.turtle.Follow("R2")
	this.turtle.Follow("R2")

	this.So(this.turtle.Position(), should.Equal, "0,-2")
	this.So(this.turtle.TaxiDistanceToEndingLocation(), should.Equal, 2)
}

func (this *TurtleFixture) Test3() {
	this.turtle.FollowAll("R5, L5, R5, R3")
	this.So(this.turtle.TaxiDistanceToEndingLocation(), should.Equal, 12)
}

func (this *TurtleFixture) Test4_MultiDigitSteps() {
	this.turtle.Follow("R10")
	this.So(this.turtle.Position(), should.Equal, "10,0")
}

func (this *TurtleFixture) Test_DistanceToFirstPointVisitedTwice() {
	this.turtle.Follow("R8")
	this.So(this.turtle.PositionFirstVisitedTwice(), should.Equal, "")

	this.turtle.FollowAll("R4, R4, R8")
	this.So(this.turtle.PositionFirstVisitedTwice(), should.Equal, "4,0")
}
