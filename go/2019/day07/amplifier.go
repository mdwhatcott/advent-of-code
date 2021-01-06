package advent

import (
	"sync"

	"advent/2019/intcode"
)

type Amplifier struct {
	in      chan int
	out     chan int
	machine *intcode.Interpreter
	waiter  *sync.WaitGroup
}

func NewAmplifier(in, out chan int, program []int, waiter *sync.WaitGroup) *Amplifier {
	this := new(Amplifier)
	this.in = in
	this.out = out
	this.waiter = waiter
	this.waiter.Add(1)
	this.machine = intcode.NewIntCodeInterpreter(program, this.input, this.output)
	return this
}

func (this *Amplifier) Amplify() {
	this.machine.RunProgram()
	this.waiter.Done()
}

func (this *Amplifier) input() (result int) { return <-this.in }
func (this *Amplifier) output(value int)    { this.out <- value }
