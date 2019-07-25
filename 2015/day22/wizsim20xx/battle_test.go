package wizsim20xx

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"github.com/smartystreets/logging"
)

func TestBattleFixture(t *testing.T) {
	gunit.Run(new(BattleFixture), t)
}

type BattleFixture struct {
	*gunit.Fixture
	battle *Battle
}

func (this *BattleFixture) Setup() {
	this.battle = NewBattle()
	this.battle.log = logging.Capture()
	this.battle.log.SetOutput(this)
}

func (this *BattleFixture) TestOneHitMagicMissileBattle_PlayerWins() {
	this.battle.BossHits = 4
	this.battle.BossDamage = 0
	this.battle.PlayerHits = 10
	this.battle.PlayerMana = MagicMissileManaCost
	this.battle.WizardSpell = this.magicMissileForever

	this.So(this.battle.Simulate(), should.Resemble, Result{
		PlayerMana: 0,
		PlayerHits: 10,
		BossHits:   0,
		Turns:      1,
	})
}

func (this *BattleFixture) TestFewHitsMagicMissileBattle_PlayerWins() {
	this.battle.BossHits = 8
	this.battle.BossDamage = 0
	this.battle.PlayerHits = 20
	this.battle.PlayerMana = MagicMissileManaCost * 2
	this.battle.WizardSpell = this.magicMissileForever

	this.So(this.battle.Simulate(), should.Resemble, Result{
		PlayerMana: 0,
		PlayerHits: 19,
		BossHits:   0,
		Turns:      3,
	})
}

func (this *BattleFixture) TestBossWins() {
	this.battle.BossHits = 100
	this.battle.BossDamage = 10
	this.battle.PlayerHits = 20
	this.battle.PlayerMana = 1000
	this.battle.WizardSpell = this.magicMissileForever

	this.So(this.battle.Simulate(), should.Resemble, Result{
		PlayerHits: 0,
		PlayerMana: 1000 - MagicMissileManaCost - MagicMissileManaCost,
		BossHits:   100 - MagicMissileDamage - MagicMissileDamage,
		Turns:      4,
	})
}

func (this *BattleFixture) TestExample1() {
	this.battle.PlayerHits = 10
	this.battle.PlayerMana = 250
	this.battle.BossHits = 13
	this.battle.BossDamage = 8
	this.battle.WizardSpell = this.example1Spells

	this.So(this.battle.Simulate(), should.Resemble, Result{
		PlayerHits: 2,
		PlayerMana: 24,
		BossHits:   0,
		Turns:      4,
	})
}

func (this *BattleFixture) FocusTestExample2() {
	this.battle.PlayerHits = 10
	this.battle.PlayerMana = 250
	this.battle.BossHits = 14
	this.battle.BossDamage = 8
	this.battle.WizardSpell = this.example2Spells

	this.So(this.battle.Simulate(), should.Resemble, Result{
		PlayerHits: 1,
		PlayerMana: 114,
		BossHits:   -1,
		Turns:      10,
	})
}

func (this *BattleFixture) magicMissileForever() spell {
	return MagicMissileSpell
}
func (this *BattleFixture) example1Spells() spell {
	return []spell{
		PoisonSpell,
		MagicMissileSpell,
	}[this.battle.Turns/2]
}
func (this *BattleFixture) example2Spells() spell {
	return []spell{
		RechargeSpell,
		ShieldSpell,
		DrainSpell,
		PoisonSpell,
		MagicMissileSpell,
	}[this.battle.Turns/2]
}
