package advent

import (
	"fmt"
)

func RunProgram(program []int, input func() int, output func(int)) {
	NewIntCodeInterpreter(program, input, output).RunProgram()
}

const (
	AddInstruction         = 1
	MultiplyInstruction    = 2
	InputInstruction       = 3
	OutputInstruction      = 4
	JumpIfTrueInstruction  = 5
	JumpIfFalseInstruction = 6
	LessThanInstruction    = 7
	EqualsInstruction      = 8
	ExitInstruction        = 99

	PositionMode  = 0
	ImmediateMode = 1

	Slot1 = 2
	Slot2 = 1
	Slot3 = 0
)

var offsets = map[int]int{
	Slot1: 1,
	Slot2: 2,
	Slot3: 3,
}

type IntCodeInterpreter struct {
	pointer int
	program []int
	inputs  func() int
	outputs func(int)
	modes   []int
}

func NewIntCodeInterpreter(program []int, input func() int, output func(int)) *IntCodeInterpreter {
	return &IntCodeInterpreter{
		pointer: 0,
		program: program,
		inputs:  input,
		outputs: output,
	}
}

func (this *IntCodeInterpreter) RunProgram() {
	for {
		this.modes = splitDigits(this.value(this.pointer))
		//log.Println(this.pointer, this.modes)

		switch opCode(this.modes) {

		case AddInstruction:
			this.pointer += this.add()

		case MultiplyInstruction:
			this.pointer += this.multiply()

		case InputInstruction:
			this.pointer += this.input()

		case OutputInstruction:
			this.pointer += this.output()

		case JumpIfTrueInstruction:
			this.pointer = this.jumpIfTrue()

		case JumpIfFalseInstruction:
			this.pointer = this.jumpIfFalse()

		case LessThanInstruction:
			this.pointer += this.less()

		case EqualsInstruction:
			this.pointer += this.equals()

		case ExitInstruction:
			return

		default:
			panic(fmt.Sprintf("unknown instruction: %d", this.modes))
		}
	}
}

func (this *IntCodeInterpreter) access(slot int) int {
	address := this.pointer + offsets[slot]
	switch this.modes[slot] {

	case PositionMode:
		return this.reference(address)

	case ImmediateMode:
		return this.value(address)

	default:
		panic("not possible")
	}
}

func (this *IntCodeInterpreter) value(address int) int {
	return this.program[address]
}
func (this *IntCodeInterpreter) setValue(address, value int) {
	this.program[address] = value
}
func (this *IntCodeInterpreter) setReference(address, value int) {
	this.setValue(this.value(address), value)
}
func (this *IntCodeInterpreter) reference(address int) int {
	return this.value(this.value(address))
}

func (this *IntCodeInterpreter) add() int {
	this.setReference(this.pointer+3, this.access(Slot1)+this.access(Slot2))
	return 4
}
func (this *IntCodeInterpreter) multiply() int {
	this.setReference(this.pointer+3, this.access(Slot1)*this.access(Slot2))
	return 4
}
func (this *IntCodeInterpreter) input() int {
	this.setReference(this.pointer+1, this.inputs())
	return 2
}
func (this *IntCodeInterpreter) output() int {
	this.outputs(this.access(Slot1))
	return 2
}
func (this *IntCodeInterpreter) jumpIfTrue() int {
	if this.access(Slot1) != 0 {
		return this.access(Slot2)
	}
	return this.pointer + 3
}
func (this *IntCodeInterpreter) jumpIfFalse() int {
	if this.access(Slot1) == 0 {
		return this.access(Slot2)
	}
	return this.pointer + 3
}
func (this *IntCodeInterpreter) less() int {
	if this.access(Slot1) < this.access(Slot2) {
		this.setReference(this.pointer+3, 1)
	} else {
		this.setReference(this.pointer+3, 0)
	}
	return 4
}
func (this *IntCodeInterpreter) equals() int {
	if this.access(Slot1) == this.access(Slot2) {
		this.setReference(this.pointer+3, 1)
	} else {
		this.setReference(this.pointer+3, 0)
	}
	return 4
}

func opCode(digits []int) int {
	if len(digits) == 1 {
		return digits[0]
	}
	return digits[len(digits)-2]*10 + digits[len(digits)-1]
}

func splitDigits(value int) (digits []int) {
	magnitude := 10
	for value > 0 {
		remainder := value % magnitude
		digit := remainder / (magnitude / 10)
		digits = append([]int{digit}, digits...)
		value -= remainder
		magnitude *= 10
	}
	for len(digits) < 5 {
		digits = append([]int{0}, digits...)
	}
	return digits
}
