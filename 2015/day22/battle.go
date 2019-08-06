package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	BossAttack = iota
	Missile
	Drain
	Shield
	Poison
	Recharge
)

var Cost = map[int]int{
	BossAttack: 0,
	Missile:    53,
	Drain:      73,
	Shield:     113,
	Poison:     173,
	Recharge:   229,
}

type Battle struct {
	IsPlayerTurn bool

	PlayerHitPoints int
	PlayerArmor     int
	PlayerMana      int
	PlayerManaSpent int

	BossHitPoints int
	BossDamage    int

	ShieldCounter   int
	PoisonCounter   int
	RechargeCounter int
}

func (this Battle) TakeTurn(attack int) (Battle, string) {
	log := new(bytes.Buffer)

	if this.IsPlayerTurn {
		log.WriteString("-- Player turn --\n")
	} else {
		log.WriteString("-- Boss turn --\n")
	}
	fmt.Fprintf(log, "- Player has %d hit points, %d armor, %d mana\n", this.PlayerHitPoints, this.PlayerArmor, this.PlayerMana)
	fmt.Fprintf(log, "- Boss has %d hit points\n", this.BossHitPoints)

	this.IsPlayerTurn = !this.IsPlayerTurn

	if this.ShieldCounter > 0 {
		this.ShieldCounter--
		fmt.Fprintf(log, "Shield's timer is now %d.\n", this.ShieldCounter)
		if this.ShieldCounter == 0 {
			fmt.Fprintf(log, "Shield wears off, decreasing armor by 7.\n")
			this.PlayerArmor = 0
		}
	}

	if this.PoisonCounter > 0 {
		this.PoisonCounter--
		this.BossHitPoints -= 3
		fmt.Fprintf(log, "Poison deals 3 damage")
		if this.BossHitPoints > 0 {
			fmt.Fprintf(log, "; its timer is now %d.\n", this.PoisonCounter)
		} else {
			fmt.Fprintln(log, ". This kills the boss, and the player wins.")
			return this, strings.TrimSpace(log.String())
		}
	}

	if this.RechargeCounter > 0 {
		this.RechargeCounter--
		this.PlayerMana += 101
		fmt.Fprintf(log, "Recharge provides 101 mana; its timer is now %d.\n", this.RechargeCounter)
		if this.RechargeCounter == 0 {
			fmt.Fprintln(log, "Recharge wears off.")
		}
	}

	this.PlayerMana -= Cost[attack]
	this.PlayerManaSpent += Cost[attack]

	if this.PlayerMana < 1 {
		fmt.Fprintf(log, "Player mana points depleted, and the boss wins.")
		return this, strings.TrimSpace(log.String())
	}

	switch attack {

	case BossAttack:
		damage := this.BossDamage
		armor := this.PlayerArmor
		adjustedDamage := damage - armor
		if adjustedDamage < 1 {
			adjustedDamage = 1
		}
		if this.PlayerArmor > 0 {
			fmt.Fprintf(log, "Boss attacks for %d - %d = %d damage!\n", this.BossDamage, this.PlayerArmor, adjustedDamage)
		} else {
			fmt.Fprintf(log, "Boss attacks for %d damage!\n", this.BossDamage)
		}
		this.PlayerHitPoints -= adjustedDamage

	case Missile:
		log.WriteString("Player casts Magic Missile, dealing 4 damage.")
		this.BossHitPoints -= 4

	case Drain:
		log.WriteString("Player casts Drain, dealing 2 damage, and healing 2 hit points.")
		this.BossHitPoints -= 2
		this.PlayerHitPoints += 2

	case Poison:
		if this.PoisonCounter > 0 {
			log.WriteString("Player casts Poison illegally, and the boss wins.")
		} else {
			log.WriteString("Player casts Poison.\n")
			this.PoisonCounter = 6
		}

	case Shield:
		if this.ShieldCounter > 0 {
			log.WriteString("Player casts Shield illegally, and the boss wins.")
		} else {
			log.WriteString("Player casts Shield, increasing armor by 7.")
			this.ShieldCounter += 6
			this.PlayerArmor = 7
		}

	case Recharge:
		if this.RechargeCounter > 0 {
			log.WriteString("Player casts Recharge illegally, and the boss wins.")
		} else {
			log.WriteString("Player casts Recharge.\n")
			this.RechargeCounter = 5
		}
	}

	if this.BossHitPoints < 1 {
		log.WriteString(" This kills the boss, and the player wins.")
	} else if this.PlayerHitPoints < 1 {
		log.WriteString(" This kills the player, and the boss wins.")
	} else if this.IsPlayerTurn && this.PlayerMana < 53 && this.RechargeCounter == 0 {
		fmt.Fprintf(log, " The player has only %d mana, and the boss wins.", this.PlayerMana)
	}

	return this, strings.TrimSpace(log.String())
}
