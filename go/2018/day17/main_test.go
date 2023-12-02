package day17

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/testing/should"
)

func TestDay17(t *testing.T) {
	t.Parallel()

	should.So(t, Part1(), should.Equal, nil)
	should.So(t, Part2(), should.Equal, nil)
}

var example1Lines = []string{
	"x=495, y=2..7",
	"y=7, x=495..501",
	"x=501, y=3..7",
	"x=498, y=2..4",
	"x=506, y=1..2",
	"x=498, y=10..13",
	"x=504, y=10..13",
	"y=13, x=498..504",
}

func TestParseInput(t *testing.T) {
	cave := ParseInput(example1Lines)
	should.So(t, cave, should.Equal, set.Of(
		XY(495, 2), XY(495, 3), XY(495, 4), XY(495, 5), XY(495, 6), XY(495, 7),
		XY(496, 7), XY(497, 7), XY(498, 2), XY(498, 3), XY(498, 4), XY(498, 7),
		XY(498, 10), XY(498, 11), XY(498, 12), XY(498, 13), XY(499, 7), XY(499, 13),
		XY(500, 7), XY(500, 13), XY(501, 3), XY(501, 4), XY(501, 5), XY(501, 6),
		XY(501, 7), XY(501, 13), XY(502, 13), XY(503, 13), XY(504, 10), XY(504, 11),
		XY(504, 12), XY(504, 13), XY(506, 1), XY(506, 2)),
	)
	rendered := Display(cave, nil)
	t.Log("\n" + rendered)
	should.So(t, strings.Split(rendered, "\n"), should.Equal, []string{
		"............",
		"...........#",
		"#..#.......#",
		"#..#..#.....",
		"#..#..#.....",
		"#.....#.....",
		"#.....#.....",
		"#######.....",
		"............",
		"............",
		"...#.....#..",
		"...#.....#..",
		"...#.....#..",
		"...#######..",
	})
}
func TestDrip(t *testing.T) {
	cave := ParseInput(example1Lines)
	pools := DumpWater(cave)
	should.So(t, len(pools), should.Equal, 57)
	t.Log("\n" + Display(cave, pools))
}
