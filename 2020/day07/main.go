package advent

import (
	"strings"

	"advent/lib/util"
)

var exampleInput1 = strings.Split(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`, "\n")

var exampleInput2 = strings.Split(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`, "\n")

func Part1() interface{} {
	colors := make(map[string]struct{})
	bags := make(Bags)
	for _, line := range util.InputLines() {
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")
		leftRight := strings.Split(line, " contain ")
		right := strings.Split(leftRight[1], ", ")
		for _, color := range right {
			words := strings.Fields(color)
			color := strings.Join(words[1:], " ")
			bags[color] = append(bags[color], leftRight[0])
		}
		colors[leftRight[0]] = struct{}{}
	}

	return len(bags.HowManyCanHold("shiny gold"))
}

type Bags map[string][]string

func (this Bags) HowManyCanHold(color string) (containers map[string]struct{}) {
	outers := this[color]
	if len(outers) == 0 {
		return nil
	}
	set := make(map[string]struct{})
	for _, outer := range outers {
		set[outer] = struct{}{}
		for container := range this.HowManyCanHold(outer) {
			set[container] = struct{}{}
		}
	}
	return set
}

func Part2() interface{} {
	bags := make(NestedBags)
	for _, line := range util.InputLines() {
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")
		leftRight := strings.Split(line, " contain ")
		outer := leftRight[0]
		inners := strings.Split(leftRight[1], ", ")
		for _, inner := range inners {
			words := strings.Fields(inner)
			bags[outer] = append(bags[outer], Bag{
				Color: strings.Join(words[1:], " "),
				Quantity: util.ParseInt(words[0]),
			})
		}
	}
	return bags.HowManyAreHeldIn("shiny gold")
}

type NestedBags map[string][]Bag

func (this NestedBags) HowManyAreHeldIn(color string) (result int) {
	inners := this[color]
	for _, inner := range inners {
		result += inner.Quantity
		result += inner.Quantity * this.HowManyAreHeldIn(inner.Color)
	}
	return result
}

type Bag struct {
	Color    string
	Quantity int
}
