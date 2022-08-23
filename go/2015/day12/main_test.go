package main

import (
	"log"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test(t *testing.T) {
	input := util.InputString()
	assert.Error(t).So(part1(input), should.Equal, 156366)
	assert.Error(t).So(part2(input), should.Equal, 96852)
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
