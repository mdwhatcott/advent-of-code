package advent

import "fmt"

func RunTweakedProgram(program []int, noun, verb int) {
	program[1] = noun
	program[2] = verb
	RunProgram(program)
}
func RunProgram(program []int) {
	NewIntCodeInterpreter(program).RunProgram()
}

const (
	AddInstruction      = 1
	MultiplyInstruction = 2
	ExitInstruction     = 99
)

type IntCodeInterpreter struct {
	program []int
	pointer int
}

func NewIntCodeInterpreter(program []int) *IntCodeInterpreter {
	return &IntCodeInterpreter{program: program, pointer: 0}
}

func (this *IntCodeInterpreter) RunProgram() {
	for {
		switch this.program[this.pointer] {
		case AddInstruction:
			this.pointer += this.add()
		case MultiplyInstruction:
			this.pointer += this.multiply()
		case ExitInstruction:
			return
		default:
			panic(fmt.Sprintf("unknown opcode: %d", this.program[this.pointer]))
		}
	}
}

func (this *IntCodeInterpreter) SetValue(address, value int) {
	this.program[address] = value
}
func (this *IntCodeInterpreter) setReference(address, value int) {
	this.SetValue(this.GetValue(address), value)
}
func (this *IntCodeInterpreter) GetValue(address int) int {
	return this.program[address]
}
func (this *IntCodeInterpreter) reference(address int) int {
	return this.GetValue(this.GetValue(address))
}

func (this *IntCodeInterpreter) add() int {
	this.setReference(this.pointer+3,
		this.reference(this.pointer+1)+
			this.reference(this.pointer+2))
	return 4
}
func (this *IntCodeInterpreter) multiply() int {
	this.setReference(this.pointer+3,
		this.reference(this.pointer+1)*
			this.reference(this.pointer+2))
	return 4
}
