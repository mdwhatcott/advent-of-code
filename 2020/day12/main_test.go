package advent

import (
	"strings"
	"testing"

	"advent/lib/intgrid"
)

var exampleInput = strings.Split(`F10
N3
F7
R90
F11`, "\n")

func TestStuff(t *testing.T) {
	end := part2(exampleInput)
	t.Log(end)
}

func TestRotate(t *testing.T) {
	t.Log("RIGHT:")
	p := intgrid.NewPoint(10, 4)
	for x := 0; x < 4; x++ {
		t.Log(p)
		p = RotateRight(p)
	}

	t.Log("LEFT:")
	p = intgrid.NewPoint(10, 4)
	for x := 0; x < 4; x++ {
		t.Log(p)
		p = RotateLeft(p)
	}
}
