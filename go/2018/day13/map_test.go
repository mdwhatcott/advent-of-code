package day13

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
}

func (this *StuffFixture) SkipTestA() {
	m := NewMap(TestInputA)
	go func() {
		m.Tick()
		m.Tick()
	}()
	this.So(<-m.Signals, should.Resemble, NewPoint(0, 3))
	this.So(m.String(), should.Equal, "|\n|\n|\nX\n|\n|\n|")
}
func (this *StuffFixture) TestB() {
	m := NewMap(TestInputB)
	go func() {
		for m.Tick() {
		}
	}()
	this.So(<-m.Signals, should.Resemble, NewPoint(7, 3))
}

var TestInputA = strings.TrimLeft(`
|
v
|
|
|
^
|`, "\n")

var TestInputB = strings.TrimLeft(`
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `, "\n")
