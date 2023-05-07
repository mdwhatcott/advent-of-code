package day12

import (
	"log"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test(t *testing.T) {
	input := util.InputString()
	should.So(t, part1(input), should.Equal, 156366)
	should.So(t, part2(input), should.Equal, 96852)
}

func part1(input string) (sum int) {
	scanner := prepare(input)

	for scanner.Scan() {
		sum += parse.Int(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func part2(input string) int {
	scanner := prepare(input)
	scope := new(Scope)

	for scanner.Scan() {
		scope = scope.Receive(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return scope.Sum()
}
