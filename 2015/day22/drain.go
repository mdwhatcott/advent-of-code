package main

type Drain struct{}

func (this Drain) Perform(state Battle) Battle {
	state.BossHitPoints -= 2
	state.PlayerHitPoints += 2
	return state
}
