package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestShuntingYardSuite(t *testing.T) {
	suite.Run(&ShuntingYardSuite{T: suite.New(t)}, suite.Options.UnitTests())
}

type ShuntingYardSuite struct{ *suite.T }

func (this *ShuntingYardSuite) parse1(input string) string {
	return ParseShuntingYard(part1Precedence, input)
}
func (this *ShuntingYardSuite) parse2(input string) string {
	return ParseShuntingYard(part2Precedence, input)
}

func (this *ShuntingYardSuite) eval1(input string) int {
	return EvalPostfix(this.parse1(input))
}
func (this *ShuntingYardSuite) eval2(input string) int {
	return EvalPostfix(this.parse2(input))
}

func (this *ShuntingYardSuite) TestParseShuntingYard() {
	this.So(this.parse1("3"), should.Equal, "3")
	this.So(this.parse1("3 + 4"), should.Equal, "34+")

	this.So(this.parse1("3 * (4 + 5) * 6"), should.Equal, "345+*6*")
	this.So(this.parse2("3 *  4 + 5  * 6"), should.Equal, "345+*6*")

	this.So(this.parse1("3 * 2 + 7"), should.Equal, "32*7+")
	this.So(this.parse2("3 * 2 + 7"), should.Equal, "327+*")

	this.So(this.parse1("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, "12+3*4+5*6+")
	this.So(this.parse2("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, "12+34+*56+*")
}
func (this *ShuntingYardSuite) TestEvalPostfix() {
	this.So(EvalPostfix("34+"), should.Equal, 7)
	this.So(EvalPostfix("56+7*8+"), should.Equal, 85)
	this.So(EvalPostfix("795+55+6**72+4+2++*9+"), should.Equal, 5994)
}

func (this *ShuntingYardSuite) TestPart1() {
	this.So(this.eval1("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, 71)
	this.So(this.eval1("1 + (2 * 3) + (4 * (5 + 6))"), should.Equal, 51)
	this.So(this.eval1("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), should.Equal, 13632)
	this.So(Part1(), should.Equal, 6923486965641)
}
func (this *ShuntingYardSuite) TestPart2() {
	this.So(this.eval2("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, 231)
	this.So(Part2(), should.Equal, 70722650566361)
}
