package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestShuntingYardSuite(t *testing.T) {
	suite.Run(&ShuntingYardSuite{T: suite.New(t)}, suite.Options.UnitTests())
}

type ShuntingYardSuite struct {
	*suite.T
}

func (this *ShuntingYardSuite) shunt1(input string) string {
	return string(ParseShuntingYard(part1Precedence, input))
}
func (this *ShuntingYardSuite) shunt2(input string) string {
	return string(ParseShuntingYard(part2Precedence, input))
}
func (this *ShuntingYardSuite) part1(input string) int {
	return EvalPostfix(this.shunt1(input))
}
func (this *ShuntingYardSuite) part2(input string) int {
	return EvalPostfix(this.shunt2(input))
}
func (this *ShuntingYardSuite) TestShuntingYard() {
	this.So(this.shunt1("3"), should.Equal, "3")
	this.So(this.shunt1("3 + 4"), should.Equal, "34+")
	this.So(this.shunt1("3 * (4 + 5) * 6"), should.Equal, "345+*6*")

	this.So(this.shunt1("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, "")
	this.So(this.shunt2("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, "")
}

func (this *ShuntingYardSuite) TestEvalPostfix() {
	this.So(EvalPostfix("34+"), should.Equal, 7)
	this.So(EvalPostfix("56+7*8+"), should.Equal, 85)
}

func (this *ShuntingYardSuite) TestPart1() {
	this.So(this.part1("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, 71)
	this.So(this.part1("1 + (2 * 3) + (4 * (5 + 6))"), should.Equal, 51)
	this.So(this.part1("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), should.Equal, 13632)

	this.So(this.part2("1 + 2 * 3 + 4 * 5 + 6"), should.Equal, 231)
}
