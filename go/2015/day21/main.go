package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Part 1: Cheapest winning combo:", part1())      // 91
	fmt.Println("Part 2: Most expensive losing combo:", part2()) // 158
}

func part2() int {
	store := loadStore()
	combinations := store.LoadPurchaseCombinations()
	sort.Slice(combinations, func(i, j int) bool {
		return combinations[i].Cost() > combinations[j].Cost()
	})
	for _, combination := range combinations {
		a := NewCharacter("player", 100, combination.Damage(), combination.Defense())
		b := NewCharacter("boss", 100, 8, 2)

		winner := Fight(a, b)
		if winner == b {
			return combination.Cost()
		}
	}
	panic("No answer found")
}

func part1() int {
	store := loadStore()
	combinations := store.LoadPurchaseCombinations()
	sort.Slice(combinations, func(i, j int) bool {
		return combinations[i].Cost() < combinations[j].Cost()
	})
	for _, combination := range combinations {
		a := NewCharacter("player", 100, combination.Damage(), combination.Defense())
		b := NewCharacter("boss", 100, 8, 2)

		winner := Fight(a, b)
		if winner == a {
			return combination.Cost()
		}
	}
	panic("No answer found")
}
