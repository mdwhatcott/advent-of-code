package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestMissileFixture(t *testing.T) {
    gunit.Run(new(MissileFixture), t)
}

type MissileFixture struct {
    *gunit.Fixture
}

func (this *MissileFixture) TestMissileDoes4DamageToBoss() {
	before := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3,
		BossHitPoints:   8,
		BossDamage:      5,
		ShieldCounter:   6,
		PoisonCounter:   7,
		RechargeCounter: 8,
	}
	after := new(Missile).Perform(before)

	this.So(&before, should.NotPointTo, &after)
	this.So(after, should.Resemble, Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3,
		BossHitPoints:   4, // was 8
		BossDamage:      5,
		ShieldCounter:   6,
		PoisonCounter:   7,
		RechargeCounter: 8,
	})
}
