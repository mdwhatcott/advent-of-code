package advent

import "advent/2019/intcode"

type IO struct {
	inputs  []int
	outputs []int

	i     int
	phase bool
}

func NewIO(inputs ...int) *IO {
	return &IO{inputs: inputs, phase: true}
}

func (this *IO) Run(program ...int) int {
	this.outputs = append(this.outputs, 0)

	for this.i < len(this.inputs) {
		intcode.NewIntCodeInterpreter(program, this.input, this.output).RunProgram()
	}
	return this.outputs[len(this.outputs)-1]
}

func (this *IO) input() (result int) {
	defer this.increment()
	if this.phase {
		result = this.inputs[this.i]
	} else {
		result = this.outputs[len(this.outputs)-1]
	}
	return result
}

func (this *IO) increment() {
	if this.phase {
		this.i++
	}
	this.phase = !this.phase
}

func (this *IO) output(output int) {
	this.outputs = append(this.outputs, output)
}
