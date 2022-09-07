package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/util"
)

func TestExample(t *testing.T) {
	origin := intgrid.NewPoint(1, 1)
	target := intgrid.NewPoint(7, 4)
	distance, _ := BreadthFirstSearch(10, origin, target)
	should.So(t, distance, should.Equal, 11)
}

func TestSolve(t *testing.T) {
	origin := intgrid.NewPoint(1, 1)
	target := intgrid.NewPoint(31, 39)
	distance, near := BreadthFirstSearch(util.InputInt(), origin, target)
	should.So(t, distance, should.Equal, 96)
	should.So(t, near, should.Equal, 141)
}
