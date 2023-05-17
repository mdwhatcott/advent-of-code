package day15

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/advent-of-code/go/lib/maths"
	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2022, 15).Lines()
	sampleLines = []string{
		"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
		"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
		"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
		"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
		"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
		"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
		"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
		"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
		"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
		"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
		"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
		"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
		"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
		"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
	}
)

func TestDay15(t *testing.T) {
	should.So(t, Part1(sampleLines, 10), should.Equal, 26)
	should.So(t, Part1(inputLines, 2000000), should.Equal, 4424278)

	should.So(t, Part2(sampleLines, 0, 20), should.Equal, 56000011)
	should.So(t, Part2(inputLines, 0, 4_000_000), should.Equal, 0)
}

func Part1(lines []string, targetY int) (result int) {
	maxDistance, minX, maxX := 0, 0xFFFFFFFF, -0xFFFFFFFF
	sensors := make(map[Point]int)
	beacons := set.Of[Point]()
	for _, line := range lines {
		fields := strings.Fields(strings.NewReplacer("=", " ", ",", " ", ":", " ").Replace(line))
		sensor := Point{X: parse.Int(fields[3]), Y: parse.Int(fields[5])}
		beacon := Point{X: parse.Int(fields[11]), Y: parse.Int(fields[13])}
		x := maths.Min(sensor.X, beacon.X)
		if x < minX {
			minX = sensor.X
		}
		if x > maxX {
			maxX = sensor.X
		}
		beacons.Add(beacon)
		distance := manhattanDistance(sensor, beacon)
		if distance > maxDistance {
			maxDistance = distance
		}
		sensors[sensor] = distance
	}
	for x := minX - maxDistance; x < maxX+maxDistance; x++ {
		p := Point{X: x, Y: targetY}
		if beacons.Contains(p) {
			continue
		}
		for sensor, toClosetsBeacon := range sensors {
			if manhattanDistance(p, sensor) <= toClosetsBeacon {
				result++
				break
			}
		}
	}
	return result
}

func Part2(lines []string, min, max int) (result int) {
	point := findDistressBeacon(lines, min, max)
	fmt.Println(point)
	return point.X*4_000_000 + point.Y
}

func TestTraceExteriorPerimeter(t *testing.T) {
	sensor := Point{X: 8, Y: 7}
	beacon := Point{X: 2, Y: 10}
	actual := traceExteriorPerimeter(sensor, manhattanDistance(sensor, beacon))
	should.So(t, set.Of(actual...), should.Equal, set.Of(
		Point{X: 8, Y: -3},
		Point{X: 9, Y: -2},
		Point{X: 10, Y: -1},
		Point{X: 11, Y: 0},
		Point{X: 12, Y: 1},
		Point{X: 13, Y: 2},
		Point{X: 14, Y: 3},
		Point{X: 15, Y: 4},
		Point{X: 16, Y: 5},
		Point{X: 17, Y: 6},

		Point{X: 8, Y: 17},
		Point{X: 7, Y: 16},
		Point{X: 6, Y: 15},
		Point{X: 5, Y: 14},
		Point{X: 4, Y: 13},
		Point{X: 3, Y: 12},
		Point{X: 2, Y: 11},
		Point{X: 1, Y: 10},
		Point{X: 0, Y: 9},
		Point{X: -1, Y: 8},

		Point{X: -2, Y: 7},
		Point{X: -1, Y: 6},
		Point{X: 0, Y: 5},
		Point{X: 1, Y: 4},
		Point{X: 2, Y: 3},
		Point{X: 3, Y: 2},
		Point{X: 4, Y: 1},
		Point{X: 5, Y: 0},
		Point{X: 6, Y: -1},
		Point{X: 7, Y: -2},

		Point{X: 18, Y: 7},
		Point{X: 17, Y: 8},
		Point{X: 16, Y: 9},
		Point{X: 15, Y: 10},
		Point{X: 14, Y: 11},
		Point{X: 13, Y: 12},
		Point{X: 12, Y: 13},
		Point{X: 11, Y: 14},
		Point{X: 10, Y: 15},
		Point{X: 9, Y: 16},
	))
}
