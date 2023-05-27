package day13_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day13"
	"github.com/mdwhatcott/testing/should"
)

func TestDay13(t *testing.T) {
	part1, part2 := day13.Answers()
	should.So(t, part1, should.Equal, 788)
	should.So(t, part2, should.Equal, 3905748)
}
