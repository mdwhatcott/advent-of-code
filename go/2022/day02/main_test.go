package day02

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"A Y",
		"B X",
		"C Z",
	}
)

func TestDay02(t *testing.T) {
	should.So(t, playTournament(sampleLines, part1Outcomes), should.Equal, 15)
	should.So(t, playTournament(sampleLines, part2Outcomes), should.Equal, 12)

	should.So(t, playTournament(inputLines, part1Outcomes), should.Equal, 10994)
	should.So(t, playTournament(inputLines, part2Outcomes), should.Equal, 12526)
}
func playTournament(rounds []string, outcomes map[string]int) (sum int) {
	for _, round := range rounds {
		sum += outcomes[round]
	}
	return sum
}

// rock 1, paper 2, scissors 3
// loss 0, draw 3, win 6
// A rock  B paper  C scissors
// X rock  Y paper  Z scissors
var part1Outcomes = map[string]int{
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

// x lose  y draw  z win
var part2Outcomes = map[string]int{
	"A X": 0 + 3,
	"A Y": 3 + 1,
	"A Z": 6 + 2,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 0 + 2,
	"C Y": 3 + 3,
	"C Z": 6 + 1,
}
