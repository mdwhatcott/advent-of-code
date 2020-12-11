package advent

import (
	"strings"
	"testing"
)

var input = strings.Split(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, "\n")

var expected = `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`

func TestExamplePart2(t *testing.T) {
	var grid [][]string
	for _, line := range input {
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	field := NewField(grid, 5, true)

	t.Log(0, field.String())

	for {
		field.Scan()
		modified := field.Update()
		t.Log(field.String())
		if modified == 0 {
			break
		}
	}

	final := field.String()
	if !strings.Contains(final, expected) {
		t.Error("Wrong ending field:\n", final)
	}

	occupied := field.CountOccupied()
	if occupied != 26 {
		t.Error("Expected 26, got:", occupied)
	}
}