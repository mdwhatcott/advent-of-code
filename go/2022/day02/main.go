package day02

import (
	"strings"

	"advent/lib/util"
)

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

func Part1() (sum int) {
	for _, line := range util.InputLines() {
		sum += part1Outcomes[line]
	}
	return sum
}

func Part2() (sum int) {
	for _, line := range util.InputLines() {
		a := parseFirstShape(line)
		b := chooseSecondShape(line, a)
		_, B := ScoreRound(a, b)
		sum += B
	}
	return sum
}

func chooseSecondShape(line string, a Shape) Shape {
	fields := strings.Fields(line)
	switch fields[1] {
	case "X":
		switch a {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	case "Y":
		return a
	case "Z":
		switch a {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}
	panic("invalid line")
}

func parseFirstShape(line string) Shape {
	switch strings.Fields(line)[0] {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}
	panic("invalid line")
}

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

func ScoreRound(a, b Shape) (A, B int) {
	switch a {
	case Rock:
		A += 1
	case Paper:
		A += 2
	case Scissors:
		A += 3
	}
	switch b {
	case Rock:
		B += 1
	case Paper:
		B += 2
	case Scissors:
		B += 3
	}
	if a == b {
		A += 3
		B += 3
	} else if a == Rock && b == Scissors {
		A += 6
	} else if a == Scissors && b == Paper {
		A += 6
	} else if a == Paper && b == Rock {
		A += 6
	} else if b == Rock && a == Scissors {
		B += 6
	} else if b == Scissors && a == Paper {
		B += 6
	} else if b == Paper && a == Rock {
		B += 6
	}
	return A, B
}
