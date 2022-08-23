package day13

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestStuffFixture(t *testing.T) {
	suite.Run(&StuffFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type StuffFixture struct {
	*suite.T
}

func (this *StuffFixture) SkipTestA() {
	m := NewMap(TestInputA)
	go func() {
		m.Tick()
		m.Tick()
	}()
	this.So(<-m.Signals, should.Equal, NewPoint(0, 3))
	this.So(m.String(), should.Equal, "|\n|\n|\nX\n|\n|\n|")
}
func (this *StuffFixture) TestB() {
	m := NewMap(TestInputB)
	go func() {
		for m.Tick() {
		}
	}()
	this.So(<-m.Signals, should.Equal, NewPoint(7, 3))
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
