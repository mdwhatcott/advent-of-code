package main

import (
	"fmt"
	"log"

	"advent/lib/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	input := util.InputString()
	fmt.Println(assert.So(part1(input), should.Equal, 156366))
	fmt.Println(assert.So(part2(input), should.Equal, 96852))
}

func part1(input string) (sum int) {
	scanner := prepare(input)

	for scanner.Scan() {
		sum += util.ParseInt(scanner.Text())
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
