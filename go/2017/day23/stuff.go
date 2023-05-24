package day23

import (
	"log"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
)

type Interpreter struct {
	registers map[string]int
	program   []string
	cursor    int
	mul       int
	pc        int
}

func NewInterpreter(program []string) *Interpreter {
	return &Interpreter{
		program:   program,
		registers: make(map[string]int),
	}
}

func (this *Interpreter) Run() {
	for this.cursorInProgram() {
		this.cursor = this.execute(this.program[this.cursor])
		this.pc++
	}
}

func (this *Interpreter) cursorInProgram() bool {
	return this.cursor >= 0 && this.cursor < len(this.program)
}

func (this *Interpreter) resolve(label string) int {
	value, found := this.registers[label]
	if found {
		return value
	}
	return parse.Int(label)
}
func (this *Interpreter) execute(instruction string) int {
	fields := strings.Fields(instruction)
	//defer this.Log(instruction)
	switch fields[0] {
	case "set":
		this.registers[fields[1]] = this.resolve(fields[2])
	case "sub":
		this.registers[fields[1]] -= this.resolve(fields[2])
	case "mul":
		this.mul++
		this.registers[fields[1]] *= this.resolve(fields[2])
	case "jnz":
		if this.resolve(fields[1]) != 0 {
			return this.cursor + this.resolve(fields[2])
		}
	}
	return this.cursor + 1
}

func (this *Interpreter) Log(instruction string) {
	log.Printf("(%d) %d %-15s a:%-10d b:%-10d c:%-10d d:%-10d e:%-10d f:%-10d g:%-10d h:%-10d",
		this.pc, this.cursor, instruction,
		this.registers["a"], this.registers["b"],
		this.registers["c"], this.registers["d"],
		this.registers["e"], this.registers["f"],
		this.registers["g"], this.registers["h"],
	)
}
