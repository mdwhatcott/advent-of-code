package day02

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	should.So(t, Part1(), should.Equal, 10994)
	should.So(t, Part2(), should.Equal, 12526)
}
