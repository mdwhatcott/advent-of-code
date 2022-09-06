package day24

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
	lines := strings.Split(`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, "\n")
	this.So(FindStrongestBridge(lines), should.Equal, 31)
}
