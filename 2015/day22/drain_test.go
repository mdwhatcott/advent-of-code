package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestDrainFixture(t *testing.T) {
	gunit.Run(new(DrainFixture), t)
}

type DrainFixture struct {
	*gunit.Fixture
}

func (this *DrainFixture) TestDrainDeals2DamageToBossAndHealsPlayerBy2() {
	before := Battle{
		IsPlayerTurn:    false,
		PlayerHitPoints: 1,
		PlayerArmor:     2,
		PlayerMana:      3,
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
