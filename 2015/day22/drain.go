package main

type Drain struct{}

func (this Drain) Perform(state Battle) Battle {
	return state
}
