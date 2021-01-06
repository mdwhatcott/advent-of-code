package main

import (
	"strconv"
	"strings"
)

type Disc struct {
	Delay     int
	Positions int
	Start     int
}

func (this Disc) PositionAtTime(t int) int {
	return (this.Start + t + this.Delay) % this.Positions
}

func ParseDisc(line string) Disc {
	line = strings.TrimRight(line, ".")
	fields := strings.Fields(line)
	delay, _ := strconv.Atoi(fields[1][1:])
	positions, _ := strconv.Atoi(fields[3])
	start, _ := strconv.Atoi(fields[len(fields)-1])

	return Disc{
		Delay:     delay,
		Positions: positions,
		Start:     start,
	}
}
