package main

import (
	"flag"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
	"github.com/mdwhatcott/advent-of-code-lcd"
)

func main() {
	animate := flag.Bool("animate", false, "Animate pixels being drawn.")
	flag.Parse()
	lcd.Display(util.InputScanner().Scanner, *animate)
}
