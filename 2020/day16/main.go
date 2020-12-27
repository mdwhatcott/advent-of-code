package advent

import (
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	return 0
}

func Part2() interface{} {
	return 0
}

type FieldDefinition struct {
	Name        string
	ValidValues map[int]bool
}

func ParseFieldDefinition(line string) (def FieldDefinition) {
	def.ValidValues = make(map[int]bool)
	line = strings.Replace(line, " or ", " ", 1)
	halves := strings.Split(line, ": ")
	def.Name = halves[0]
	line = halves[1]
	fields := strings.Fields(line)
	for _, field := range fields {
		bounds := strings.Split(field, "-")
		lower := util.ParseInt(bounds[0])
		upper := util.ParseInt(bounds[1])
		for x := lower; x <= upper; x++ {
			def.ValidValues[x] = true
		}
	}
	return def
}

func (this FieldDefinition) WithinRange(candidate int) bool {
	_, ok := this.ValidValues[candidate]
	return ok
}

func (this FieldDefinition) AllWithinRange(candidates ...int) bool {
	for _, candidate := range candidates {
		if !this.WithinRange(candidate) {
			return false
		}
	}
	return true
}
