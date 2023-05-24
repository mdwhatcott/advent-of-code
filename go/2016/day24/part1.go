package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
	"github.com/mdwhatcott/astar"
)

type Turtle struct {
	maze    []string
	point   grid.Point
	visited string
}

func StartPoint(mazeInput []string) (start *Turtle) {
	return &Turtle{
		maze:    mazeInput,
		point:   startPoint(mazeInput),
		visited: "0",
	}
}

func startPoint(maze []string) grid.Point {
	for r, row := range maze {
		for c, char := range row {
			if char == '0' {
				return grid.NewPoint(float64(c), float64(r))
			}
		}
	}

	log.Fatal("Could not find '0'")
	return grid.Point{}
}

func (this Turtle) AdjacentPositions() (neighbors []astar.Turtle) {
	for _, adjacent := range this.point.Neighbors4() {
		if c := this.at(adjacent); c != "#" {
			visited := this.visited
			if !strings.Contains(this.visited, c) && c != "." {
				visited = Sort(visited + c)
			}
			neighbor := &Turtle{point: adjacent, maze: this.maze, visited: visited}
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (this Turtle) EstimatedDistanceToTarget() float64 {
	return float64(8 - len(this.visited)) // len("01234567") == 8
}

func (this Turtle) StepCost() float64 {
	return 1.0
}

func (this Turtle) at(p grid.Point) string {
	return string(this.maze[int(p.Y())][int(p.X())])
}

func (this Turtle) Hash() string {
	hash := sha1.New()
	fmt.Fprint(hash, this.point)
	fmt.Fprint(hash, this.visited)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sort(v string) string {
	var a []string
	for _, x := range v {
		a = append(a, string(x))
	}
	sort.Strings(a)
	return strings.Join(a, "")
}
