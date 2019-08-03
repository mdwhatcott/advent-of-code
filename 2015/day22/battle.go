package main

type Battle struct {
	IsPlayerTurn bool

	PlayerHitPoints int
	PlayerArmor     int
	PlayerMana      int
	//PlayerManaSpent int

	BossHitPoints int
	BossDamage    int

	ShieldCounter   int
	PoisonCounter   int
	RechargeCounter int
}

func (this Battle) NextMoves() (moves []interface{}) {
	if this.gameOver() {
		return nil
	} else if this.IsPlayerTurn {
		return this.collectPlayerMoves()
	} else {
		return append(moves, new(BossAttack))
	}
}

func (this Battle) collectPlayerMoves() (moves []interface{}) {
	if this.canCastMissile() {
		moves = append(moves, new(Missile))
	}
	if this.canCastDrain() {
		moves = append(moves, new(Drain))
	}
	if this.canCastShield() {
		moves = append(moves, new(Shield))
	}
	if this.canCastPoison() {
		moves = append(moves, new(Poison))
	}
	if this.canCastRecharge() {
		moves = append(moves, new(Recharge))
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
	return this.PlayerMana >= 113 && this.ShieldCounter < 2
}
func (this Battle) canCastPoison() bool {
	return this.PlayerMana >= 173 && this.PoisonCounter < 2
}
func (this Battle) canCastRecharge() bool {
	return this.PlayerMana >= 229 && this.RechargeCounter < 2
}

func (this Battle) gameOver() bool {
	return (this.IsPlayerTurn && this.playerDead()) || this.bossDead()
}
func (this Battle) bossDead() bool   { return this.BossHitPoints < 1 }
func (this Battle) playerDead() bool { return this.PlayerHitPoints < 1 }

func (this Battle) Handle(e interface{}) Battle {
	switch e.(type) {
	case *Drain:
		this.BossHitPoints -= 2
		this.PlayerHitPoints += 2
	case *Missile:
		this.BossHitPoints -= 4
	}
	return this
}
