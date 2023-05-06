package advent

import (
	"bufio"
	"log"
	"strings"

	"advent/lib/parse"
	"advent/lib/util"
)

func Part1() interface{} {
	return NewInterpreter(util.InputScanner().Scanner).Accumulate()
}

func Part2() interface{} {
	return NewInterpreter(util.InputScanner().Scanner).AccumulateByFixingNonTerminatingProgram()
}

type Interpreter struct {
	program     []Statement
	cursor      int
	accumulator int
	executed    map[int]bool
}

func NewInterpreter(scanner *bufio.Scanner) *Interpreter {
	this := &Interpreter{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		words := strings.Fields(line)
		this.program = append(this.program, Statement{
			Operation: words[0],
			Value:     parse.Int(words[1]),
		})
	}
	return this
}

func (this *Interpreter) Accumulate() int {
	this.cursor = 0
	this.accumulator = 0
	this.executed = make(map[int]bool)

	for {
		if this.cursor >= len(this.program) {
			return this.accumulator
		}

		if this.executed[this.cursor] {
			return this.accumulator
		}

		this.executed[this.cursor] = true

		statement := this.program[this.cursor]
		switch statement.Operation {
		case "nop":
			this.cursor++
		case "acc":
			this.accumulator += statement.Value
			this.cursor++
		case "jmp":
			this.cursor += statement.Value
		default:
			log.Panicln("BAD INSTRUCTION:", this.cursor, statement)
		}
	}
}

func (this *Interpreter) AccumulateByFixingNonTerminatingProgram() (result int) {
	for x := 0; x < len(this.program); x++ {
		if this.program[x].Operation == "acc" {
			continue
		}

		this.program[x] = this.program[x].Swap()

		result = this.Accumulate()
		if this.terminatedCleanly() {
			return result
		}

		this.program[x] = this.program[x].Swap()

	}
	panic("failed to terminate cleanly in any attempted configuration")
}

func (this *Interpreter) terminatedCleanly() bool {
	return this.cursor == len(this.program)
}

type Statement struct {
	Operation string
	Value     int
}

func (this Statement) Swap() Statement {
	return Statement{
		Operation: swap[this.Operation],
		Value:     this.Value,
	}
}

var swap = map[string]string{
	"jmp": "nop",
	"nop": "jmp",
}
