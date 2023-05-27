package day13

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuffFixture(t *testing.T) {
	should.Run(&StuffFixture{T: should.New(t)}, should.Options.UnitTests())
}

type StuffFixture struct {
	*should.T
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
func (this *StuffFixture) SkipTestB() {
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
