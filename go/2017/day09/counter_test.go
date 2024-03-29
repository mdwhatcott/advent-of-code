package day09

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuff(t *testing.T) {
	should.Run(&Stuff{T: should.New(t)}, should.Options.UnitTests())
}

type Stuff struct {
	*should.T
}

func (this *Stuff) TestGroupScoring() {
	this.So(groupScore("{}"), should.Equal, 1)
	this.So(groupScore("{{}}"), should.Equal, 3)
	this.So(groupScore("{{},{}}"), should.Equal, 5)
	this.So(groupScore("{{{},{},{{}}}}"), should.Equal, 16)
	this.So(groupScore("{<a>,<a>,<a>,<a>}"), should.Equal, 1)
	this.So(groupScore("{{<ab>},{<ab>},{<ab>},{<ab>}}"), should.Equal, 9)
	this.So(groupScore("{{<!!>},{<!!>},{<!!>},{<!!>}}"), should.Equal, 9)
	this.So(groupScore("{{<a!>},{<a!>},{<a!>},{<ab>}}"), should.Equal, 3)
}

func (this *Stuff) TestGarbageScoring() {
	this.So(garbageScore("<>"), should.Equal, 0)
	this.So(garbageScore("<random characters>"), should.Equal, 17)
	this.So(garbageScore("<<<<>"), should.Equal, 3)
	this.So(garbageScore("<{!>}>"), should.Equal, 2)
	this.So(garbageScore("<!!>"), should.Equal, 0)
	this.So(garbageScore("<!!!>>"), should.Equal, 0)
	this.So(garbageScore(`<{o"i!a,<{i<a>`), should.Equal, 10)

}
