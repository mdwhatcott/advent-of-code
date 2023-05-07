package day20

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
)

func ParsePixel(id int, line string) *Pixel {
	pixel := &Pixel{ID: id}
	pixel.Point = ParseFirstTriplet(line)
	line = line[strings.Index(line, ">")+1:]
	pixel.Velocity = ParseFirstTriplet(line)
	line = line[strings.Index(line, ">")+1:]
	pixel.Acceleration = ParseFirstTriplet(line)
	return pixel
}

func ParseFirstTriplet(line string) Triplet {
	start := strings.Index(line, "<")
	stop := strings.Index(line, ">")
	excerpt := line[start+1 : stop]
	fields := strings.Split(excerpt, ",")
	return Triplet{
		X: parse.Int(fields[0]),
		Y: parse.Int(fields[1]),
		Z: parse.Int(fields[2]),
	}
}

type Pixel struct {
	ID           int
	Point        Triplet
	Velocity     Triplet
	Acceleration Triplet
}

type Triplet struct{ X, Y, Z int }

func (this Triplet) Apply(that Triplet) Triplet {
	return Triplet{
		X: this.X + that.X,
		Y: this.Y + that.Y,
		Z: this.Z + that.Z,
	}
}

func (this *Pixel) Update() {
	this.Velocity = this.Velocity.Apply(this.Acceleration)
	this.Point = this.Point.Apply(this.Velocity)
}

func (this *Pixel) DistanceFromOrigin() int {
	return abs(this.Point.X) + abs(this.Point.Y) + abs(this.Point.Z)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func detectCollisions(pixels []*Pixel) map[Triplet][]int {
	collisions := make(map[Triplet][]int)
	for _, pixel := range pixels {
		collisions[pixel.Point] = append(collisions[pixel.Point], pixel.ID)
	}
	for point := range collisions {
		if len(collisions[point]) == 1 {
			delete(collisions, point)
		}
	}
	return collisions
}
