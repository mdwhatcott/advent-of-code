package day25_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day25"
	"github.com/mdwhatcott/testing/should"
)

func TestDay25(t *testing.T) {
	should.So(t, day25.Part1(), should.Equal, 633)
}
