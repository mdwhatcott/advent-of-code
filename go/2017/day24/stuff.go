package day24

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
)

func FindStrongestBridge(lines []string) (max int) {
	out := make(chan Bridge)

	go BuildBridges(parseNodes(lines), 0, 0, 0, out)

	for bridge := range out {
		if bridge.Strength > max {
			max = bridge.Strength
		}
	}

	return max
}

func FindStrongestLongestBridge(lines []string) (strongest int) {
	out := make(chan Bridge)

	go BuildBridges(parseNodes(lines), 0, 0, 0, out)

	var longest int
	for bridge := range out {
		if bridge.Length > longest {
			longest = bridge.Length
			strongest = 0
		}
		if bridge.Length >= longest {
			if bridge.Strength >= strongest {
				strongest = bridge.Strength
			}
		}
	}

	return strongest
}

type Node struct {
	A, B int
	Used bool
}

func (this *Node) other(that int) int {
	if that == this.A {
		return this.B
	} else {
		return this.A
	}
}

func parseNodes(lines []string) (result []*Node) {
	for _, line := range lines {
		fields := strings.Split(line, "/")
		a, b := parse.Int(fields[0]), parse.Int(fields[1])
		result = append(result, &Node{A: a, B: b})
	}
	return result
}

func BuildBridges(nodes []*Node, from, strength, length int, out chan Bridge) {
	for _, node := range nodes {
		if node.Used {
			continue
		}
		if node.A == from || node.B == from {
			bridge := Bridge{
				Strength: strength + node.A + node.B,
				Length:   length + 1,
			}
			out <- bridge
			node.Used = true
			BuildBridges(nodes, node.other(from), bridge.Strength, bridge.Length, out)
			node.Used = false
		}
	}
	if strength == 0 {
		close(out)
	}
}

type Bridge struct {
	Length   int
	Strength int
}
