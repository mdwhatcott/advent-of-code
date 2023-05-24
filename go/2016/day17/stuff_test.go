package main

import (
	"crypto/md5"
	"hash"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func TestDoorHashMazeFixture(t *testing.T) {
	should.Run(&DoorHashMazeFixture{T: should.New(t)}, should.Options.UnitTests())
}

type DoorHashMazeFixture struct {
	*should.T
	hasher   hash.Hash
	passcode string
}

func (this *DoorHashMazeFixture) Setup() {
	this.hasher = md5.New()
	this.passcode = "hijkl"
}

func (this *DoorHashMazeFixture) TestOpenDoors() {
	start := NewLocation(1, 1, "")
	adjacents := start.AdjacentOpenRooms(this.passcode, this.hasher)
	this.So(adjacents, should.Equal, []*Location{{X: 1, Y: 2, Directions: "D"}})

	next := adjacents[0]
	adjacents = next.AdjacentOpenRooms(this.passcode, this.hasher)
	this.So(adjacents, should.Equal, []*Location{
		NewLocation(1, 1, "DU"),
		NewLocation(2, 2, "DR"),
	})
}

func (this *DoorHashMazeFixture) assertResult(passcode string, shortPath string, longDistance int) {
	short, long := Navigate(passcode)
	this.So(short, should.Equal, shortPath)
	this.So(long, should.Equal, longDistance)
}
func (this *DoorHashMazeFixture) TestCases() {
	this.assertResult("ihgpwlah", "DDRRRD", 370)
	this.assertResult("kglvqrro", "DDUDRLRRUDRD", 492)
	this.assertResult("ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR", 830)
}
func (this *DoorHashMazeFixture) TestSolve() {
	shortestPath, longestDistance := Navigate(util.InputString())
	this.So(shortestPath, should.Equal, "RLRDRDUDDR")
	this.So(longestDistance, should.Equal, 420)
}
