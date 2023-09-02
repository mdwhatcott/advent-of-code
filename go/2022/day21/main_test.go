package day21

import (
	"log"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	. "github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

const TODO = 21

var (
	inputLines  = inputs.Read(2022, 21).Lines()
	sampleLines = []string{
		"root: pppw + sjmn",
		"dbpl: 5",
		"cczh: sllz + lgvd",
		"zczc: 2",
		"ptdq: humn - dvpt",
		"dvpt: 3",
		"lfqf: 4",
		"humn: 5",
		"ljgn: 2",
		"sjmn: drzm * dbpl",
		"sllz: 4",

		"pppw: cczh / lfqf",
		"lgvd: ljgn * ptdq",
		"drzm: hmdt - zczc",
		"hmdt: 32",
	}
)

func Test(t *testing.T) {
	log.SetFlags(0)

	//should.So(t, Part1(sampleLines), should.Equal, 152)
	//should.So(t, Part1(inputLines), should.Equal, 62386792426088)

	should.So(t, Part2(sampleLines), should.Equal, 301)
	//should.So(t, Part2(inputLines), should.Equal, TODO)
}

func Part1(lines []string) any {
	monkeys := set.Of(Map(parse, lines)...)
	human := First(Filter(isHuman, monkeys.Slice()))
	monkeys.Remove(human)
	return try(monkeys.Slice(), human)
}
func Part2(lines []string) any {
	monkeys := set.Of(Map(parse2, lines)...)
	human := First(Filter(isHuman, monkeys.Slice()))
	monkeys.Remove(human)
	noHuman := monkeys.Slice()

	x := 0
	diff := 1
	for times := 0; times < 100; times++ {
		a := try(noHuman, Monkey{Name: "humn", Value: x})
		b := try(noHuman, Monkey{Name: "humn", Value: x + diff})
		if a == 0 {
			return x
		}
		if b == 0 {
			return x + 1
		}
		diff = a - b
		log.Printf("x:%-20d a:%-20d b:%-20d diff:%d", x, a, b, diff)
		if diff > a {
			diff = a
		} else {
			x += diff
		}
	}
	panic("boink")
}

func try(monkeys []Monkey, human Monkey) int {
	index := make(map[string]int)
	population := append(monkeys[:], human)
	for {
		for _, monkey := range population {
			if monkey.Operator == "" {
				index[monkey.Name] = monkey.Value
			} else if a, containsA := index[monkey.OperandA]; !containsA {
				continue
			} else if b, containsB := index[monkey.OperandB]; !containsB {
				continue
			} else {
				switch monkey.Operator {
				case "+":
					index[monkey.Name] = a + b
				case "-":
					index[monkey.Name] = a - b
				case "*":
					index[monkey.Name] = a * b
				case "/":
					index[monkey.Name] = a / b
				}
				monkey.Operator = ""
			}
			if value, contains := index["root"]; contains {
				return value
			}
		}
	}
}

func isHuman(m Monkey) bool { return m.Name == "humn" }
