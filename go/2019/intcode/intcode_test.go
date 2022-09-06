package intcode

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestOpCodeFixture(t *testing.T) {
	should.Run(&OpCodeFixture{T: should.New(t)}, should.Options.UnitTests())
}

type OpCodeFixture struct {
	*should.T
	i       int
	inputs  []int
	outputs []int
}

func (this *OpCodeFixture) input() int {
	this.i++
	return this.inputs[(this.i-1)%len(this.inputs)]
}
func (this *OpCodeFixture) output(value int) {
	this.outputs = append(this.outputs, value)
}
func (this *OpCodeFixture) run(program []int) []int {
	return RunProgram(program, this.input, this.output)
}

func (this *OpCodeFixture) TestRunProgramA() {
	program := []int{
		1, 9, 10, 3,
		2, 3, 11, 0,
		99,
		30, 40, 50,
	}
	program = this.run(program)
	this.So(program, should.Equal, []int{
		3500, 9, 10, 70,
		2, 3, 11, 0,
		99,
		30, 40, 50,
	})
}
func (this *OpCodeFixture) TestRunProgramB() {
	program := []int{1, 0, 0, 0, 99}
	program = this.run(program)
	this.So(program, should.Equal, []int{2, 0, 0, 0, 99})
}
func (this *OpCodeFixture) TestRunProgramC() {
	program := []int{2, 3, 0, 3, 99}
	program = this.run(program)
	this.So(program, should.Equal, []int{2, 3, 0, 6, 99})
}
func (this *OpCodeFixture) TestRunProgramD() {
	program := []int{2, 4, 4, 5, 99, 0}
	program = this.run(program)
	this.So(program, should.Equal, []int{2, 4, 4, 5, 99, 9801})
}
func (this *OpCodeFixture) TestRunProgramE() {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	program = this.run(program)
	this.So(program, should.Equal, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
func (this *OpCodeFixture) TestRunProgramInputOutput() {
	this.inputs = append(this.inputs, 42)
	program := []int{3, 0, 4, 0, 99}
	program = this.run(program)
	this.So(program, should.Equal, []int{42, 0, 4, 0, 99})
	this.So(this.outputs, should.Equal, []int{42})
}
func (this *OpCodeFixture) TestRunProgramF() {
	program := []int{1002, 4, 3, 4, 33}
	program = this.run(program)
	this.So(program, should.Equal, []int{1002, 4, 3, 4, 99})
}
func (this *OpCodeFixture) TestRunProgramG() {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	this.inputs = append(this.inputs, 8)
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{1})
}
func (this *OpCodeFixture) TestRunProgramG2() {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	this.inputs = append(this.inputs, 8+1)
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{0})
}
func (this *OpCodeFixture) TestRunProgramH() {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	this.inputs = append(this.inputs, 8)
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{1})
}
func (this *OpCodeFixture) TestRunProgramH2() {
	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	this.inputs = append(this.inputs, 8+1)
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{0})
}
func (this *OpCodeFixture) TestRunProgramI() {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	this.inputs = append(this.inputs, 0)
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{0})
}
func (this *OpCodeFixture) TestRunProgramI1() {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	this.inputs = append(this.inputs, 1)
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{1})
}
func (this *OpCodeFixture) TestFinalProgramA() {
	program := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	this.inputs = []int{8 - 1}
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{999})
}
func (this *OpCodeFixture) TestFinalProgramB() {
	program := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	this.inputs = []int{8 - 0}
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{1000})
}
func (this *OpCodeFixture) TestFinalProgramC_Below() {
	program := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	this.inputs = []int{8 - 1}
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{999})
}
func (this *OpCodeFixture) TestFinalProgramC_At() {
	program := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	this.inputs = []int{8 - 0}
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{1000})
}
func (this *OpCodeFixture) TestFinalProgramC_Above() {
	program := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	this.inputs = []int{8 + 1}
	program = this.run(program)
	this.So(this.outputs, should.Equal, []int{1001})
}
func (this *OpCodeFixture) TestDay9Part1ExampleA() {
	original := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	_ = this.run(original)
	this.So(this.outputs, should.Equal, original)
}
func (this *OpCodeFixture) TestDay9Part1ExampleB() {
	original := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	_ = this.run(original)
	output := fmt.Sprint(this.outputs[0])
	this.So(output, should.HaveLength, 16)
}
func (this *OpCodeFixture) TestDay9Part1ExampleC() {
	original := []int{104, 1125899906842624, 99}
	_ = this.run(original)
	this.So(this.outputs, should.Equal, []int{original[1]})
}

func (this *OpCodeFixture) TestParsing() {
	this.So(splitDigits(5), should.Equal, []int{0, 0, 0, 0, 5})
	this.So(splitDigits(45), should.Equal, []int{0, 0, 0, 4, 5})
	this.So(splitDigits(345), should.Equal, []int{0, 0, 3, 4, 5})
	this.So(splitDigits(2345), should.Equal, []int{0, 2, 3, 4, 5})
	this.So(splitDigits(12345), should.Equal, []int{1, 2, 3, 4, 5})

	this.So(opCode(splitDigits(5)), should.Equal, 5)
	this.So(opCode(splitDigits(45)), should.Equal, 45)
	this.So(opCode(splitDigits(345)), should.Equal, 45)
	this.So(opCode(splitDigits(2345)), should.Equal, 45)
	this.So(opCode(splitDigits(12345)), should.Equal, 45)

	digits := splitDigits(109)
	this.So(digits, should.Equal, []int{0, 0, 1, 0, 9})
	this.So(opCode(digits), should.Equal, 9)
}
