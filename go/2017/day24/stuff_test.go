package day24

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
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
