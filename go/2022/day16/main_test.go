package day16

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	_ "github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

var (
	inputLines  = inputs.Read(TODO, TODO).Lines()
	sampleLines = []string{
		fmt.Sprint(TODO),
	}
)

func Test(t *testing.T) {
	should.So(t, Part1(sampleLines), should.Equal, TODO)
	should.So(t, Part1(inputLines), should.Equal, TODO)

	should.So(t, Part2(sampleLines), should.Equal, TODO)
	should.So(t, Part2(inputLines), should.Equal, TODO)
}

func Part1(lines []string) any {
	return TODO
}

func Part2(lines []string) any {
	return TODO
}
