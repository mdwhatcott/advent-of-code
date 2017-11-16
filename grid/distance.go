package grid

import "math"

func CityBlockDistance(p, q Point) float64 {
	x, y := diff(p, q)
	return math.Abs(x) + math.Abs(y)
}

func EuclideanDistance(p, q Point) float64 {
	return math.Hypot(diff(p, q))
}

func diff(p, q Point) (x float64, y float64) {
	return p.X() - q.X(), p.Y() - q.Y()
}
