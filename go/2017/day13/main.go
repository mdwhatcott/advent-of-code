package day13

import (
	"sort"
	"strings"

	"advent/lib/parse"
	"advent/lib/util"
)

// Inspired by: https://github.com/lukaszroz/advent-of-code-2017/blob/master/day13.go

func Answers() (part1, part2 int) {
	firewall := loadFirewall(util.InputScanner())
	return CalculateSeverityOfInitialTrip(firewall), CalculateDelayNecessaryForSafeJourney(firewall)
}

func loadFirewall(scanner *util.Scanner) (firewall Firewall) {
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ": ")
		firewall = append(firewall, Layer{
			Level:    parse.Int(fields[0]),
			Width:    parse.Int(fields[1]),
			WidthMod: (parse.Int(fields[1]) - 1) * 2,
		})
	}
	firewall.Sort()
	return firewall
}

func CalculateSeverityOfInitialTrip(firewall Firewall) (total int) {
	for _, layer := range firewall {
		if layer.TripsAlarm(0) {
			total += layer.Level * layer.Width
		}
	}
	return total
}

func CalculateDelayNecessaryForSafeJourney(firewall Firewall) int {
	for delay := 0; ; delay++ {
		if !firewall.AlarmTrippedAt(delay) {
			return delay
		}
	}
}

type Firewall []Layer

func (this Firewall) Sort() {
	sort.Slice(this, func(i, j int) bool { return this[i].WidthMod < this[j].WidthMod })
}

func (firewall Firewall) AlarmTrippedAt(delay int) bool {
	for _, layer := range firewall {
		if layer.TripsAlarm(delay) {
			return true
		}
	}
	return false
}

type Layer struct {
	Level    int
	Width    int
	WidthMod int
}

func (this Layer) TripsAlarm(at int) bool {
	return (this.Level+at)%this.WidthMod == 0
}
