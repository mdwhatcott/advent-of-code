package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestRPGFixture(t *testing.T) {
	gunit.Run(new(RPGFixture), t)
}

type RPGFixture struct {
	*gunit.Fixture
}

func (this *RPGFixture) TestFight() {
	player := NewCharacter("player", 8, 5, 5)
	enemy := NewCharacter("boss", 12, 7, 2)
	winner := Fight(player, enemy)
	this.So(winner, should.Equal, player)
	this.So(player.hits, should.Equal, 2)
	this.So(enemy.hits, should.Equal, 0)
}
