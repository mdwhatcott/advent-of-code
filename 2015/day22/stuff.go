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
	if this.IsPlayerTurn {
		moves = append(moves, new(Missile))
		moves = append(moves, new(Drain))
		moves = append(moves, new(Shield))
		moves = append(moves, new(Poison))
		moves = append(moves, new(Recharge))
	} else {
		moves = append(moves, new(BossAttack))
	}
	return moves
}
