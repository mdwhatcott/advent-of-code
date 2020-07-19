package advent

import "sort"

type Blaster struct {
	field AsteroidField
	angle float64
}

func NewBlaster(field AsteroidField) *Blaster {
	return &Blaster{field: field}
}

func (this *Blaster) Aim() (index int) {
	sort.Slice(this.field, func(i, j int) bool {
		angleI := this.field[i].AngleFromOrigin()
		angleJ := this.field[j].AngleFromOrigin()
		return angleI < angleJ
	})

	if len(this.field) == 0 {
		return -1
	}
	if len(this.field) == 1 {
		return 0
	}

	// TODO: select the closest of asteroids at the front of the list that have the same angle
	return -1
}

func (this *Blaster) Fire(i int) {
	this.field = append(this.field[:i], this.field[i:]...)
}

func (this *Blaster) Field() AsteroidField {
	return this.field
}


