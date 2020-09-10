// This solution originally derived from:
// https://github.com/fogleman/AdventOfCode2019/blob/master/14.py
package advent

import (
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	return NewReactor(Parse(util.InputScanner())).ResolveOreCost()
}

func Part2() interface{} {
	return nil
}

func Parse(scanner *util.Scanner) map[string]Reaction {
	result := make(map[string]Reaction)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "=> ", "")
		fields := strings.Fields(line)
		key := fields[len(fields)-1]
		recipe := Reaction{OutputQuantity: util.ParseInt(fields[len(fields)-2])}
		fields = fields[:len(fields)-2]
		for x := 0; x < len(fields); x += 2 {
			recipe.Reactants = append(recipe.Reactants, Reactant{
				Count:    util.ParseInt(fields[x]),
				Material: fields[x+1],
			})
		}
		result[key] = recipe
	}
	return result
}

type Reaction struct {
	OutputQuantity int
	Reactants      []Reactant
}

type Reactant struct {
	Count    int
	Material string
}

type Reactor struct {
	reactions   map[string]Reaction
	batchCounts map[string]int
	reactants   map[string]int
	products    map[string]int
}

func NewReactor(reactions map[string]Reaction) *Reactor {
	return &Reactor{
		reactions:   reactions,
		batchCounts: map[string]int{"FUEL": 1},
	}
}

func (this *Reactor) ResolveOreCost() (ore int) {
	for !this.oreCostAlreadyResolved() {
		this.takeInventoryOfResolvedBatches()
		this.updateBatchCounts()
	}

	return this.ore()
}

func (this *Reactor) oreCostAlreadyResolved() bool {
	for name, neededQuantity := range this.reactants {
		if name != "ORE" && this.products[name] < neededQuantity {
			return false
		}
	}
	return this.ore() > 0
}

func (this *Reactor) updateBatchCounts() {
	for name, neededQuantity := range this.reactants {
		if name == "ORE" {
			continue
		}
		alreadyProduced := this.products[name]
		if alreadyProduced >= neededQuantity {
			continue
		}

		target := this.reactions[name].OutputQuantity
		lacking := neededQuantity - alreadyProduced
		this.batchCounts[name] += this.calculateBatchesNeeded(target, lacking)
	}
}

func (this *Reactor) calculateBatchesNeeded(outputQuantity, lacking int) int {
	neededBatches := lacking / outputQuantity
	remainder := lacking % outputQuantity
	if remainder > 0 {
		neededBatches++
	}
	return neededBatches
}

func (this *Reactor) takeInventoryOfResolvedBatches() {
	this.reactants = make(map[string]int)
	this.products = make(map[string]int)

	for name, multiplier := range this.batchCounts {
		recipe := this.reactions[name]

		this.products[name] = recipe.OutputQuantity * multiplier

		for _, ingredient := range recipe.Reactants {
			this.reactants[ingredient.Material] += ingredient.Count * multiplier
		}
	}
}

func (this *Reactor) ore() int {
	return this.reactants["ORE"]
}
