package starter

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

func Test(t *testing.T) {
	t.Skip("not ready yet")
	inputLines := inputs.Read(2018, 15).Lines()
	t.Log("\n" + strings.Join(inputLines, "\n"))
	should.So(t, Part1(inputLines), should.Equal, TODO)
	should.So(t, Part2(inputLines), should.Equal, TODO)
}

func Part1(lines []string) any {
	return TODO
}
func Part2(lines []string) any {
	return TODO
}

////////////////////////////////////////////////////////////////////////////////

func TestSortReadingOrder(t *testing.T) {
	points := []*Unit{
		{point: NewPoint(9, 2)},
		{point: NewPoint(0, 3)},
		{point: NewPoint(4, 9)},
		{point: NewPoint(0, 4)},
		{point: NewPoint(7, 3)},
		{point: NewPoint(8, 8)},
		{point: NewPoint(8, 8)},
		{point: NewPoint(7, 9)},
		{point: NewPoint(6, 0)},
		{point: NewPoint(9, 4)},
	}

	SortReadingOrder(points)

	should.So(t, points, should.Equal, []*Unit{
		{point: NewPoint(6, 0)},
		{point: NewPoint(9, 2)},
		{point: NewPoint(0, 3)},
		{point: NewPoint(7, 3)},
		{point: NewPoint(0, 4)},
		{point: NewPoint(9, 4)},
		{point: NewPoint(8, 8)},
		{point: NewPoint(8, 8)},
		{point: NewPoint(4, 9)},
		{point: NewPoint(7, 9)},
	})
}
func TestParseCave(t *testing.T) {
	lines := []string{
		"#######",
		"#E..G.#",
		"#...#.#",
		"#.G.#G#",
		"#######",
	}
	elves, goblins, obstacles := ParseCave(lines)
	should.So(t, elves, should.Equal, []*Unit{
		NewUnit(1, 1),
	})
	should.So(t, goblins, should.Equal, []*Unit{
		NewUnit(4, 1),
		NewUnit(2, 3),
		NewUnit(5, 3),
	})
	should.So(t, obstacles, should.Equal, set.Of(
		// top side walls
		NewPoint(0, 0),
		NewPoint(1, 0),
		NewPoint(2, 0),
		NewPoint(3, 0),
		NewPoint(4, 0),
		NewPoint(5, 0),
		NewPoint(6, 0),

		// left side walls
		NewPoint(0, 1),
		NewPoint(0, 2),
		NewPoint(0, 3),

		// right side walls
		NewPoint(6, 1),
		NewPoint(6, 2),
		NewPoint(6, 3),

		// bottom side walls
		NewPoint(0, 4),
		NewPoint(1, 4),
		NewPoint(2, 4),
		NewPoint(3, 4),
		NewPoint(4, 4),
		NewPoint(5, 4),
		NewPoint(6, 4),

		// middle walls
		NewPoint(4, 2),
		NewPoint(4, 3),
	))
}
func TestNextMove_ExampleA(t *testing.T) {
	/*
		Targets:      In range:     Reachable:    Nearest:      Chosen:
		#######       #######       #######       #######       #######
		#E..G.#       #E.?G?#       #E.@G.#       #E.!G.#       #E.+G.#
		#...#.#  -->  #.?.#?#  -->  #.@.#.#  -->  #.!.#.#  -->  #...#.#
		#.G.#G#       #?G?#G#       #@G@#G#       #!G.#G#       #.G.#G#
		#######       #######       #######       #######       #######
	*/
	elves, goblins, walls := ParseCave([]string{
		"#######",
		"#E..G.#",
		"#...#.#",
		"#.G.#G#",
		"#######",
	})
	MoveActor(elves[0], goblins, walls.Union(set.Of(Points(goblins)...)))

	should.So(t, elves[0].Point(), should.Equal, NewPoint(2, 1))
}
func TestNextMove_ExampleB(t *testing.T) {
	/*
		In range:     Nearest:      Chosen:       Distance:     Step:
		#######       #######       #######       #######       #######
		#.E...#       #.E...#       #.E...#       #4E212#       #..E..#
		#...?.#  -->  #...!.#  -->  #...+.#  -->  #32101#  -->  #.....#
		#..?G?#       #..!G.#       #...G.#       #432G2#       #...G.#
		#######       #######       #######       #######       #######
	*/
	elves, goblins, walls := ParseCave([]string{
		"#######",
		"#.E...#",
		"#.....#",
		"#...G.#",
		"#######",
	})
	MoveActor(elves[0], goblins, walls.Union(set.Of(Points(goblins)...)))

	should.So(t, elves[0].Point(), should.Equal, NewPoint(3, 1))
}
