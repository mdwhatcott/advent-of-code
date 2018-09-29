package main

import (
	"flag"

	"github.com/mdwhatcott/advent-of-code-2016/util/lcd"
	"advent/lib/util"
)

func main() {
	animate := flag.Bool("animate", false, "Animate pixels being drawn.")
	flag.Parse()
	lcd.Display(util.InputScanner(), *animate)
}
