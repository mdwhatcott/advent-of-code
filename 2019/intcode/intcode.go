package intcode

func RunProgram(program []int, input func() int, output func(int)) []int {
	return NewIntCodeInterpreter(program, input, output).RunProgram()
}

const (
	_01_AddInstruction         = 1
	_02_MultiplyInstruction    = 2
	_03_InputInstruction       = 3
	_04_OutputInstruction      = 4
	_05_JumpIfTrueInstruction  = 5
	_06_JumpIfFalseInstruction = 6
	_07_LessThanInstruction    = 7
	_08_EqualsInstruction      = 8
	_99_ExitInstruction        = 99

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

type Interpreter struct {
	pointer  int
	program  []int
	inputs   func() int
	outputs  func(int)
	modes    []int
	finished bool
}

func NewIntCodeInterpreter(program []int, input func() int, output func(int)) *Interpreter {
	this := new(Interpreter)
	this.inputs = input
	this.outputs = output
	this.program = make([]int, len(program))
	copy(this.program, program)
	return this
}

func (this *Interpreter) RunProgram() []int {
	for !this.finished {
		this.processInstruction()
	}
	return this.program
}

func (this *Interpreter) processInstruction() {
	this.modes = splitDigits(this.value(this.pointer))

	//println(this.pointer, this.modes)

	switch opCode(this.modes) {

	case _01_AddInstruction:
		this.add()

	case _02_MultiplyInstruction:
		this.multiply()

	case _03_InputInstruction:
		this.input()

	case _04_OutputInstruction:
		this.output()

	case _05_JumpIfTrueInstruction:
		this.jumpIfTrue()

	case _06_JumpIfFalseInstruction:
		this.jumpIfFalse()

	case _07_LessThanInstruction:
		this.less()

	case _08_EqualsInstruction:
		this.equals()

	case _99_ExitInstruction:
		this.finished = true

	default:
		panic("not possible")
	}
}

func (this *Interpreter) access(slot int) int {
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

func (this *Interpreter) value(address int) int {
	return this.program[address]
}
func (this *Interpreter) setValue(address, value int) {
	this.program[address] = value
}
func (this *Interpreter) setReference(address, value int) {
	this.setValue(this.value(address), value)
}
func (this *Interpreter) reference(address int) int {
	return this.value(this.value(address))
}

func (this *Interpreter) add() {
	this.setReference(this.pointer+3, this.access(Slot1)+this.access(Slot2))
	this.pointer += 4
}
func (this *Interpreter) multiply() {
	this.setReference(this.pointer+3, this.access(Slot1)*this.access(Slot2))
	this.pointer += 4
}
func (this *Interpreter) input() {
	this.setReference(this.pointer+1, this.inputs())
	this.pointer += 2
}
func (this *Interpreter) output() {
	this.outputs(this.access(Slot1))
	this.pointer += 2
}
func (this *Interpreter) jumpIfTrue() {
	if this.access(Slot1) != 0 {
		this.pointer = this.access(Slot2)
	} else {
		this.pointer += 3
	}
}
func (this *Interpreter) jumpIfFalse() {
	if this.access(Slot1) == 0 {
		this.pointer = this.access(Slot2)
	} else {
		this.pointer += 3
	}
}
func (this *Interpreter) less() {
	if this.access(Slot1) < this.access(Slot2) {
		this.setReference(this.pointer+3, 1)
	} else {
		this.setReference(this.pointer+3, 0)
	}
	this.pointer += 4
}
func (this *Interpreter) equals() {
	if this.access(Slot1) == this.access(Slot2) {
		this.setReference(this.pointer+3, 1)
	} else {
		this.setReference(this.pointer+3, 0)
	}
	this.pointer += 4
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
