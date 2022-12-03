package day01

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	topElf, top3Elves := sums()
	should.So(t, topElf, should.Equal, 71924)
	should.So(t, top3Elves, should.Equal, 210406)
}
