package day06_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day06"
	"github.com/mdwhatcott/testing/should"
)

func TestDay06(t *testing.T) {
	should.So(t, day06.Part1(), should.Equal, 5042)
	should.So(t, day06.Part2(), should.Equal, 1086)
}
