package day17

import (
	"container/ring"

	"advent/lib/util"
)

func Part1() interface{} {
	steps := util.InputInt()
	circle := ring.New(1)
	circle.Value = 0
	for x := 1; x < 2018; x++ {
		for y := 0; y < steps; y++ {
			circle = circle.Next()
		}
		insert := ring.New(1)
		insert.Value = x
		circle = circle.Link(insert)
		circle = insert
	}
	return circle.Next().Value
}

// https://www.reddit.com/r/adventofcode/comments/7kc0xw/2017_day_17_solutions/drd5yek/
func Part2() (out int) {
	steps := util.InputInt()

	for current, i := 0, 1; i <= 50000000+1; i++ {
		next := (current+steps)%i + 1
		if next == 1 {
			out = i
		}
		current = next
	}
	return out
}
