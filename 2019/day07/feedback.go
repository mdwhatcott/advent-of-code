package advent

import (
	"log"
	"sync"

	"advent/2019/intcode"
)

func part2(program []int, phases ...int) (result int) {
	var (
		ab = make(chan int)
		bc = make(chan int)
		cd = make(chan int)
		de = make(chan int)
		ea = make(chan int, 1)
	)
	var (
		a = NewAmplifier("a", ea, ab, program)
		b = NewAmplifier("b", ab, bc, program)
		c = NewAmplifier("c", bc, cd, program)
		d = NewAmplifier("d", cd, de, program)
		e = NewAmplifier("e", de, ea, program)
	)

	ea <- 0

	waiter := new(sync.WaitGroup)
	waiter.Add(5)

	go e.Amplify(phases[0], waiter)
	go d.Amplify(phases[1], waiter)
	go c.Amplify(phases[2], waiter)
	go b.Amplify(phases[3], waiter)
	go a.Amplify(phases[4], waiter)

	waiter.Wait()

	log.Println("ab:", len(ab))
	log.Println("bc:", len(bc))
	log.Println("cd:", len(cd))
	log.Println("de:", len(de))
	log.Println("ea:", len(ea))

	select {
	case result = <-ea:
		log.Println("All finished:", result)
		return result
	default:
		log.Println("No answer for phase combo:", phases)
		return 0
	}
}

type Amplifier struct {
	id      string
	in      chan int
	out     chan int
	program []int
}

func NewAmplifier(id string, in, out chan int, program []int) *Amplifier {
	return &Amplifier{
		id:      id,
		in:      in,
		out:     out,
		program: program,
	}
}

func (this *Amplifier) Amplify(phase int, waiter *sync.WaitGroup) {
	defer this.done(waiter)
	go func() { this.in <- phase }()
	intcode.NewIntCodeInterpreter(this.program, this.input, this.output).RunProgram()
}

func (this *Amplifier) input() (result int) {
	result = <-this.in
	log.Printf("%s: [%d] <- \n", this.id, result)
	return result
}

func (this *Amplifier) output(value int) {
	log.Printf("%s: -> [%d]\n", this.id, value)
	this.out <- value
}

func (this *Amplifier) done(waiter *sync.WaitGroup) {
	log.Printf("%s: finished!\n", this.id)
	waiter.Done()
}
