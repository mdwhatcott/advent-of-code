package intcode

type Harness struct {
	original []int
	final    []int

	inputs  []int
	outputs []int

	in int
}

func NewHarness(program []int, inputs ...int) *Harness {
	return &Harness{original: program, inputs: inputs}
}

func (this *Harness) increment() {
	this.in++
}
func (this *Harness) input() int {
	defer this.increment()
	return this.inputs[this.in%len(this.inputs)]
}
func (this *Harness) output(i int) {
	this.outputs = append(this.outputs, i)
}

func (this *Harness) Run() {
	this.final = NewIntCodeInterpreter(this.original, this.input, this.output).RunProgram()
}
func (this *Harness) Outputs() []int {
	return this.outputs
}
func (this *Harness) FinalProgram() []int {
	return this.final
}
