package main

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBattleFixture(t *testing.T) {
	gunit.Run(new(BattleFixture), t)
}

type BattleFixture struct {
	*gunit.Fixture
}

func (this *BattleFixture) Test_Example1() {
	a := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 10,
		PlayerArmor:     0,
		PlayerMana:      250,
		PlayerManaSpent: 0,
		BossHitPoints:   13,
		BossDamage:      8,
		ShieldCounter:   0,
		PoisonCounter:   0,
		RechargeCounter: 0,
	}

	b, B := a.TakeTurn(Poison)

	this.So(b.IsPlayerTurn, should.BeFalse)
	this.So(b.PlayerMana, should.Equal, 250-Cost[Poison])
	this.So(b.PlayerManaSpent, should.Equal, Cost[Poison])
	this.So(b.PoisonCounter, should.Equal, 6)
	this.So(B, should.Equal, strings.TrimSpace(`
-- Player turn --
- Player has 10 hit points, 0 armor, 250 mana
- Boss has 13 hit points
Player casts Poison.
`))

	c, C := b.TakeTurn(BossAttack)

	this.So(c.IsPlayerTurn, should.BeTrue)
	this.So(c.PoisonCounter, should.Equal, 5)
	this.So(c.BossHitPoints, should.Equal, 10)
	this.So(c.PlayerHitPoints, should.Equal, 2)
	this.So(C, should.Equal, strings.TrimSpace(`
-- Boss turn --
- Player has 10 hit points, 0 armor, 77 mana
- Boss has 13 hit points
Poison deals 3 damage; its timer is now 5.
Boss attacks for 8 damage!
`))

	d, D := c.TakeTurn(Missile)

	this.So(d.IsPlayerTurn, should.BeFalse)
	this.So(d.PoisonCounter, should.Equal, 4)
	this.So(d.BossHitPoints, should.Equal, 3)
	this.So(d.PlayerMana, should.Equal, 250-Cost[Poison]-Cost[Missile])
	this.So(d.PlayerManaSpent, should.Equal, Cost[Poison]+Cost[Missile])
	this.So(D, should.Equal, strings.TrimSpace(`
-- Player turn --
- Player has 2 hit points, 0 armor, 77 mana
- Boss has 10 hit points
Poison deals 3 damage; its timer is now 4.
Player casts Magic Missile, dealing 4 damage.
`))

	e, E := d.TakeTurn(BossAttack)

	this.So(e.IsPlayerTurn, should.BeTrue)
	this.So(e.PoisonCounter, should.Equal, 3)
	this.So(e.BossHitPoints, should.Equal, 0)
	this.So(e.PlayerHitPoints, should.Equal, 2)
	this.So(e.PlayerMana, should.Equal, 250-Cost[Poison]-Cost[Missile])
	this.So(e.PlayerManaSpent, should.Equal, Cost[Poison]+Cost[Missile])
	this.So(E, should.Equal, strings.TrimSpace(`
-- Boss turn --
- Player has 2 hit points, 0 armor, 24 mana
- Boss has 3 hit points
Poison deals 3 damage. This kills the boss, and the player wins.
`))
}

func (this *BattleFixture) Test_Example2() {
	battle := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 10,
		PlayerArmor:     0,
		PlayerMana:      250,
		PlayerManaSpent: 0,
		BossHitPoints:   14,
		BossDamage:      8,
		ShieldCounter:   0,
		PoisonCounter:   0,
		RechargeCounter: 0,
	}

	for t, turn := range turns {
		this.Println("Turn:", t)
		var log string
		battle, log = battle.TakeTurn(turn.Attack)
		log = strings.ReplaceAll(log, "\n", " | ")
		turn.ExpectedLog = strings.ReplaceAll(turn.ExpectedLog, "\n", " | ")
		if !this.So(log, should.Equal, turn.ExpectedLog) {
			break
		}
	}
}

var turns = []Example2Turn{
	{
		Attack: Recharge, // 0
		ExpectedLog: "" +
			"-- Player turn --" + "\n" +
			"- Player has 10 hit points, 0 armor, 250 mana" + "\n" +
			"- Boss has 14 hit points" + "\n" +
			"Player casts Recharge.",
	},
	{
		Attack: BossAttack, // 1
		ExpectedLog: "" +
			"-- Boss turn --" + "\n" +
			"- Player has 10 hit points, 0 armor, 21 mana" + "\n" +
			"- Boss has 14 hit points" + "\n" +
			"Recharge provides 101 mana; its timer is now 4." + "\n" +
			"Boss attacks for 8 damage!",
	},
	{
		Attack: Shield, // 2
		ExpectedLog: "" +
			"-- Player turn --" + "\n" +
			"- Player has 2 hit points, 0 armor, 122 mana" + "\n" +
			"- Boss has 14 hit points" + "\n" +
			"Recharge provides 101 mana; its timer is now 3." + "\n" +
			"Player casts Shield, increasing armor by 7.",
	},
	{
		Attack: BossAttack, // 3
		ExpectedLog: "" +
			"-- Boss turn --" + "\n" +
			"- Player has 2 hit points, 7 armor, 110 mana" + "\n" +
			"- Boss has 14 hit points" + "\n" +
			"Shield's timer is now 5." + "\n" +
			"Recharge provides 101 mana; its timer is now 2." + "\n" +
			"Boss attacks for 8 - 7 = 1 damage!",
	},
	{
		Attack: Drain, // 4
		ExpectedLog: "" +
			"-- Player turn --" + "\n" +
			"- Player has 1 hit points, 7 armor, 211 mana" + "\n" +
			"- Boss has 14 hit points" + "\n" +
			"Shield's timer is now 4." + "\n" +
			"Recharge provides 101 mana; its timer is now 1." + "\n" +
			"Player casts Drain, dealing 2 damage, and healing 2 hit points.",
	},
	{
		Attack: BossAttack, // 5
		ExpectedLog: "" +
			"-- Boss turn --" + "\n" +
			"- Player has 3 hit points, 7 armor, 239 mana" + "\n" +
			"- Boss has 12 hit points" + "\n" +
			"Shield's timer is now 3." + "\n" +
			"Recharge provides 101 mana; its timer is now 0." + "\n" +
			"Recharge wears off." + "\n" +
			"Boss attacks for 8 - 7 = 1 damage!",
	},
	{
		Attack: Poison, // 6
		ExpectedLog: "" +
			"-- Player turn --" + "\n" +
			"- Player has 2 hit points, 7 armor, 340 mana" + "\n" +
			"- Boss has 12 hit points" + "\n" +
			"Shield's timer is now 2." + "\n" +
			"Player casts Poison.",
	},
	{
		Attack: BossAttack, // 7
		ExpectedLog: "" +
			"-- Boss turn --" + "\n" +
			"- Player has 2 hit points, 7 armor, 167 mana" + "\n" +
			"- Boss has 12 hit points" + "\n" +
			"Shield's timer is now 1." + "\n" +
			"Poison deals 3 damage; its timer is now 5." + "\n" +
			"Boss attacks for 8 - 7 = 1 damage!",
	},
	{
		Attack: Missile, // 8
		ExpectedLog: "" +
			"-- Player turn --" + "\n" +
			"- Player has 1 hit points, 7 armor, 167 mana" + "\n" +
			"- Boss has 9 hit points" + "\n" +
			"Shield's timer is now 0." + "\n" +
			"Shield wears off, decreasing armor by 7." + "\n" +
			"Poison deals 3 damage; its timer is now 4." + "\n" +
			"Player casts Magic Missile, dealing 4 damage.",
	},
	{
		Attack: BossAttack, // 9
		ExpectedLog: "" +
			"-- Boss turn --" + "\n" +
			"- Player has 1 hit points, 0 armor, 114 mana" + "\n" +
			"- Boss has 2 hit points" + "\n" +
			"Poison deals 3 damage. This kills the boss, and the player wins.",
	},
}

type Example2Turn struct {
	Attack      int
	ExpectedLog string
}
