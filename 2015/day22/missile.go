package main

type Missile struct{}

func (this *Missile) Perform(state Battle) Battle {
	return state
}
