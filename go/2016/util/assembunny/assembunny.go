package assembunny

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Interpreter struct {
	registers map[string]int
	program   map[int][]string
	cursor    int
	counter   int
	out       []byte
	maxOut    int
}

func (this *Interpreter) SetMaxOutputLength(maxLength int) {
	this.maxOut = maxLength
}

func NewInterpreter(lines []string) *Interpreter {
	program := make(map[int][]string)
	for x, line := range lines {
		program[x] = strings.Fields(line)
	}
	return &Interpreter{
		program:   program,
		registers: make(map[string]int),
		out:       []byte{},
	}
}

func (this *Interpreter) Get(register string) int {
	return this.registers[register]
}

func (this *Interpreter) Set(register string, value int) {
	this.registers[register] = value
}

func (this *Interpreter) Out() []byte {
	return this.out
}

func (this *Interpreter) ExecuteProgram() {
	for this.cursor = 0; this.cursor < len(this.program); this.log() {
		this.cursor += this.execute(this.program[this.cursor])
	}
}

func (this *Interpreter) execute(args []string) (cursorOffset int) {
	switch instruction := args[0]; instruction {
	case OUT:
		return this.executeOut(args)
	case COPY:
		this.executeCopy(args)
	case INCREMENT:
		this.executeIncrement(args)
	case DECREMENT:
		this.executeDecrement(args)
	case TOGGLE:
		this.toggleInstruction(args)
	case JUMP:
		return this.calculateJumpCursorOffset(args)
	default:
		log.Println("Ignoring instruction:", args)
	}
	return nextInstructionCursorOffset
}
func (this *Interpreter) executeOut(args []string) int {
	this.out = append(this.out, byte(this.evaluate(args[1])))
	if this.maxOut == 0 || len(this.out) < this.maxOut {
		return nextInstructionCursorOffset
	} else {
		return len(this.program) - this.cursor + 1
	}
}

func (this *Interpreter) log() {
	if this.counter++; this.counter%10000000 == 0 {
		log.Printf("Counter: %d Cursor: %d Registers: %v", this.counter, this.cursor, this.registers)
	}
}
func (this *Interpreter) executeCopy(args []string) {
	register := args[2]
	value := this.evaluate(args[1])
	this.Set(register, value)
}
func (this *Interpreter) executeIncrement(args []string) {
	register := args[1]
	this.increment(register)
}
func (this *Interpreter) executeDecrement(args []string) {
	register := args[1]
	this.decrement(register)
}
func (this *Interpreter) toggleInstruction(args []string) {
	offset := this.evaluate(args[1])
	if index := this.cursor + offset; this.instructionWithinBounds(index) {
		this.toggle(index)
	}
}
func (this *Interpreter) instructionWithinBounds(instructionIndex int) bool {
	return instructionIndex >= 0 && instructionIndex < len(this.program)
}
func (this *Interpreter) toggle(instructionIndex int) {
	old := this.program[instructionIndex]

	switch instruction := old[0]; instruction {
	case INCREMENT:
		this.program[instructionIndex] = []string{DECREMENT, old[1]}
	case DECREMENT:
		this.program[instructionIndex] = []string{INCREMENT, old[1]}
	case TOGGLE:
		this.program[instructionIndex] = []string{INCREMENT, old[1]}
	case JUMP:
		this.program[instructionIndex] = []string{COPY, old[1], old[2]}
	case COPY:
		this.program[instructionIndex] = []string{JUMP, old[1], old[2]}
	}
}
func (this *Interpreter) calculateJumpCursorOffset(args []string) int {
	if canJump := this.evaluate(args[1]) != 0; canJump {
		offset := args[2]
		return this.evaluate(offset)
	} else {
		return nextInstructionCursorOffset
	}
}

func (this *Interpreter) evaluate(value string) int {
	if n, err := strconv.Atoi(value); err == nil {
		return n
	} else {
		return this.registers[value]
	}
}

func (this *Interpreter) increment(register string) {
	this.registers[register]++
}
func (this *Interpreter) decrement(register string) {
	this.registers[register]--
}

func (this *Interpreter) String() string {
	buffer := new(bytes.Buffer)
	fmt.Fprintf(buffer, "Counter: %d Cursor: %d Registers: %v\n", this.counter, this.cursor, this.registers)

	for x := 0; x < len(this.program); x++ {
		if x == this.cursor {
			fmt.Fprint(buffer, "> ")
		} else {
			fmt.Fprint(buffer, "  ")
		}
		fmt.Fprintf(buffer, "%s\n", this.program[x])
	}

	fmt.Fprintln(buffer)
	fmt.Fprintln(buffer, this.registers)
	fmt.Fprintln(buffer, "---------------------------")
	return buffer.String()
}

const (
	COPY      = "cpy"
	INCREMENT = "inc"
	DECREMENT = "dec"
	JUMP      = "jnz"
	TOGGLE    = "tgl"
	OUT       = "out"

	nextInstructionCursorOffset = 1
)
