package main

type Turn interface {
	Perform(previous Battle) (result Battle)
}
