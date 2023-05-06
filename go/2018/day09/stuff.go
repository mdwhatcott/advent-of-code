package day09

import (
	"container/ring"
	"fmt"
	"strings"

	"advent/lib/maths"
	"advent/lib/parse"
)

func Parse(input string) (players int, maxMarble int) {
	fields := strings.Fields(input)
	players = parse.Int(fields[0])
	maxMarble = parse.Int(fields[6])
	return players, maxMarble
}

func Parse10X(input string) (players int, maxMarble int) {
	players, maxMarble = Parse(input)
	return players, maxMarble * 100
}

func MarbleHighScore(playerCount int, maxMarble int) int {
	var players = make([]int, playerCount)
	start := ring.New(1)
	start.Value = 0
	circle := start
	for marble := 1; marble <= maxMarble; marble++ {
		if marble%23 == 0 {
			players[marble%playerCount] += marble
			for x := 0; x < 9; x++ {
				circle = circle.Prev()
			}
			removed := circle.Unlink(1)
			players[marble%playerCount] += removed.Value.(int)
			circle = circle.Next().Next()
		} else {
			insert := ring.New(1)
			insert.Value = marble
			circle = circle.Link(insert)
		}
		//debug(start, marble)
	}
	return maths.Max(players...)
}

func debug(circle *ring.Ring, marble int) {
	circle.Do(func(f interface{}) {
		if f == marble {
			fmt.Printf("(%v) ", f)
		} else {
			fmt.Print(f, " ")
		}
	})
	fmt.Println()
}
