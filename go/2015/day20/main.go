package main

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	fmt.Println("Part 1 Answer:", part1()) // 665280
	fmt.Println("Part 2 Answer:", part2()) // 705600
}

func part1() int {
	input := util.InputInt()
	max := 0
	startingHouse := 660000 // found by trial and error
	for house := startingHouse; ; house++ {
		count := presents(house)
		if count > max {
			fmt.Println("Progress:", house, count)
			max = count
		}
		if count >= input {
			return house
		}
	}
}

func presents(house int) (presents int) {
	for elf := 1; elf <= house; elf++ {
		if house%elf == 0 {
			presents += elf * 10
		}
	}
	return presents
}

func part2() int {
	input := util.InputInt()
	max := 0
	presents := NewPresents()
	startingHouse := 700000
	for house := startingHouse; ; house++ {
		count := presents.Deliver(house)
		if count > max {
			fmt.Println("Progress:", house, count)
			max = count
		}
		if count >= input {
			return house
		}
	}
}

type Presents struct {
	deliveries map[int]int
	minElf     int
}

func NewPresents() *Presents {
	return &Presents{deliveries: make(map[int]int), minElf: 1}
}

func (this *Presents) Deliver(house int) (presents int) {
	for elf := this.minElf; elf <= house; elf++ {
		delivered := this.deliveries[elf]
		if house%elf == 0 && delivered < 10 {
			presents += elf * 10
			this.deliveries[elf]++
		}
	}
	this.incrementMinElf()
	return presents
}

func (this *Presents) incrementMinElf() {
	for {
		delivered := this.deliveries[this.minElf]
		if delivered == 10 {
			this.minElf++
		} else {
			break
		}
	}
}
