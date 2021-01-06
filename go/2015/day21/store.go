package main

import (
	"strings"

	"advent/lib/util"
)

type Equipment struct {
	Name    string
	Cost    int
	Defense int
	Damage  int
}

func (this Equipment) IsRing() bool {
	return strings.Contains(this.Name, "+")
}

type Store struct {
	weapons []Equipment
	armor   []Equipment
	rings   []Equipment
}

func newStore() *Store {
	store := &Store{}
	return store
}

type Purchase []Equipment

func (this Purchase) Cost() (cost int) {
	for _, item := range this {
		cost += item.Cost
	}
	return cost
}

func (this Purchase) Defense() (total int) {
	for _, item := range this {
		total += item.Defense
	}
	return total
}

func (this Purchase) Damage() (total int) {
	for _, item := range this {
		total += item.Damage
	}
	return total
}

func (this *Store) LoadPurchaseCombinations() (all []Purchase) {
	for _, weapon := range this.weapons {
		all = append(all, Purchase{weapon})

		for _, armor := range this.armor {
			all = append(all, Purchase{weapon, armor})

			for _, ring1 := range this.rings {
				all = append(all, Purchase{weapon, armor, ring1})
				all = append(all, Purchase{weapon, ring1})

				for _, ring2 := range this.rings {
					if ring1 != ring2 {
						all = append(all, Purchase{weapon, armor, ring1, ring2})
						all = append(all, Purchase{weapon, ring1, ring2})
					}
				}
			}
		}
	}
	return all
}

func (this *Store) Stock(equipment Equipment) {
	if equipment.IsRing() {
		this.rings = append(this.rings, equipment)
	} else if equipment.Damage > 0 {
		this.weapons = append(this.weapons, equipment)
	} else if equipment.Defense > 0 {
		this.armor = append(this.armor, equipment)
	}
}

func loadStore() *Store {
	store := newStore()
	scanner := util.NewScanner(strings.NewReader(strings.TrimSpace(rawStore)))
	for scanner.Scan() {
		fields := scanner.Fields()
		store.Stock(Equipment{
			Name:    fields[0],
			Cost:    util.ParseInt(fields[1]),
			Damage:  util.ParseInt(fields[2]),
			Defense: util.ParseInt(fields[3]),
		})
	}
	return store
}

// Store:   Cost  Damage  Armor
const rawStore = `
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5
Damage+1     25     1       0
Damage+2     50     2       0
Damage+3    100     3       0
Defense+1    20     0       1
Defense+2    40     0       2
Defense+3    80     0       3
`
