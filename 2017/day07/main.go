package day07

import "advent/lib/util"

const test = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func Answers() (part1 string, part2 int) {
	tower := NewTower()
	scanner := util.InputScanner()
	for scanner.Scan() {
		tower.AddProgram(scanner.Text())
	}
	bottom := tower.FindBottom()
	node := tower.listing[bottom]
	diff, value := node.FindUnbalance()

	return bottom, diff+value
}
