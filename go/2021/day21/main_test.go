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

func (this *DiracDiceSuite) TestExample() {
	game := NewDeterministicGame(4, 8)
	for game.Turn() {
	}
	this.So(game.Answer(), should.Equal, 739785)
}
func (this *DiracDiceSuite) TestPart1() {
	game := NewDeterministicGame(6, 3)
	for game.Turn() {
	}
	this.So(game.Answer(), should.Equal, 752745)
}
