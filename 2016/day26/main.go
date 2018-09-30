package main

import (
	"bufio"
	"bytes"

	"advent/2016/util/assembunny"
	"advent/2016/util/lcd"
	"advent/lib/util"
)

func main() {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.ExecuteProgram()
	lcdInstructions := interpreter.Out()
	scanner := bufio.NewScanner(bytes.NewReader(lcdInstructions))
	lcd.Display(scanner, true)
}
