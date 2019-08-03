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

func (this *BattleFixture) TestBossAttackDoesSpecifiedDamageToPlayer() {
	before := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 100,
		PlayerArmor:     0,
		PlayerMana:      1,
		BossHitPoints:   2,
		BossDamage:      30,
		ShieldCounter:   4,
		PoisonCounter:   5,
		RechargeCounter: 6,
	}
	after := before.Handle(new(BossAttack))

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 70,
		PlayerArmor:     0,
		PlayerMana:      1,
		BossHitPoints:   2,
		BossDamage:      30,
		ShieldCounter:   4,
		PoisonCounter:   5,
		RechargeCounter: 6,
	})

}

func (this *BattleFixture) Test_Missile_Does4DamageToBoss_CostsPlayer53Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3+53,
		BossHitPoints:   8,
		BossDamage:      5,
		ShieldCounter:   6,
		PoisonCounter:   7,
		RechargeCounter: 8,
	}
	after := before.Handle(new(Missile))

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3, // was 3+53
		BossHitPoints:   4, // was 8
		BossDamage:      5,
		ShieldCounter:   6,
		PoisonCounter:   7,
		RechargeCounter: 8,
	})
}

func (this *BattleFixture) Test_Drain_Deals2DamageToBoss_HealsPlayerBy2_CostsPlayer73Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3+73,
		BossHitPoints:   6,
		BossDamage:      5,
		ShieldCounter:   6,
		PoisonCounter:   7,
		RechargeCounter: 8,
	}
	after := before.Handle(new(Drain))

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 3, // was 1
		PlayerArmor:     2,
		PlayerMana:      3,
		BossHitPoints:   4, // was 6
		BossDamage:      5,
		ShieldCounter:   6,
		PoisonCounter:   7,
		RechargeCounter: 8,
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
