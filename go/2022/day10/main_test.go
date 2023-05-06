package day10

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/parse"
	"advent/lib/util"
)

var sampleLines = []string{
	"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8", "addx 13", "addx 4", "noop",
	"addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx -35",
	"addx 1", "addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop", "noop", "addx 21", "addx -15",
	"noop", "noop", "addx -3", "addx 9", "addx 1", "addx -3", "addx 8", "addx 1", "addx 5", "noop",
	"noop", "noop", "noop", "noop", "addx -36", "noop", "addx 1", "addx 7", "noop", "noop",
	"noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop", "noop", "addx 1", "noop",
	"noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7", "noop", "addx 1", "addx -33",
	"noop", "noop", "noop", "addx 2", "noop", "noop", "noop", "addx 8", "noop", "addx -1",
	"addx 2", "addx 1", "noop", "addx 17", "addx -9", "addx 1", "addx 1", "addx -3", "addx 11", "noop",
	"noop", "addx 1", "noop", "addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1", "addx 3",
	"addx 26", "addx -30", "addx 12", "addx -1", "addx 3", "addx 1", "noop", "noop", "noop", "addx -9",
	"addx 18", "addx 1", "addx 2", "noop", "noop", "addx 9", "noop", "noop", "noop", "addx -1",
	"addx 2", "addx -37", "addx 1", "addx 3", "noop", "addx 15", "addx -21", "addx 22", "addx -6", "addx 1",
	"noop", "addx 2", "addx 1", "noop", "addx -10", "noop", "noop", "addx 20", "addx 1", "addx 2",
	"addx 2", "addx -6", "addx -11", "noop", "noop", "noop",
}

func TestDay10(t *testing.T) {
	samplePart1, samplePart2 := Render(sampleLines)
	should.So(t, samplePart1, should.Equal, 13140)
	should.So(t, samplePart2, should.Equal, strings.Join([]string{"",
		"##  ##  ##  ##  ##  ##  ##  ##  ##  ##  ",
		"###   ###   ###   ###   ###   ###   ### ",
		"####    ####    ####    ####    ####    ",
		"#####     #####     #####     #####     ",
		"######      ######      ######      ####",
		"#######       #######       #######     ",
	}, "\n"))

	inputPart1, inputPart2 := Render(util.InputLines())
	should.So(t, inputPart1, should.Equal, 13060)
	should.So(t, inputPart2, should.Equal, strings.Join([]string{"",
		"####   ## #  # ###  #  # #    ###  #### ",
		"#       # #  # #  # #  # #    #  #    # ",
		"###     # #  # ###  #  # #    #  #   #  ",
		"#       # #  # #  # #  # #    ###   #   ",
		"#    #  # #  # #  # #  # #    # #  #    ",
		"#     ##   ##  ###   ##  #### #  # #### ",
	}, "\n")) // F J U B U L R Z
}

func Render(lines []string) (Part1 int, Part2 string) {
	const screenWidthInPixels = 40
	var (
		part2    strings.Builder
		cycle    = 1
		register = 1
		line     string
	)
	for len(lines) > 0 {
		if (cycle-20)%40 == 0 {
			Part1 += cycle * register
		}

		pixel := (cycle - 1) % screenWidthInPixels
		if pixel == 0 {
			part2.WriteString("\n")
		}
		if pixel-1 <= register && register <= pixel+1 {
			part2.WriteString("#")
		} else {
			part2.WriteString(" ")
		}

		if len(line) == 0 {
			line = lines[0]
			if line == "noop" {
				lines = lines[1:]
				line = ""
			}
			cycle++
		} else {
			register += parse.Int(strings.Fields(line)[1])
			lines = lines[1:]
			line = ""
			cycle++
		}
	}
	return Part1, part2.String()
}
