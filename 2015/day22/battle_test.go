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

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		Shield,
		Poison,
		Recharge,
	})
}

func (this *BattleFixture) Test_BossTurn_BothStillAlive_Attack() {
	game := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerMana:      250,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []int{BossAttack})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForRecharge() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      228,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		Shield,
		Poison,
		//Recharge,
	})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForPoison() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      172,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		Shield,
		//Poison,
		//Recharge,
	})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForShield() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      112,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		//Shield,
		//Poison,
		//Recharge,
	})
}

func (this *BattleFixture) Test_PlayerTurn_NotEnoughManaForDraim() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerMana:      72,
		BossHitPoints:   1,
	}

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		//Drain,
		//Shield,
		//Poison,
		//Recharge,
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

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		//Shield,
		Poison,
		Recharge,
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

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		Shield,
		//Poison,
		Recharge,
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

	this.So(game.Attack(), should.Resemble, []int{
		Missile,
		Drain,
		Shield,
		Poison,
		//Recharge,
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
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	}
	after := before.Handle(BossAttack)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 70,
		PlayerArmor:     0,
		PlayerMana:      1,
		BossHitPoints:   2,
		BossDamage:      30,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	})

}

func (this *BattleFixture) Test_Missile_Does4DamageToBoss_CostsPlayer53Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3 + SpellCost[Missile],
		BossHitPoints:   8,
		BossDamage:      5,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	}
	after := before.Handle(Missile)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3, // was 3+53
		BossHitPoints:   4, // was 8
		BossDamage:      5,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	})
}

func (this *BattleFixture) Test_Drain_Deals2DamageToBoss_HealsPlayerBy2_CostsPlayer73Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3 + SpellCost[Drain],
		BossHitPoints:   6,
		BossDamage:      5,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	}
	after := before.Handle(Drain)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 3, // was 1
		PlayerArmor:     2,
		PlayerMana:      3,
		BossHitPoints:   4, // was 6
		BossDamage:      5,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	})
}

func (this *BattleFixture) Test_Poison_Deals3DamageToBoss_For6Turns_CostsPlayer73Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     1,
		PlayerMana:      3 + SpellCost[Poison],
		BossHitPoints:   100,
		BossDamage:      1,
		ShieldCounter:   1,
		PoisonCounter:   0,
		RechargeCounter: 1,
	}
	after := before.Handle(Poison)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     1,
		PlayerMana:      3,
		BossHitPoints:   97, // was 100
		BossDamage:      1,
		ShieldCounter:   1,
		PoisonCounter:   5,
		RechargeCounter: 1,
	})

	for x := 0; x < 50; x++ {
		after = after.Handle(-1)
	}

	this.So(after.BossHitPoints, should.Equal, 100-(3*6))
}

// TODO: armor may reduce damage to 1, but no lower

///////////////////////////////////////////////////////////////////

func (this *BattleFixture) TestExample1() {
	game := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 10,
		PlayerMana:      250,
		BossHitPoints:   13,
		BossDamage:      8,
	}

	moves := []int{
		Poison,
		BossAttack,
		Missile,
	}
	for _, turn := range moves {
		game = game.Handle(turn)
	}

	this.So(game, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 2,
		PlayerMana:      24,
		PoisonCounter:   3,
		BossDamage:      8,
		BossHitPoints:   0,
	})
}
