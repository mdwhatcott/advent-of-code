package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	generator := NewGenerator(util.InputString())
	fmt.Println("64th one-time pad:", generator.IndexOfKey(64))

	generator = NewGenerator(util.InputString())
	generator.stretch = true
	fmt.Println("64th stretched one-time pad:", generator.IndexOfKey(64))
}
