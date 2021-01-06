package advent

import (
	"strconv"
	"strings"
)

type RecursiveCircuit struct {
	instructions map[string][]string
	known        map[string]uint16
}

func NewRecursiveCircuit(instructions []string) *RecursiveCircuit {
	circuit := &RecursiveCircuit{
		instructions: make(map[string][]string),
		known:        make(map[string]uint16),
	}
	for _, line := range instructions {
		words := strings.Fields(line)
		destination := words[len(words)-1]
		operation := words[0 : len(words)-2]
		circuit.instructions[destination] = operation
	}
	return circuit
}

func (this *RecursiveCircuit) SolveFor(wire string) uint16 {
	if value, err := strconv.ParseUint(wire, 10, 16); err == nil {
		return uint16(value)
	}
	if value, found := this.known[wire]; found {
		return value
	}
	switch operation := this.instructions[wire]; decideOperator(operation) {
	case "SET":
		this.known[wire] = this.SolveFor(operation[0])
	case "NOT":
		this.known[wire] = ^this.SolveFor(operation[1])
	case "AND":
		this.known[wire] = this.SolveFor(operation[0]) & this.SolveFor(operation[2])
	case "OR":
		this.known[wire] = this.SolveFor(operation[0]) | this.SolveFor(operation[2])
	case "LSHIFT":
		this.known[wire] = this.SolveFor(operation[0]) << this.SolveFor(operation[2])
	case "RSHIFT":
		this.known[wire] = this.SolveFor(operation[0]) >> this.SolveFor(operation[2])
	}
	return this.known[wire]
}

func decideOperator(operation []string) string {
	for _, word := range operation {
		if word == "AND" || word == "OR" || word == "LSHIFT" || word == "RSHIFT" || word == "NOT" {
			return word
		}
	}
	return "SET"
}
