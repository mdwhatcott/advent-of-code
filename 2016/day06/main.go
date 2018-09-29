package main

import (
	"fmt"

	"advent/lib/util"
)

const wordLength = 8

func main() {
	code := NewRepetitionCode(util.InputScanner(), wordLength)

	fmt.Println("Answer based on frequent character:", code.DecodeFrequent())
	fmt.Println("Answer based on infrequent character:", code.DecodeInfrequent())
}
