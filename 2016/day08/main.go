package main

import (
	"flag"

	"advent/2016/util/lcd"
	"advent/lib/util"
)

func main() {
	animate := flag.Bool("animate", false, "Animate pixels being drawn.")
	flag.Parse()
	lcd.Display(util.InputScanner().Scanner, *animate)
}
