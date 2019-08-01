package main

type Poison struct{}

func (this *Poison) Perform(previous Battle) (result Battle) {
	panic("implement me")
}
