package main

import (
	"fmt"

	"advent/lib/astar"
	"github.com/ferhatelmas/levenshtein"
)

type MoleculeSearch struct {
	goal         string
	current      string
	replacements []string
}

func (this *MoleculeSearch) EstimatedDistanceToTarget() float64 {
	distance := float64(levenshtein.Dist(this.goal, this.current))
	fmt.Printf("distance from %s: %0.f\n", this.current, distance)
	return distance
}

func (this *MoleculeSearch) AdjacentPositions() (adjacent []astar.Turtle) {
	machine := NewMoleculeMachine()
	for _, replacement := range this.replacements {
		machine.RegisterReplacement(replacement)
	}

	for _, transformation := range machine.Calibrate(this.current) {
		// TODO: filter out transformations we've already seen? (infinite loops?)
		adjacent = append(adjacent, &MoleculeSearch{
			goal:         this.goal,
			current:      transformation,
			replacements: this.replacements,
		})
	}
	return adjacent
}

func (this *MoleculeSearch) Hash() string {
	return this.current
}
