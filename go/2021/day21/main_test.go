package day21

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestDiracDiceSuite(t *testing.T) {
	should.Run(&DiracDiceSuite{T: should.New(t)}, should.Options.UnitTests())
}

type DiracDiceSuite struct {
	*should.T
}

func (this *DiracDiceSuite) TestExample1() {
	this.So(NewDeterministicGame(4, 8).Play(), should.Equal, 739785)
}
func (this *DiracDiceSuite) TestPart1() {
	this.So(NewDeterministicGame(6, 3).Play(), should.Equal, 752745)
}
func (this *DiracDiceSuite) TestExample2() {
	this.So(PlayDirac(4, 8), should.Equal, 444356092776315)
}
func (this *DiracDiceSuite) TestPart2() {
	this.So(PlayDirac(6, 3), should.Equal, 309196008717909)
}
