package day17

import (
	"container/ring"
	"fmt"
	"time"

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

func Part2() interface{} {
	now := time.Now()
	length := 50000000
	circle := make([]int, 1, length)
	steps := util.InputInt()
	current := 0
	for x := 1; x < length+1; x++ {
		if x%100000 == 0 {
			fmt.Println(x, time.Since(now))
			now = time.Now()
		}
		insert := advance(current, len(circle), steps) + 1
		current = insert

		circle = append(circle, 0)
		copy(circle[insert+1:], circle[insert:])
		circle[insert] = x

	}
	fmt.Println(circle)
	return circle[current+1]
}

func advance(current, length, steps int) int {
	return (current + steps) % length
}
