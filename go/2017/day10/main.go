package day10

import (
	"github.com/mdwhatcott/advent-of-code/go/2017/knot"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Part1() int {
	circle := knot.NewLoop()
	circle.TwistAll(toBytes(util.InputInts(",")))
	a, b := circle.FirstTwo()
	return int(a) * int(b)
}

func Part2() string {
	return knot.HashString(util.InputString())
}

func toBytes(input []int) (bytes []byte) {
	for _, l := range input {
		bytes = append(bytes, byte(l))
	}
	return bytes
}
