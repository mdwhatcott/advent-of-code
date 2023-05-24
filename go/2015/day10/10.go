package main

import (
	"bytes"
	"fmt"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func main() {
	input := util.InputString()
	forty := LookSayMany(input, 40)
	fifty := LookSayMany(forty, 10)
	fmt.Println("Length of result after 40 rounds:", len(forty))
	fmt.Println("Length of result after 50 rounds:", len(fifty))
}

func LookSayMany(value string, times int) string {
	for x := 0; x < times; x++ {
		value = LookSay(value)
	}
	return value
}

func LookSay(look string) string {
	say := new(bytes.Buffer)
	var block int
	var current rune
	for _, c := range look {
		if block == 0 {
			current = c
			block++
		} else if c != current {
			fmt.Fprint(say, block)
			say.WriteRune(current)
			block = 1
			current = c
		} else {
			block++
		}
	}
	if block > 0 {
		fmt.Fprint(say, block)
		say.WriteRune(current)
	}
	return say.String()
}
