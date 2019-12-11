package advent

import "fmt"

func RunProgram(program []int) {
	NewIntCodeInterpreter(program, nil).RunProgram()
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
	input   func() int
	output  chan int
}

func NewIntCodeInterpreter(program []int, input func() int) *IntCodeInterpreter {
	return &IntCodeInterpreter{
		program: program,
		pointer: 0,
		input:   input,
		output:  make(chan int),
	}
}

type OpCode struct {
	OpCode     int
	Parameters []Parameter
}

type Parameter struct {
	Mode  int
	Raw   int
	Value int
}

func ParseOpCode(program []int, index int) OpCode {
	return OpCode{}
}

func (this *IntCodeInterpreter) RunProgram() {
	defer close(this.output)

	for {
		opCode := this.program[this.pointer]
		switch opCode {
		case AddInstruction:
			this.pointer += this.add()
		case MultiplyInstruction:
			this.pointer += this.multiply()
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
