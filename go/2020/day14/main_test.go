package advent_test

import (
	"testing"

	day14 "github.com/mdwhatcott/advent-of-code/go/2020/day14"
	"github.com/mdwhatcott/testing/should"
)

func TestDay14(t *testing.T) {
	should.So(t, day14.Part1(), should.Equal, 18630548206046)
	should.So(t, day14.Part2(), should.Equal, 4254673508445)
}
