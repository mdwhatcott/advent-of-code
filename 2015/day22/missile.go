package main

type Missile struct{}

func (this Missile) Perform(state Battle) Battle {
	state.BossHitPoints -= 4
	return state
}
