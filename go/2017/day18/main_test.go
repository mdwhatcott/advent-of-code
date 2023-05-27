package day18_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day18"
	"github.com/mdwhatcott/testing/should"
)

func TestDay18(t *testing.T) {
	should.So(t, day18.Part1(), should.Equal, 3423)
	should.So(t, day18.Part2(), should.Equal, 7493)
}
