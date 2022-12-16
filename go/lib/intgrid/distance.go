package intgrid

import "math"

func ManhattanDistance(p, q Point) int {
	return CityBlockDistance(p, q)
}

func CityBlockDistance(p, q Point) int {
	x, y := diff(p, q)
	return int(math.Abs(x)) + int(math.Abs(y))
}

func EuclideanDistance(p, q Point) float64 {
	return math.Hypot(diff(p, q))
}

func diff(p, q Point) (x float64, y float64) {
	return float64(p.X() - q.X()), float64(p.Y() - q.Y())
}
