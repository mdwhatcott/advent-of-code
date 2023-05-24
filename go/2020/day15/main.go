package advent

import (
	"container/list"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	return part1(2020, util.InputInts(",")...)
}

func part1(until int, inputs ...int) interface{} {
	game := list.New()

	for _, input := range inputs {
		game.PushFront(input)
	}

	for game.Len() < until {
		next := sinceLastOccurrence(game.Front())
		game.PushFront(next)
	}

	return game.Front().Value
}

func sinceLastOccurrence(cursor *list.Element) int {
	search := cursor.Value
	for x := 1; ; x++ {
		cursor = cursor.Next()
		if cursor == nil {
			return 0
		}
		if cursor.Value == search {
			return x
		}
	}
}

func Part2() interface{} {
	return part2(30000000, util.InputInts(",")...)
}

func part2(until int, inputs ...int) interface{} {
	game := list.New()
	memory := make(map[interface{}][]int) // map[num][]occurrences

	for _, input := range inputs {
		game.PushFront(input)
		memory[input] = append(memory[input], game.Len())
	}

	for game.Len() < until {
		next := sinceLastOccurrenceCache(memory, game.Front())
		game.PushFront(next)
		memory[next] = append(memory[next], game.Len())
	}

	return game.Front().Value
}

func sinceLastOccurrenceCache(memory map[interface{}][]int, front *list.Element) int {
	values := memory[front.Value]
	if len(values) == 1 {
		return 0
	}

	if len(values) > 2 {
		values = values[len(values)-2:]
	}

	return values[len(values)-1] - values[len(values)-2]
}
