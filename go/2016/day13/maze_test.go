package main

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/util"
)

func TestExample(t *testing.T) {
	distance, _ := BreadthFirstSearch(10, intgrid.NewPoint(7, 4))
	assert.Error(t).So(distance, should.Equal, 11)
}

func TestSolve(t *testing.T) {
	target := intgrid.NewPoint(31, 39)
	distance, near := BreadthFirstSearch(util.InputInt(), target)
	assert.Error(t).So(distance, should.Equal, 96)
	assert.Error(t).So(near, should.Equal, 141)
}
