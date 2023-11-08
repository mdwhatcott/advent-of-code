package starter

import "github.com/mdwhatcott/grid"

type Unit struct {
	species  rune
	location grid.Point[int]
	targets  []*Unit
	damage   int
}

func NewUnit(species rune, x, y int) *Unit {
	return &Unit{species: species, location: grid.NewPoint(x, y)}
}
func (this *Unit) HP() int {
	return max(0, 200-this.damage)
}

func AssociateEnemyUnits(all ...*Unit) {
	for c, c1 := range all {
		for _, c2 := range all[c+1:] {
			if c1.species == c2.species {
				continue
			}
			c1.targets = append(c1.targets, c2)
			c2.targets = append(c2.targets, c1)
		}
	}
}
