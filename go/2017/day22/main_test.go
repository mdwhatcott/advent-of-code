package day22_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day22"
	"github.com/mdwhatcott/testing/should"
)

func TestDay22(t *testing.T) {
	should.So(t, day22.Part1(), should.Equal, 5460)
	should.So(t, day22.Part2(), should.Equal, 2511702)
}
