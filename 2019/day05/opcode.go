package advent

import "fmt"

func RunProgram(program []int, input func() int, output func(int)) {
	NewIntCodeInterpreter(program, input, output).RunProgram()
}

const (
	AddInstruction      = 1
	MultiplyInstruction = 2
	InputInstruction    = 3
	OutputInstruction   = 4
	ExitInstruction     = 99

	PositionMode  = 0
	ImmediateMode = 1
)

type IntCodeInterpreter struct {
	program []int
	pointer int
	inputs  func() int
	outputs func(int)
}

func NewIntCodeInterpreter(program []int, input func() int, output func(int)) *IntCodeInterpreter {
	return &IntCodeInterpreter{
		program: program,
		pointer: 0,
		inputs:  input,
		outputs: output,
	}
}

func (this *IntCodeInterpreter) RunProgram() {
	for {
		opCode := this.program[this.pointer]
		switch opCode {
		case AddInstruction:
			this.pointer += this.add()
		case MultiplyInstruction:
			this.pointer += this.multiply()
		case InputInstruction:
			this.pointer += this.input()
		case OutputInstruction:
			this.pointer += this.output()
		case ExitInstruction:
			return
		default:
			panic(fmt.Sprintf("unknown opcode: %d", opCode))
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

func (this *IntCodeInterpreter) input() int {
	this.setReference(this.pointer+1, this.inputs())
	return 2
}

func (this *IntCodeInterpreter) output() int {
	this.outputs(this.reference(this.pointer+1))
	return 2
}
