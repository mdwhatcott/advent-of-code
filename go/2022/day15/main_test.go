package day15

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/intgrid"
	"advent/lib/maths"
	"advent/lib/util"
)

var (
	inputLines  = util.InputLines()
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
}

func Part1(lines []string, targetY int) (result int) {
	maxDistance, minX, maxX := 0, 0xFFFFFFFF, -0xFFFFFFFF
	sensors := make(map[intgrid.Point]int)
	beacons := set.New[intgrid.Point](0)
	for _, line := range lines {
		fields := strings.Fields(strings.NewReplacer("=", " ", ",", " ", ":", " ").Replace(line))
		sensor := intgrid.NewPoint(util.ParseInt(fields[3]), util.ParseInt(fields[5]))
		beacon := intgrid.NewPoint(util.ParseInt(fields[11]), util.ParseInt(fields[13]))
		x := maths.Min(sensor.X(), beacon.X())
		if x < minX {
			minX = sensor.X()
		}
		if x > maxX {
			maxX = sensor.X()
		}
		beacons.Add(beacon)
		distance := intgrid.ManhattanDistance(sensor, beacon)
		if distance > maxDistance {
			maxDistance = distance
		}
		sensors[sensor] = distance
	}
	for x := minX - maxDistance; x < maxX+maxDistance; x++ {
		p := intgrid.NewPoint(x, targetY)
		if beacons.Contains(p) {
			continue
		}
		for sensor, toClosetsBeacon := range sensors {
			if intgrid.ManhattanDistance(p, sensor) <= toClosetsBeacon {
				result++
				break
			}
		}
	}
	return result
}
