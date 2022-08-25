package main

import (
	"bufio"
	"strings"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func main() {
	people := ParseRelations(util.InputScanner().Scanner)
	assert.So(nil, ComputeHappiestArrangement(people...), should.Equal, 733)

	people = append(people, NewPerson("Me"))
	assert.So(nil, ComputeHappiestArrangement(people...), should.Equal, 725)
}

func ParseRelations(scanner *bufio.Scanner) (people []*Person) {
	names := make(map[string]*Person)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line) // "Alice would gain 54 happiness units by sitting next to Bob."
		if line == "" {
			continue
		}
		line = strings.Replace(line, ".", "", 1)
		words := strings.Fields(line)
		from := words[0]
		to := words[10]
		value := util.ParseInt(words[3])
		if words[2] == "lose" {
			value = -value
		}

		a := names[from]
		if a == nil {
			a = NewPerson(from)
			names[from] = a
		}

		b := names[to]
		if b == nil {
			b = NewPerson(to)
			names[to] = b
		}

		a.Relations[b] = value
	}

	for _, person := range names {
		people = append(people, person)
	}
	return people
}
