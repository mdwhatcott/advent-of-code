package advent

import (
	"strings"

	"advent/lib/util"
)

var exampleInput = strings.Split(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`, "\n")

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
	//for in, out := range bags {
	//	fmt.Println(in, "--", strings.Join(out, " | "))
	//}

	return len(bags.HowManyCanHold("shiny gold"))
}

func Part2() interface{} {
	return nil
}

type Bags map[string][]string

func (this Bags) HowManyCanHold(color string) (containers map[string]struct{}) {
	outers := this[color]
	if len(outers) == 0 {
		//fmt.Println("DONE:", color, this[color], this)
		return nil
	}
	set := make(map[string]struct{})
	for _, outer := range outers {
		set[outer] = struct{}{}
		for container := range this.HowManyCanHold(outer) {
			set[container] = struct{}{}
		}
	}
	//fmt.Println("SET:", set)
	return set
}

