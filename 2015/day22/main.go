package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	log.SetFlags(0)
	fmt.Println("Least mana spent (expected: 1824) was...", part1())
}

func part1() (min int) {
	min = 999999999

	battle := Battle{
		IsPlayerTurn:    true,
		PlayerHitPoints: 50,
		PlayerMana:      500,
		BossHitPoints:   71,
		BossDamage:      10,
	}

	battles := 0
	wins := 0
	queue := []Battle{battle}

	for len(queue) > 0 {
		battle := queue[0]
		queue = queue[1:]

		var moves = bossAttack
		if battle.IsPlayerTurn {
			moves = allSpells
		}
		for _, move := range moves {
			next, report := battle.TakeTurn(move)

			if strings.Contains(report, "player wins") {
				wins++
				log.Println("Player!", strings.ReplaceAll(report, "\n", " | "))
				if next.PlayerManaSpent < min {
					log.Println("MIN!", next.PlayerManaSpent)
					min = next.PlayerManaSpent
				}
				battles++
			} else if strings.Contains(report, "boss wins") {
				//log.Println("Boss!", strings.ReplaceAll(report, "\n", " | "))
				battles++
			} else if next.PlayerManaSpent < min {
				queue = append(queue, next)
			}
		}
	}
	fmt.Println("Battles:", battles)
	fmt.Println("Player wins:", wins)

	return min
}

var bossAttack = []int{BossAttack}
var allSpells = []int{Missile, Drain, Shield, Poison, Recharge}
