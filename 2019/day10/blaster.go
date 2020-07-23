package advent

type Blaster struct {
	field AsteroidField
	angle float64
}

func NewBlaster(field AsteroidField) *Blaster {
	return &Blaster{
		field: field,
		angle: -1,
	}
}

func (this *Blaster) Angle() float64 {
	angle := this.angle
	for angle > 360 {
		angle -= 360
	}
	return angle
}

func (this *Blaster) SetAngle(update float64) {
	for update < this.angle {
		update += 360
	}
	this.angle = update
}

func (this *Blaster) Aim() (index int) {
	if len(this.field) == 0 {
		return -1
	}

	minAngle := 361.0
	minAsteroid := -1
	minDistance := 1000.0

	for i, a := range this.field {
		angleA := a.AngleFromOrigin()
		if angleA <= this.Angle() {
			continue
		}
		if angleA < minAngle {
			minAngle = angleA
			minAsteroid = i
		}
		distanceA := a.DistanceFromOrigin()
		if angleA == minAngle && distanceA < minDistance {
			minAsteroid = i
			minDistance = distanceA
		}
	}

	if minAsteroid == -1 && len(this.field) > 0 {
		this.angle -= 360
		return this.Aim()
	}

	this.SetAngle(minAngle)
	return minAsteroid
}

func (this *Blaster) Fire(i int) {
	this.field = append(this.field[:i], this.field[i+1:]...)
}

func (this *Blaster) Field() AsteroidField {
	return this.field
}
