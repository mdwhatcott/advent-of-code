package day17

import (
	"advent/lib/intgrid"
	"advent/lib/maths"
)

func SimulateLaunches(target intgrid.BoundingBox) (maxElevation, hits int) {
	for dx := 0; dx <= target.MaxX(); dx++ {
		for dy := -1000; dy < 1000; dy++ {
			probe := NewProbe(intgrid.NewDirection(dx, dy))
			targetHit := Launch(probe, target)
			if targetHit {
				hits++
				if probe.MaxY > maxElevation {
					maxElevation = probe.MaxY
				}
			}
		}
	}
	return maxElevation, hits
}

func Launch(probe *Probe, target intgrid.BoundingBox) (targetHit bool) {
	probe.Advance()

	if target.Contains(probe.Location) {
		return true
	}
	if probe.Location.X() > target.MaxX() {
		return false
	}
	if probe.Location.Y() < target.MinY() {
		return false
	}
	return Launch(probe, target)
}

type Probe struct {
	direction intgrid.Direction
	Location  intgrid.Point
	MaxY      int
}

func NewProbe(direction intgrid.Direction) *Probe {
	return &Probe{direction: direction}
}

func (this *Probe) Advance() {
	this.Location = this.Location.Move(this.direction)
	this.direction = intgrid.NewDirection(this.applyDrag(), this.applyGravity())
	this.MaxY = maths.Max(this.MaxY, this.Location.Y())
}

func (this *Probe) applyGravity() int {
	return this.direction.Dy() - 1
}

func (this *Probe) applyDrag() int {
	dx := this.direction.Dx()
	if dx > 0 {
		dx--
	} else if dx < 0 {
		dx++
	}
	return dx
}
