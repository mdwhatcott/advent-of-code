package day05_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day05"
	"github.com/mdwhatcott/testing/should"
)

func TestDay05(t *testing.T) {
	should.So(t, day05.Part1(), should.Equal, 351282)
	should.So(t, day05.Part2(), should.Equal, 24568703)
}
