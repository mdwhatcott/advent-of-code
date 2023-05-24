package day18

import (
	"strings"
	"time"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
)

type Agent struct {
	id        int
	input     chan int
	output    chan int
	registers map[string]int
	program   []string
	cursor    int
	sendCount int
}

func NewAgent(id int, input, output chan int, program []string) *Agent {
	return &Agent{
		id:        id,
		program:   program,
		input:     input,
		output:    output,
		registers: map[string]int{"p": id},
	}
}

func (this *Agent) Run() {
	for this.cursorInProgram() {
		this.cursor = this.execute(this.program[this.cursor])
	}
}

func (this *Agent) cursorInProgram() bool {
	return this.cursor >= 0 && this.cursor < len(this.program)
}

func (this *Agent) resolve(label string) int {
	value, found := this.registers[label]
	if found {
		return value
	}
	return parse.Int(label)
}
func (this *Agent) execute(instruction string) int {
	fields := strings.Fields(instruction)
	switch fields[0] {
	case "set":
		this.registers[fields[1]] = this.resolve(fields[2])
	case "add":
		this.registers[fields[1]] += this.resolve(fields[2])
	case "mul":
		this.registers[fields[1]] *= this.resolve(fields[2])
	case "mod":
		this.registers[fields[1]] %= this.resolve(fields[2])
	case "snd":
		select {
		case this.output <- this.resolve(fields[1]):
			this.sendCount++
		case <-time.After(time.Second):
			return -1
		}
	case "rcv":
		select {
		case value := <-this.input:
			this.registers[fields[1]] = value
		case <-time.After(time.Second):
			return -1
		}
	case "jgz":
		if this.resolve(fields[1]) > 0 {
			return this.cursor + this.resolve(fields[2])
		}
	}
	return this.cursor + 1
}
