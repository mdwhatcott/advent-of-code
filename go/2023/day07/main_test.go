package starter

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2023, 7).Lines()
	sampleLines = []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct{ *should.T }

func (this *Suite) TestPart1A() {
	this.So(Part1(sampleLines), should.Equal, 6440)
}
func (this *Suite) TestPart1Full() {
	this.So(Part1(inputLines), should.Equal, 250453939)
}
func (this *Suite) TestPart2A() {
	this.So(Part2(sampleLines), should.Equal, 5905)
}
func (this *Suite) TestPart2Full() {
	this.So(Part2(inputLines), should.Equal, 248652697)
}
func TestPart2Type(t *testing.T) {
	tests := []struct {
		name     string
		expected HandType
	}{
		{name: "JJJJJ 1", expected: FiveOfAKind},
		{name: "1JJJJ 1", expected: FiveOfAKind},
		{name: "12JJJ 1", expected: FourOfAKind},
		{name: "11JJJ 1", expected: FiveOfAKind},
		{name: "111JJ 1", expected: FiveOfAKind},
		{name: "123JJ 1", expected: ThreeOfAKind},
		{name: "122JJ 1", expected: FourOfAKind},
		{name: "1222J 1", expected: FourOfAKind},
		{name: "1122J 1", expected: FullHouse},
		{name: "1112J 1", expected: FourOfAKind},
		{name: "1234J 1", expected: OnePair},
		{name: "1111J 1", expected: FiveOfAKind},
		{name: "1123J 1", expected: ThreeOfAKind},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hand := ParsePart2Hand(test.name)
			t.Log("Original:", hand.Part2String())
			t.Log("Parsed:  ", string(hand))
			should.So(t, hand.Part2String(), should.Equal, test.name)
			should.So(t, hand.Part2Type(), should.Equal, test.expected)
		})
	}
}
