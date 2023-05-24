package main

import (
	"bufio"
	"bytes"

	"github.com/mdwhatcott/advent-of-code-assembunny"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
	"github.com/mdwhatcott/advent-of-code-lcd"
)

func main() {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.ExecuteProgram()
	lcdInstructions := interpreter.Out()
	scanner := bufio.NewScanner(bytes.NewReader(lcdInstructions))
	lcd.Display(scanner, true)
}
