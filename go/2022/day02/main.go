package day02

import "advent/lib/util"

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

func Part1() int { return playTournament(part1Outcomes) }
func Part2() int { return playTournament(part2Outcomes) }

func playTournament(outcomes map[string]int) (sum int) {
	for _, line := range util.InputLines() {
		sum += outcomes[line]
	}
	return sum
}
