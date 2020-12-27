package advent

import (
	"strings"
	"unicode"

	"advent/lib/util"
)

func Part1() interface{} {
	input := util.InputString()
	definitions := ParseAllFieldDefinitions(input)
	tickets := ParseAllTickets(input)
	return CalculateErrorRate(definitions, tickets)
}

func Part2() interface{} {
	input := util.InputString()
	definitions := ParseAllFieldDefinitions(input)
	tickets := FilterValidTickets(definitions, ParseAllTickets(input))
	candidates := IdentifyFieldPlacementCandidates(definitions, tickets)
	finalized := FinalizeFieldPlacements(candidates)
	return CalculateDepartureProduct(finalized, tickets[0])
}

func ParseAllFieldDefinitions(input string) (result []*FieldDefinition) {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		result = append(result, ParseFieldDefinition(line))
	}
	return result
}

type FieldDefinition struct {
	Name        string
	ValidValues map[int]bool
	Placement   map[int]bool
}

func ParseFieldDefinition(line string) (definition *FieldDefinition) {
	definition = &FieldDefinition{
		ValidValues: make(map[int]bool),
		Placement:   make(map[int]bool),
	}
	line = strings.Replace(line, " or ", " ", 1)
	halves := strings.Split(line, ": ")
	definition.Name = halves[0]
	line = halves[1]
	fields := strings.Fields(line)
	for _, field := range fields {
		bounds := strings.Split(field, "-")
		lower := util.ParseInt(bounds[0])
		upper := util.ParseInt(bounds[1])
		for x := lower; x <= upper; x++ {
			definition.ValidValues[x] = true
		}
	}
	return definition
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

func ParseAllTickets(input string) (result [][]int) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 && unicode.IsDigit(rune(line[0])) {
			result = append(result, util.ParseInts(strings.Split(line, ",")))
		}
	}
	return result
}

func CalculateErrorRate(definitions []*FieldDefinition, tickets [][]int) (result int) {
	for _, ticket := range tickets {
		result += CalculateTicketErrorRate(definitions, ticket)
	}
	return result
}

func CalculateTicketErrorRate(definitions []*FieldDefinition, ticket []int) (result int) {
	for _, field := range ticket {
		valid := false
		for _, definition := range definitions {
			if definition.WithinRange(field) {
				valid = true
			}
		}
		if !valid {
			result += field
		}
	}
	return result
}

func FilterValidTickets(definitions []*FieldDefinition, tickets [][]int) (result [][]int) {
	for _, ticket := range tickets {
		if ticketIsValid(definitions, ticket) {
			result = append(result, ticket)
		}
	}
	return result
}

func ticketIsValid(definitions []*FieldDefinition, ticket []int) bool {
	return CalculateTicketErrorRate(definitions, ticket) == 0
}

func IdentifyFieldPlacementCandidates(definitions []*FieldDefinition, tickets [][]int) (result map[string][]int) {
	result = make(map[string][]int)
	for _, definition := range definitions {
		for f := 0; f < len(tickets[0]); f++ {
			valid := 0
			for _, ticket := range tickets {
				if definition.WithinRange(ticket[f]) {
					valid++
				}
			}
			if valid == len(tickets) {
				result[definition.Name] = append(result[definition.Name], f)
			}
		}
	}
	return result
}

func FinalizeFieldPlacements(candidates map[string][]int) (result map[string]int) {
	result = make(map[string]int)

	for len(result) < len(candidates) {

		for n1, values := range candidates {
			if len(values) == 1 {
				result[n1] = values[0]

				for n2, others := range candidates {
					if n1 != n2 && len(others) > 1 && contains(others, values[0]) {
						candidates[n2] = remove(others, values[0])
					}
				}
			}
		}

	}
	return result
}

func contains(s []int, v int) bool {
	for _, ss := range s {
		if ss == v {
			return true
		}
	}
	return false
}

// super inefficient, but doesn't require any slice tricks ;)
func remove(s []int, v int) (result []int) {
	values := make(map[int]bool, len(s))
	for _, ss := range s {
		values[ss] = true
	}
	delete(values, v)
	result = make([]int, 0, len(values))
	for ss := range values {
		result = append(result, ss)
	}
	return result
}

func CalculateDepartureProduct(finalized map[string]int, ticket []int) (result int) {
	result = 1
	for name, field := range finalized {
		if strings.Contains(name, "departure") {
			result *= ticket[field]
		}
	}
	return result
}
