package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestRPGFixture(t *testing.T) {
	suite.Run(&RPGFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type RPGFixture struct {
	*suite.T
}

func (this *RPGFixture) TestFight() {
	player := NewCharacter("player", 8, 5, 5)
	enemy := NewCharacter("boss", 12, 7, 2)
	winner := Fight(player, enemy)
	this.So(winner, should.Equal, player)
	this.So(player.hits, should.Equal, 2)
	this.So(enemy.hits, should.Equal, 0)
}
