package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestRPGFixture(t *testing.T) {
	should.Run(&RPGFixture{T: should.New(t)}, should.Options.UnitTests())
}

type RPGFixture struct {
	*should.T
}

func (this *RPGFixture) TestFight() {
	player := NewCharacter("player", 8, 5, 5)
	enemy := NewCharacter("boss", 12, 7, 2)
	winner := Fight(player, enemy)
	this.So(winner, should.Equal, player)
	this.So(player.hits, should.Equal, 2)
	this.So(enemy.hits, should.Equal, 0)
}
