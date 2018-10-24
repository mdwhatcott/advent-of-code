package day18

import "advent/lib/util"

func Part1() int {
	interpreter := NewSoundBlaster(util.InputLines())
	interpreter.Run()
	return interpreter.recovered
}

func Part2() int {
	var (
		a, b = make(chan int, 100), make(chan int, 100)
		zero = NewAgent(0, a, b, util.InputLines())
		one  = NewAgent(1, b, a, util.InputLines())
	)
	go zero.Run()
	one.Run()
	return one.sendCount
}
