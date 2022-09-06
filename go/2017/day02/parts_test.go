package day02

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuff(t *testing.T) {
	should.Run(&Stuff{T: should.New(t)}, should.Options.UnitTests())
}

type Stuff struct {
	*should.T
}

func (this *Stuff) Test() {
	this.So(part1Checksum(strings.Split("5 1 9 5\n7 5 3\n2 4 6 8", "\n")), should.Equal, 18)
	this.So(part2Checksum(strings.Split("5 9 2 8\n9 4 7 3\n3 8 6 5", "\n")), should.Equal, 9)
}
