package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture
}

func (this *Stuff) Test_PlayerTurn_BothStillAlive_NoEffectsInPlay_PlentyOfMana__AllSpellsAvailable() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 10,
		PlayerMana:      250,
		BossHitPoints:   13,
	}

	this.So(game.NextMoves(), should.Resemble, []Turn{
		new(Missile),
		new(Drain),
		new(Shield),
		new(Poison),
		new(Recharge),
	})
}

func (this *Stuff) Test_BossTurn_BothStillAlive_Attack() {
	game := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 10,
		PlayerMana:      250,
		BossHitPoints:   13,
	}

	this.So(game.NextMoves(), should.Resemble, []Turn{new(BossAttack)})
}
