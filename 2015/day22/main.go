package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	fmt.Println("Least mana spent:", part1())
}

func part1() (min int) {
	min = 1000000

	battle := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 50,
		PlayerMana:      250,
		BossHitPoints:   71,
		BossDamage:      10,
	}

	queue := []Battle{battle}

	for len(queue) > 0 {

		before := queue[0]
		if len(queue) == 1 {
			queue = nil
		} else {
			queue = queue[1:]
		}

		moves := before.Attack()
		for _, move := range moves {
			after := before.Handle(move)

			if !after.GameOver() {
				queue = append(queue, after)
			} else if after.PlayerHitPoints > 0 {
				log.Printf("Player won: %+v", after)
				min = after.PlayerManaSpent
			} else {
				log.Printf("Boss won: %+v", after)
			}
		}
	}

	return min
}
