package day14_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day14"
	"github.com/mdwhatcott/testing/should"
)

func TestDay14(t *testing.T) {
	should.So(t, day14.Part1(), should.Equal, 8222)
	should.So(t, day14.Part2(), should.Equal, 1086)
}
