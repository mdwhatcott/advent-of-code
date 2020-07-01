package advent

import "sync"

func part2(program []int, phases ...int) (result int) {
	waiter := new(sync.WaitGroup)

	ea := make(chan int, 2)
	ab := make(chan int, 1)
	bc := make(chan int, 1)
	cd := make(chan int, 1)
	de := make(chan int, 1)

	a := NewAmplifier(ea, ab, program, waiter)
	b := NewAmplifier(ab, bc, program, waiter)
	c := NewAmplifier(bc, cd, program, waiter)
	d := NewAmplifier(cd, de, program, waiter)
	e := NewAmplifier(de, ea, program, waiter)

	ea <- phases[0]
	ab <- phases[1]
	bc <- phases[2]
	cd <- phases[3]
	de <- phases[4]

	ea <- 0 // first input to amplifier a

	go e.Amplify()
	go d.Amplify()
	go c.Amplify()
	go b.Amplify()
	go a.Amplify()

	waiter.Wait()

	return <-ea // final output from amplifier e
}
