package main

import (
	"log"
	"strings"

	"advent/lib/util"
)

type Interpreter struct {
	registers map[string]int
	program   map[int][]string
	cursor    int
	counter   int
}

func NewInterpreter(lines []string) *Interpreter {
	program := make(map[int][]string)
	for x, line := range lines {
		line = strings.Replace(line, ",", "", 1)
		program[x] = strings.Fields(line)
	}
	return &Interpreter{
		program:   program,
		registers: make(map[string]int),
	}
}

func (this *Interpreter) ExecuteProgram() {
	for this.cursor = 0; this.cursor < len(this.program); this.log() {
		this.cursor += this.execute(this.program[this.cursor])
	}
}

func (this *Interpreter) execute(args []string) (cursorOffset int) {
	switch instruction := args[0]; instruction {
	case HALF:
		this.registers[args[1]] /= 2
	case TRIPLE:
		this.registers[args[1]] *= 3
	case INCREMENT:
		this.registers[args[1]]++
	case JUMP:
		return util.ParseInt(args[1])
	case JUMPEVEN:
		if this.registers[args[1]]%2 == 0 {
			return util.ParseInt(args[2])
		}
	case JUMP1:
		if this.registers[args[1]] == 1 {
			return util.ParseInt(args[2])
		}
	default:
		log.Println("Ignoring instruction:", args)
	}
	return nextInstructionCursorOffset
}

func (this *Interpreter) log() {
	if this.counter++; this.counter%10000000 == 0 {
		log.Printf("Counter: %d Cursor: %d Registers: %v", this.counter, this.cursor, this.registers)
	}
}

const (
	HALF      = "hlf"
	TRIPLE    = "tpl"
	INCREMENT = "inc"
	JUMP      = "jmp"
	JUMPEVEN  = "jie"
	JUMP1     = "jio"

	nextInstructionCursorOffset = 1
)
