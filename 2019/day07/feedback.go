package advent

import (
	"log"
	"sync"

	"advent/2019/intcode"
)

func part2(program []int, phases ...int) (result int) {
	var (
		waiter = new(sync.WaitGroup)

		ea = make(chan int, 2)
		ab = make(chan int, 1)
		bc = make(chan int, 1)
		cd = make(chan int, 1)
		de = make(chan int, 1)
	)

	var (
		a = NewAmplifier("a", ea, ab, program, waiter)
		b = NewAmplifier("b", ab, bc, program, waiter)
		c = NewAmplifier("c", bc, cd, program, waiter)
		d = NewAmplifier("d", cd, de, program, waiter)
		e = NewAmplifier("e", de, ea, program, waiter)
	)

	ea <- phases[0]
	ab <- phases[1]
	bc <- phases[2]
	cd <- phases[3]
	de <- phases[4]

	ea <- 0

	go e.Amplify()
	go d.Amplify()
	go c.Amplify()
	go b.Amplify()
	go a.Amplify()

	waiter.Wait()

	log.Println("ab:", len(ab), "bc:", len(bc), "cd:", len(cd), "de:", len(de), "ea:", len(ea))

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
	id          string
	in          chan int
	out         chan int
	interpreter *intcode.Interpreter
	waiter      *sync.WaitGroup
}

func NewAmplifier(id string, in, out chan int, program []int, waiter *sync.WaitGroup) *Amplifier {
	this := Amplifier{id: id, in: in, out: out, waiter: waiter}
	this.waiter.Add(1)
	this.interpreter = intcode.NewIntCodeInterpreter(program, this.input, this.output)
	return &this
}

func (this *Amplifier) Amplify() {
	this.interpreter.RunProgram()
	log.Printf("%s: finished!\n", this.id)
	this.waiter.Done()
}

func (this *Amplifier) input() (result int) {
	result = <-this.in
	log.Printf("%s: <- [%d] \n", this.id, result)
	return result
}

func (this *Amplifier) output(value int) {
	log.Printf("%s: [%d] ->\n", this.id, value)
	this.out <- value
}
