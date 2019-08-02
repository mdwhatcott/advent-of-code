package main

type Battle struct {
	IsPlayerTurn bool

	PlayerHitPoints int
	PlayerArmor     int
	PlayerMana      int

	BossHitPoints int
	BossDamage    int

	ShieldCounter   int
	PoisonCounter   int
	RechargeCounter int
}

func (this Battle) NextMoves() (moves []Turn) {
	if this.gameOver() {
		return nil
	} else if this.IsPlayerTurn {
		return this.collectPlayerMoves()
	} else {
		return append(moves, new(BossAttack))
	}
}

func (this Battle) collectPlayerMoves() (moves []Turn) {
	if this.PlayerMana >= 53 {
		moves = append(moves, new(Missile))
	}
	if this.PlayerMana >= 73 {
		moves = append(moves, new(Drain))
	}
	if this.PlayerMana >= 113 {
		moves = append(moves, new(Shield))
	}
	if this.PlayerMana >= 173 {
		moves = append(moves, new(Poison))
	}
	if this.PlayerMana >= 229 {
		moves = append(moves, new(Recharge))
	}
	return moves
}

func (this Battle) gameOver() bool {
	return (this.IsPlayerTurn && this.playerDead()) || this.bossDead()
}
func (this Battle) bossDead() bool   { return this.BossHitPoints < 1 }
func (this Battle) playerDead() bool { return this.PlayerHitPoints < 1 }
