package day11

import (
	"strings"

	"advent/lib/util"
)

func Answers() (part1, part2 int) {
	var maxDistance int
	var start Hex
	var position Hex

	path := strings.Split(util.InputString(), ",")

	for _, direction := range path {
		position = position.Offset(Offsets[direction])
		if distance := position.DistanceTo(start); distance > maxDistance {
			maxDistance = distance
		}
	}

	return position.DistanceTo(start), maxDistance
}

var Offsets = map[string]Hex{
	"n":  North,
	"ne": NorthEast,
	"nw": NorthWest,
	"s":  South,
	"se": SouthEast,
	"sw": SouthWest,
}
