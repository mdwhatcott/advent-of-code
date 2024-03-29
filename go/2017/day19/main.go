package day19

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() string {
	turtle := NewTurtle(strings.Split(string(util.InputBytes()), "\n"))
	for turtle.Orient() {
		turtle.Move()
	}
	return turtle.sequence
}

func Part2() (steps int) {
	turtle := NewTurtle(strings.Split(string(util.InputBytes()), "\n"))
	for turtle.Orient() {
		turtle.Move()
		steps++
	}
	return steps
}
