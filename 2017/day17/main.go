package day17

import (
	"container/ring"

	"advent/lib/util"
)

func Part1() interface{} {
	circle := ring.New(1)
	circle.Value = 0
	for x := 1; x < 2018; x++ {
		for y := 0; y < util.InputInt(); y++ {
			circle = circle.Next()
		}
		insert := ring.New(1)
		insert.Value = x
		circle = circle.Link(insert)
		circle = insert
	}
	return circle.Next().Value
}

func Part2() int {
	return 0
}
