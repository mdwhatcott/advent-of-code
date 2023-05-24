package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func main() {
	input := util.InputString()

	keypad := NewKeypad(nineKeyDirections)
	fmt.Println("9-key Code:", keypad.DeriveCode(input))

	keypad = NewKeypad(thirteenKeyDirections)
	fmt.Println("13-key Code:", keypad.DeriveCode(input))
}
