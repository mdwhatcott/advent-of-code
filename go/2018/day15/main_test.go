package starter

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

func Test(t *testing.T) {
	inputLines := inputs.Read(2018, 15).Lines()
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
