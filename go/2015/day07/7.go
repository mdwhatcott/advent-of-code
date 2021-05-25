package advent

import (
	"fmt"
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
	numeric, err := strconv.ParseUint(wire, 10, 16)
	if err == nil {
		return uint16(numeric)
	}
	value, found := this.known[wire]
	if found {
		return value
	}
	this.known[wire] = this.solve(wire)
	return this.known[wire]
}
func (this *RecursiveCircuit) solve(wire string) uint16 {
	operation := this.instructions[wire]
	operator := decideOperator(operation)
	switch operator {
	case "SET":
		return this.SolveFor(operation[0])
	case "NOT":
		return ^this.SolveFor(operation[1])
	case "AND":
		return this.SolveFor(operation[0]) & this.SolveFor(operation[2])
	case "OR":
		return this.SolveFor(operation[0]) | this.SolveFor(operation[2])
	case "LSHIFT":
		return this.SolveFor(operation[0]) << this.SolveFor(operation[2])
	case "RSHIFT":
		return this.SolveFor(operation[0]) >> this.SolveFor(operation[2])
	}
	panic(fmt.Sprintln(wire, operation))
}

func decideOperator(operation []string) string {
	for _, word := range operation {
		if word == "AND" || word == "OR" || word == "LSHIFT" || word == "RSHIFT" || word == "NOT" {
			return word
		}
	}
	return "SET"
}
