package advent

import (
	"io/ioutil"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestRecipeFixture(t *testing.T) {
	gunit.Run(new(RecipeFixture), t)
}

type RecipeFixture struct {
	*gunit.Fixture

	calculator *Calculator
}

func (this *RecipeFixture) Setup() {
	this.calculator = NewCalculator()
	this.calculator.log.SetFlags(0)
	this.calculator.log.SetOutput(ioutil.Discard)
}

func (this *RecipeFixture) TestExample1() {
	this.calculator.ParseAllConversions([]string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	})

	this.calculator.Resolve(map[string]int{"FUEL": 1}, 0)

	this.So(this.calculator.TotalOREConsumed(), should.Equal, 31)
}

func (this *RecipeFixture) TestExample2() {
	this.calculator.ParseAllConversions([]string{
		"9 ORE => 2 A",
		"8 ORE => 3 B",
		"7 ORE => 5 C",
		"3 A, 4 B => 1 AB",
		"5 B, 7 C => 1 BC",
		"4 C, 1 A => 1 CA",
		"2 AB, 3 BC, 4 CA => 1 FUEL",
	})

	this.calculator.Resolve(map[string]int{"FUEL": 1}, 0)

	this.So(this.calculator.TotalOREConsumed(), should.Equal, 165)
}

func (this *RecipeFixture) TestExample3() {
	this.calculator.ParseAllConversions([]string{
		"157 ORE => 5 A",
		"165 ORE => 6 B",
		"179 ORE => 7 C",
		"177 ORE => 5 D",
		"165 ORE => 2 E",

		"12 D, 1 E, 8 C => 9 F",
		"7 B, 7 C => 2 G",
		"3 B, 7 A, 5 D, 10 C => 8 H",

		"44 G, 5 H, 1 F, 29 A, 9 E, 48 D => 1 FUEL",
	})

	this.calculator.Resolve(map[string]int{"FUEL": 1}, 0)

	this.So(this.calculator.TotalOREConsumed(), should.Equal, 13312)
}

func (this *RecipeFixture) SkipTestExample4() {
	this.calculator.log.SetOutput(this)

	this.calculator.ParseAllConversions([]string{
		"139 ORE  => 4 A",
		"144 ORE  => 7 B",
		"145 ORE  => 6 C",
		"176 ORE  => 6 D",

		"1  A        => 8 E",
		"17 A, 3  B  => 8 F",
		"22 D, 37 C  => 5 G",
		"1  D, 6  C  => 4 H",

		"2 F, 7 G, 2 E, 11 C        => 1 I",
		"5 C, 7 H, 2 G, 2  F, 19 E  => 3 J",
		"5 D, 7 C, 9 F, 37 E        => 6 K",

		"53 I, 6 C, 46 D, 81 J, 68 E, 25 K  => 1 FUEL",
	})

	this.calculator.Resolve(map[string]int{"FUEL": 1}, 0)

	this.So(this.calculator.TotalOREConsumed(), should.Equal, 180697)
}
