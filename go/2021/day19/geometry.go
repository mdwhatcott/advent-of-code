package day19

type rot func(x, y, z int) (X, Y, Z int)

var faces = []rot{
	func(x, y, z int) (X, Y, Z int) { return x, y, z },
	func(x, y, z int) (X, Y, Z int) { return x, -y, -z },
	func(x, y, z int) (X, Y, Z int) { return x, -z, y },
	func(x, y, z int) (X, Y, Z int) { return -y, -z, x },
	func(x, y, z int) (X, Y, Z int) { return y, -z, -x },
	func(x, y, z int) (X, Y, Z int) { return -x, -z, -y },
}
var spins = []rot{
	func(x, y, z int) (X, Y, Z int) { return x, y, z },
	func(x, y, z int) (X, Y, Z int) { return -y, x, z },
	func(x, y, z int) (X, Y, Z int) { return -x, -y, z },
	func(x, y, z int) (X, Y, Z int) { return y, -x, z },
}

func Rotate(i int, p Point) Point {
	return NewPoint(spins[i%4](faces[i/4](p.X(), p.Y(), p.Z())))
}

func RotateAll(i int, points ...Point) (results []Point) {
	for _, point := range points {
		results = append(results, Rotate(i, point))
	}
	return results
}
