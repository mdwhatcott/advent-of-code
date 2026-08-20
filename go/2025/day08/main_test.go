package day08

import (
	"bufio"
	"cmp"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Run("part 1", func(t *testing.T) { execute(t, false) })
	t.Run("part 2", func(t *testing.T) { execute(t, true) })
}

func execute(t *testing.T, part2 bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	// gather all points
	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			t.Log(err)
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			t.Log(err)
		}
		z, err := strconv.Atoi(fields[2])
		if err != nil {
			t.Log(err)
		}
		points = append(points, Point{X: x, Y: y, Z: z})
	}
	// calculate distance between all pairs
	var pairs []Pair
	for a := range points {
		A := points[a]
		for b := a + 1; b < len(points); b++ {
			B := points[b]
			pairs = append(pairs, Pair{P1: A, P2: B, Distance: CalculateDistance(A, B)})
		}
	}
	slices.SortStableFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.Distance, b.Distance)
	})
	t.Log("how many pairs?", len(pairs))

	// Part 1: connect 1000 closest pairs
	if !part2 {
		pairs = pairs[:1000]
	}

	allPoints := make(map[Point]struct{})

	// each unique point in the pairs starts on its own circuit
	var circuits []map[Point]struct{}
	for _, pair := range pairs {
		_, p1ok := allPoints[pair.P1]
		if !p1ok {
			allPoints[pair.P1] = struct{}{}
			circuits = append(circuits, map[Point]struct{}{pair.P1: {}})
		}
		_, p2ok := allPoints[pair.P2]
		if !p2ok {
			allPoints[pair.P2] = struct{}{}
			circuits = append(circuits, map[Point]struct{}{pair.P2: {}})
		}
	}
	t.Log("How many unique points in the involved pairs?", len(circuits)) // 862 for part 1

	// connect each pair, merging circuits as needed
	for _, pair := range pairs {
		// locate the circuit for each point in the pair
		p1Index := -1
		p2Index := -1
		for c, circuit := range circuits {
			_, p1ok := circuit[pair.P1]
			if p1ok {
				p1Index = c
			}
			_, p2ok := circuit[pair.P2]
			if p2ok {
				p2Index = c
			}
		}
		if p1Index < 0 || p2Index < 0 || p1Index == p2Index {
			continue
		}
		// merge the second point's circuit into the first point's circuit
		maps.Copy(circuits[p1Index], circuits[p2Index])
		circuits = slices.Delete(circuits, p2Index, p2Index+1)

		if part2 && len(circuits) == 1 {
			part2Answer := pair.P1.X * pair.P2.X
			if part2Answer != 31182420 {
				t.Fatal("expected 31182420, got:", part2Answer)
			}
			t.Log("[PART 2] product of final pair's X values:", part2Answer)
			return
		}
	}

	t.Log("How many circuits from the closest 1000 pairs?", len(circuits))
	var counts []int
	for _, circuit := range circuits {
		counts = append(counts, len(circuit))
	}
	slices.SortFunc(counts, func(a, b int) int { return -cmp.Compare(a, b) })
	if len(counts) >= 3 {
		t.Log("3 largest circuits:", counts[:3])
		part1Answer := counts[0] * counts[1] * counts[2]
		if part1Answer != 79560 {
			t.Fatal("expected 79560 for part 1, got:", part1Answer)
		}
		t.Log("[PART 1] product of the sizes of the 3 largest circuits:", part1Answer)
	}
}

type Pair struct {
	P1, P2   Point
	Distance float64
}

type Point struct {
	X, Y, Z int
}

func CalculateDistance(p1, p2 Point) float64 {
	return math.Sqrt(
		math.Pow(float64(p2.X-p1.X), 2) +
			math.Pow(float64(p2.Y-p1.Y), 2) +
			math.Pow(float64(p2.Z-p1.Z), 2),
	)
}
