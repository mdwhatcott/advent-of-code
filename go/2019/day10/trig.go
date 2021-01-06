package advent

import "math"

func (this Asteroid) DistanceFromOrigin() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y)
}

func (this Asteroid) AngleFromOrigin() (result float64) {
	if this.X >= 0 && this.Y < 0 {
		return 0 + degrees(math.Atan(math.Abs(this.X)/math.Abs(this.Y)))
	} else if this.X >= 0 && this.Y >= 0 {
		return 90 + degrees(math.Atan(math.Abs(this.Y)/math.Abs(this.X)))
	} else if this.X < 0 && this.Y > 0 {
		return 180 + degrees(math.Atan(math.Abs(this.X)/math.Abs(this.Y)))
	} else {
		return 270 + degrees(math.Atan(math.Abs(this.Y)/math.Abs(this.X)))
	}
}

func degrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
