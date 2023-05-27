package day08_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day08"
	"github.com/mdwhatcott/testing/should"
)

func TestDay08(t *testing.T) {
	part1, part2 := day08.Answers()
	should.So(t, part1, should.Equal, 4902)
	should.So(t, part2, should.Equal, 7037)
}
