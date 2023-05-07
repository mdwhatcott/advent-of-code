package day14

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func TestDay14Suite(t *testing.T) {
	should.Run(&Day14Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Day14Suite struct {
	*should.T
	start       string
	conversions []string
}

func (this *Day14Suite) Setup() {
	lines := util.InputLines()
	this.start = lines[0]
	this.conversions = lines[2:]
}

const sampleStart = "NNCB"

var sampleConversions = []string{
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}

func (this *Day14Suite) TestPart1() {
	this.So(Solve(sampleStart, sampleConversions, 10), should.Equal, 1588)
	this.So(Solve(this.start, this.conversions, 10), should.Equal, 2447)
}
func (this *Day14Suite) TestPart2() {
	this.So(Solve(sampleStart, sampleConversions, 40), should.Equal, 2188189693529)
	this.So(Solve(this.start, this.conversions, 40), should.Equal, 3018019237563) // 3629372415877
}
