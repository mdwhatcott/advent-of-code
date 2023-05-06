package day01

import (
	"advent/lib/parse"
)

func sumOfIdenticalNeighborDigits(input []string) (sum int) {
	input = append(input, input[0]) // wrap-around

	for x := 0; x < len(input)-1; x++ {
		c1 := input[x]
		c2 := input[x+1]

		if c1 == c2 {
			sum += parse.Int(c2)
		}
	}

	return sum
}

func sumOfIdenticalOppositeDigits(input []string) (sum int) {
	length := len(input)
	half := length / 2

	for x := 0; x < len(input); x++ {
		c1 := input[x]
		x2 := x + half
		if x2 >= length {
			x2 -= length // wrap-around
		}
		c2 := input[x2]

		if c1 == c2 {
			sum += parse.Int(c2)
		}
	}

	return sum
}
