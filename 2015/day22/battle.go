package main

const (
	BossAttack = iota
	Missile
	Drain
	Poison
	Shield
	Recharge
)

type Battle struct {
	IsPlayerTurn bool

	PlayerHitPoints int
	PlayerArmor     int
	PlayerMana      int
	//PlayerManaSpent int // TODO

	BossHitPoints int
	BossDamage    int

	ShieldCounter   int
	PoisonCounter   int
	RechargeCounter int
}

func (this Battle) Attack() (moves []int) {
	if this.gameOver() {
		return nil
	} else if this.IsPlayerTurn {
		return this.collectPlayerMoves()
	} else {
		return append(moves, BossAttack)
	}
}

func (this Battle) collectPlayerMoves() (moves []int) {
	if this.canCastMissile() {
		moves = append(moves, Missile)
	}
	if this.canCastDrain() {
		moves = append(moves, Drain)
	}
	if this.canCastShield() {
		moves = append(moves, Shield)
	}
	if this.canCastPoison() {
		moves = append(moves, Poison)
	}
	if this.canCastRecharge() {
		moves = append(moves, Recharge)
	}
	return moves
}

func (this Battle) canCastMissile() bool {
	return this.PlayerMana >= 53
}
func (this Battle) canCastDrain() bool {
	return this.PlayerMana >= 73
}
func (this Battle) canCastShield() bool {
	return this.PlayerMana >= 113 && this.ShieldCounter < 2 // TODO: off by one?
}
func (this Battle) canCastPoison() bool {
	return this.PlayerMana >= 173 && this.PoisonCounter < 2 // TODO: off by one?
}
func (this Battle) canCastRecharge() bool {
	return this.PlayerMana >= 229 && this.RechargeCounter < 2 // TODO: off by one?
}

func (this Battle) gameOver() bool {
	return this.PlayerHitPoints < 1 || this.BossHitPoints < 1
}

func (this Battle) Handle(attack int) Battle {
	this.IsPlayerTurn = !this.IsPlayerTurn

	if this.PoisonCounter > 0 {
		this.PoisonCounter--
		this.BossHitPoints -= 3
	}

	switch attack {
	case BossAttack:
		this.PlayerHitPoints -= this.BossDamage // TODO: There is more to boss attacks than this (armor)
	case Missile:
		this.PlayerMana -= 53
		this.BossHitPoints -= 4
	case Drain:
		this.PlayerMana -= 73
		this.BossHitPoints -= 2
		this.PlayerHitPoints += 2
	case Poison:
		this.PlayerMana -= 173
		this.BossHitPoints -= 3
		this.PoisonCounter = 5
	}
	return this
}
