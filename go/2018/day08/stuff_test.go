package day08

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/suite"

	"advent/lib/util"

	"github.com/mdwhatcott/testing/should"
)

func TestStuffFixture(t *testing.T) {
	suite.Run(&StuffFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type StuffFixture struct {
	*suite.T
}

const toy = `2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`

//           A----------------------------------
//               B----------- C-----------
//                                D-----

func prepare(input string) []int {
	return util.ParseInts(strings.Fields(input))
}

func (this *StuffFixture) TestPart1() {
	this.So(SumMetadata(ParseTree(LoadInputs(prepare(toy)))), should.Equal, 138)            // A
	this.So(SumMetadata(ParseTree(LoadInputs(prepare("0 3 10 11 12")))), should.Equal, 33)  // B
	this.So(SumMetadata(ParseTree(LoadInputs(prepare("1 1 0 1 99 2")))), should.Equal, 101) // C
}

func (this *StuffFixture) TestPart2() {
	this.So(RootValue(ParseTree(LoadInputs(prepare(toy)))), should.Equal, 66)            // A
	this.So(RootValue(ParseTree(LoadInputs(prepare("0 3 10 11 12")))), should.Equal, 33) // B
	this.So(RootValue(ParseTree(LoadInputs(prepare("1 1 0 1 99 2")))), should.Equal, 0)  // C
	this.So(RootValue(ParseTree(LoadInputs(prepare("0 1 99")))), should.Equal, 99)       // D
}
