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

func (this *BattleFixture) Test_PlayerTurn_BothAlive_NoEffectsInPlay_PlentyOfMana__AllSpellsAvailable() {
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

func (this *BattleFixture) Test_ShieldAlreadyInPlace_CannotRecastShield() {
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

func (this *BattleFixture) Test_PoisonAlreadyInPlace_CannotRecastPoison() {
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

func (this *BattleFixture) Test_RechargeAlreadyInPlace_CannotRecastRecharge() {
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

func (this *BattleFixture) Test_BossAttackDoesSpecifiedDamageToPlayer() {
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

func (this *BattleFixture) Test_Poison_SetsPoisonTimerFor6Turns_CostsPlayer73Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     0,
		PlayerMana:      3 + SpellCost[Poison],
		BossHitPoints:   100,
		BossDamage:      1,
		ShieldCounter:   -1,
		PoisonCounter:   0,
		RechargeCounter: -1,
	}
	after := before.Handle(Poison)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     0,
		PlayerMana:      3,
		BossHitPoints:   100,
		BossDamage:      1,
		ShieldCounter:   -1,
		PoisonCounter:   6,
		RechargeCounter: -1,
	})

	for x := 0; x < 50; x++ {
		after = after.Handle(-1)
	}

	this.So(after.BossHitPoints, should.Equal, 100-(3*6))
}

func (this *BattleFixture) Test_Shield_StartsShieldTimerFor6Turns_CostsPlayer113Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     0,
		PlayerMana:      3 + SpellCost[Shield],
		BossHitPoints:   1,
		BossDamage:      1,
		ShieldCounter:   0,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	}
	after := before.Handle(Shield)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     0,
		PlayerMana:      3,
		BossHitPoints:   1,
		BossDamage:      1,
		ShieldCounter:   6,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	})

	for x := 0; x < 50; x++ {
		after = after.Handle(-1)
	}

	this.So(after.PlayerArmor, should.Equal, 0)
}

func (this *BattleFixture) Test_Recharge_IncreasesManaBy101_For5Turns_CostsPlayer229Mana() {
	before := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 1,
		PlayerArmor:     -1,
		PlayerMana:      3 + SpellCost[Recharge],
		BossHitPoints:   1,
		BossDamage:      1,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: -1,
	}
	after := before.Handle(Recharge)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     -1,
		PlayerMana:      3,
		BossHitPoints:   1,
		BossDamage:      1,
		ShieldCounter:   -1,
		PoisonCounter:   -1,
		RechargeCounter: 5,
	})

	for x := 0; x < 50; x++ {
		after = after.Handle(-1)
	}

	this.So(after.PlayerMana, should.Equal, 3+(101*5))
}

type ArmorTest struct {
	PlayerHitPoints   int
	PlayerArmor       int
	BossDamage        int
	ExpectedHitPoints int
}

func (this *BattleFixture) Test_ArmorCanAbsorbsAlmostAllBossDamage() {
	this.assertDamageDealt(ArmorTest{
		PlayerHitPoints:   10,
		PlayerArmor:       0,
		BossDamage:        4,
		ExpectedHitPoints: 6,
	})
	this.assertDamageDealt(ArmorTest{
		PlayerHitPoints:   10,
		PlayerArmor:       1,
		BossDamage:        4,
		ExpectedHitPoints: 7,
	})
	this.assertDamageDealt(ArmorTest{
		PlayerHitPoints:   10,
		PlayerArmor:       2,
		BossDamage:        4,
		ExpectedHitPoints: 8,
	})
	this.assertDamageDealt(ArmorTest{
		PlayerHitPoints:   10,
		PlayerArmor:       3,
		BossDamage:        4,
		ExpectedHitPoints: 9,
	})
	this.assertDamageDealt(ArmorTest{
		PlayerHitPoints:   10,
		PlayerArmor:       4,
		BossDamage:        4,
		ExpectedHitPoints: 9,
	})
}
func (this *BattleFixture) assertDamageDealt(test ArmorTest) {
	before := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: test.PlayerHitPoints,
		PlayerArmor:     test.PlayerArmor,
		BossHitPoints:   10,
		BossDamage:      test.BossDamage,
	}
	after := before.Handle(BossAttack)
	this.So(after.PlayerHitPoints, should.Equal,test.ExpectedHitPoints)
}

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
		BossAttack,
	}
	for _, turn := range moves {
		game = game.Handle(turn)
	}

	this.So(game, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 2,
		PlayerMana:      24,
		PoisonCounter:   4,
		BossDamage:      8,
		BossHitPoints:   0,
	})
}
