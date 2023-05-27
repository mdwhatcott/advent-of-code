package day01_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day01"
	"github.com/mdwhatcott/testing/should"
)

func TestDay01(t *testing.T) {
	should.So(t, day01.Part1(), should.Equal, 1251)
	should.So(t, day01.Part2(), should.Equal, 1244)
}
