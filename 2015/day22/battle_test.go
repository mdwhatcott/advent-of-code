package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(BattleFixture), t)
}

type BattleFixture struct {
	*gunit.Fixture
}

func (this *BattleFixture) Test_PlayerTurn_ButIsDead_NoMoves() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 0,
		PlayerMana:      250,
		BossHitPoints:   1,
	}
	this.So(game.Attack(), should.BeEmpty)
}

func (this *BattleFixture) Test_PlayerTurn_ButManaIsScarce_NoMoves() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      52,
		BossHitPoints:   1,
	}
	this.So(game.Attack(), should.BeEmpty)
}

func (this *BattleFixture) Test_BossTurn_ButIsDead_NoMoves() {
	game := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerMana:      250,
		BossHitPoints:   0,
	}
	this.So(game.Attack(), should.BeEmpty)
}

func (this *BattleFixture) Test_PlayerTurn_BothStillAlive_NoEffectsInPlay_PlentyOfMana__AllSpellsAvailable() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      250,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		new(Shield),
		new(Poison),
		new(Recharge),
	})
}

func (this *BattleFixture) Test_BossTurn_BothStillAlive_Attack() {
	game := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerMana:      250,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{new(BossAttack)})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForRecharge() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      228,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		new(Shield),
		new(Poison),
		//new(Recharge),
	})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForPoison() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      172,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		new(Shield),
		//new(Poison),
		//new(Recharge),
	})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForShield() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      112,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		//new(Shield),
		//new(Poison),
		//new(Recharge),
	})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForDraim() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      72,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		//new(Drain),
		//new(Shield),
		//new(Poison),
		//new(Recharge),
	})
}

func (this *BattleFixture) TestShieldAlreadyInPlace_CannotRecastShield() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      1000,
		BossHitPoints:   1,
		ShieldCounter:   2,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		//new(Shield),
		new(Poison),
		new(Recharge),
	})
}

func (this *BattleFixture) TestPoisonAlreadyInPlace_CannotRecastPoison() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      1000,
		BossHitPoints:   1,
		PoisonCounter:   2,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		new(Shield),
		//new(Poison),
		new(Recharge),
	})
}

func (this *BattleFixture) TestRechargeAlreadyInPlace_CannotRecastRecharge() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      1000,
		BossHitPoints:   1,
		RechargeCounter: 2,
	}

	this.So(game.Attack(), should.Resemble, []interface{}{
		new(Missile),
		new(Drain),
		new(Shield),
		new(Poison),
		//new(Recharge),
	})
}

///////////////////////////////////////////////////////////////////

func (this *BattleFixture) SkipTestExample1() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 10,
		PlayerMana:      250,
		BossHitPoints:   13,
		BossDamage:      8,
	}

	moves := []interface{}{
		new(Poison),
		new(BossAttack),
		new(Missile),
	}
	for _, turn := range moves {
		game = game.Handle(turn)
	}

	this.So(game, should.Resemble, Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 2,
		PlayerMana:      77,
		PoisonCounter:   4,
		BossDamage:      8,
	})
}
