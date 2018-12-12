package day06

import (
	"math"
	"strings"

	"advent/lib/grid"
	"advent/lib/util"
)

func calculateAreaWithinRadius(points []grid.Point, radius float64) (within int) {
	box := NewBoundingBox(points)
	for x := box.minX; x <= box.maxX; x++ {
		for y := box.minY; y <= box.maxY; y++ {
			current := grid.NewPoint(x, y)
			var sum float64
			for _, point := range points {
				sum += grid.ManhattanDistance(current, point)
			}
			if sum < radius {
				within++
			}
		}
	}
	return within
}

func calculateLargestFiniteArea(points []grid.Point) (max float64) {
	areas := make(map[grid.Point]float64)
	box := NewBoundingBox(points)
	for _, point := range points {
		if box.IsOnBoundary(point) {
			areas[point] = math.Inf(1)
		}
	}
	for x := box.minX; x <= box.maxX; x++ {
		for y := box.minY; y <= box.maxY; y++ {
			distances := make(map[grid.Point]float64)
			here := grid.NewPoint(x, y)
			for _, point := range points {
				distance := grid.ManhattanDistance(point, here)
				distances[point] = distance
			}
			shortest := findShortestPoints(distances)
			if len(shortest) == 1 {
				if candidate := shortest[0]; !box.IsOnBoundary(candidate) {
					areas[candidate]++
				}
			}
		}
	}

	for point, area := range areas {
		if !box.IsOnBoundary(point) && area > max {
			max = area
		}
	}

	return max
}

func findShortestPoints(distances map[grid.Point]float64) (closest []grid.Point) {
	min := 1000.0
	for _, distance := range distances {
		if distance < min {
			min = distance
		}
	}
	for point, distance := range distances {
		if distance == min {
			closest = append(closest, point)
		}
	}
	return closest
}

func parsePoints(input string) (points []grid.Point) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.Replace(line, ",", "", -1)
		fields := strings.Fields(line)
		point := grid.NewPoint(
			util.ParseFloat(fields[0]),
			util.ParseFloat(fields[1]),
		)
		points = append(points, point)
	}
	return points
}
