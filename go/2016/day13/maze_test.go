package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
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
func TestBinaryHammingWeight(t *testing.T) {
	should.So(t, binaryHammingWeight(0), should.Equal, 0)
	should.So(t, binaryHammingWeight(1), should.Equal, 1)
	should.So(t, binaryHammingWeight(2), should.Equal, 1)
	should.So(t, binaryHammingWeight(3), should.Equal, 2)
	should.So(t, binaryHammingWeight(255), should.Equal, 8)
}
