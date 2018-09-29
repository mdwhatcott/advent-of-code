package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestRoundTableFixture(t *testing.T) {
	gunit.Run(new(RoundTableFixture), t)
}

type RoundTableFixture struct {
	*gunit.Fixture
}

func (this *RoundTableFixture) TestManualRegistrationOfRelations() {
	alice := NewPerson("Alice")
	bob := NewPerson("Bob")
	carol := NewPerson("Carol")
	david := NewPerson("David")

	alice.Relations[bob] = 54
	alice.Relations[carol] = -79
	alice.Relations[david] = -2

	bob.Relations[alice] = 83
	bob.Relations[carol] = -7
	bob.Relations[david] = -63

	carol.Relations[alice] = -62
	carol.Relations[bob] = 60
	carol.Relations[david] = 55

	david.Relations[alice] = 46
	david.Relations[bob] = -7
	david.Relations[carol] = 41

	this.So(ComputeHappiestArrangement(bob, david, carol, alice), should.Equal, 330)
	this.So(ComputeHappiestArrangement(bob, carol, david, alice), should.Equal, 330)
	this.So(ComputeHappiestArrangement(alice, bob, carol, david), should.Equal, 330)
}

func (this *RoundTableFixture) TestAutomatedRegistrationOfRelations() {
	input := bufio.NewScanner(strings.NewReader(exampleInput))
	people := ParseRelations(input)
	this.So(ComputeHappiestArrangement(people...), should.Equal, 330)
}

const exampleInput = `
Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
`
