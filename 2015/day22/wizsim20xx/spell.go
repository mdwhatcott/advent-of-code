package wizsim20xx

type conjurer func() spell

type spell int

const (
	MagicMissileSpell spell = iota
	DrainSpell
	ShieldSpell
	PoisonSpell
	RechargeSpell
)

var spellCost = map[spell]int{
	MagicMissileSpell: MagicMissileManaCost,
	DrainSpell:        DrainManaCost,
	ShieldSpell:       ShieldManaCost,
	PoisonSpell:       PoisonManaCost,
	RechargeSpell:     RechargeManaCost,
}

const (
	MagicMissileManaCost = 53
	MagicMissileDamage   = 4

	DrainManaCost = 73
	DrainDamage   = 2
	DrainHealing  = 2

	ShieldManaCost              = 113
	ShieldEffectDurationInTurns = 6
	ShieldArmor                 = 7

	PoisonManaCost              = 173
	PoisonDamage                = 3
	PoisonEffectDurationInTurns = 6

	RechargeManaCost              = 229
	RechargeManaBonus             = 101
	RechargeEffectDurationInTurns = 5
)

