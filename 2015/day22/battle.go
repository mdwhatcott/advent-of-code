package main

const (
	BossAttack = iota
	Missile
	Drain
	Shield
	Poison
	Recharge
)

var Cost = map[int]int{
	BossAttack: 0,
	Missile:    53,
	Drain:      73,
	Shield:     113,
	Poison:     173,
	Recharge:   229,
}

type Battle struct {
	IsPlayerTurn bool

	PlayerHitPoints int
	PlayerArmor     int
	PlayerMana      int
	PlayerManaSpent int

	BossHitPoints int
	BossDamage    int

	ShieldCounter   int
	PoisonCounter   int
	RechargeCounter int
}
