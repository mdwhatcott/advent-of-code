package day15

import (
	"math"
	"strings"

	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/strconvmust"
)

// findDistressBeacon implements the algorithm visualized here:
// https://www.reddit.com/r/adventofcode/comments/zmfwg1/2022_day_15_part_2_seekin_for_the_beacon/
// This code was originally generated with the 'help' of chat-gpt, but it was so incorrect I had to rework most of it.
// https://chat.openai.com/c/cf3ea31c-7cc1-4308-9006-0cebd6b8e495
func findDistressBeacon(lines []string, min, max int) Point {
	var sensors []Point
	var beacons []Point

	for _, line := range lines {
		parts := strings.Fields(replacer.Replace(line))
		sensors = append(sensors, Point{X: strconvmust.Atoi(parts[3]), Y: strconvmust.Atoi(parts[5])})
		beacons = append(beacons, Point{X: strconvmust.Atoi(parts[11]), Y: strconvmust.Atoi(parts[13])})
	}

	var distinctPoints []Point
	for point := range generateOuterPoints(sensors, beacons) {
		if isValid(point, min, max) {
			if isDistinctBeaconLocation(point, sensors, beacons) {
				distinctPoints = append(distinctPoints, point)
			}
		}
	}

	for _, point := range distinctPoints {
		return point
	}
	panic("not found")
}

type Point struct{ X, Y int }

func manhattanDistance(a, b Point) (result int) {
	result += int(math.Abs(float64(a.X - b.X)))
	result += int(math.Abs(float64(a.Y - b.Y)))
	return result
}
func isValid(p Point, min, max int) bool {
	if p.X < min {
		return false
	}
	if p.Y < min {
		return false
	}
	if p.X > max {
		return false
	}
	if p.Y > max {
		return false
	}
	return true
}
func traceExteriorPerimeter(sensor Point, minDistance int) (result []Point) {
	for n := 0; n < minDistance+1; n++ {
		result = append(result,
			Point{X: sensor.X + n, Y: sensor.Y - (minDistance + 1) + n},
			Point{X: sensor.X - n, Y: sensor.Y + (minDistance + 1) - n},
			Point{X: sensor.X + (minDistance + 1) - n, Y: sensor.Y + n},
			Point{X: sensor.X - (minDistance + 1) + n, Y: sensor.Y - n},
		)
	}
	return result
}
func generateOuterPoints(sensors, beacons []Point) set.Set[Point] {
	result := set.Of[Point]()
	for n, sensor := range sensors {
		for _, p := range traceExteriorPerimeter(sensor, manhattanDistance(sensor, beacons[n])) {
			result.Add(p)
		}
	}
	return result
}

func isDistinctBeaconLocation(point Point, sensors, beacons []Point) bool {
	for n, sensor := range sensors {
		closestBeacon := beacons[n]
		minDistance := manhattanDistance(sensor, closestBeacon)
		if manhattanDistance(sensor, point) <= minDistance {
			return false
		}
	}
	return true
}

var replacer = strings.NewReplacer(
	"=", " ",
	",", " ",
	":", " ",
)
