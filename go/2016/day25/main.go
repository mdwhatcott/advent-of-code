package main

import (
	"fmt"

	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"

	"advent/2016/util/assembunny"
	"advent/lib/util"
)

func main() {
	fmt.Print("Lowest starting value for 'a' that results in 0,1,0,1,0,1... ")
	fmt.Println(assert.So(part1(), should.Equal, 158).Fatal())
}

func part1() int {
	for x := 0; ; x++ {
		interpreter := assembunny.NewInterpreter(util.InputLines())
		interpreter.SetMaxOutputLength(100)
		interpreter.Set("a", x)
		interpreter.ExecuteProgram()
		if signal := interpreter.Out(); IsCorrect(signal) {
			return x
		}
	}
}
func IsCorrect(out []byte) bool {
	for x := range out {
		if x%2 == 0 && out[x] != 0 {
			return false
		} else if x%2 == 1 && out[x] != 1 {
			return false
		}
	}
	return true
}
