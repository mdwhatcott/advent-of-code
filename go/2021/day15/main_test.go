package day15

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

var exampleLines = []string{
	"1163751742",
	"1381373672",
	"2136511328",
	"3694931569",
	"7463417111",
	"1319128137",
	"1359912421",
	"3125421639",
	"1293138521",
	"2311944581",
}

func TestPart1(t *testing.T) {
	assert.Fatal(t).So(Part1(exampleLines), should.Equal, 40)
	assert.Fatal(t).So(Part1(util.InputLines()), should.Equal, 824)
}

func TestPart2(t *testing.T) {
	assert.Fatal(t).So(Part2(exampleLines), should.Equal, 315)
	assert.Fatal(t).So(Part2(util.InputLines()), should.Equal, 3063)
}
