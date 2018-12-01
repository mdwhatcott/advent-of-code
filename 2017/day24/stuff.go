package day24

import (
	"strings"

	"advent/lib/util"
)

// TODO: implement a BFS or a* search over the graph

func buildGraph(lines []string) *Node {
	var all []*Node
	all = append(all, &Node{A: -1, B: 0})
	for _, line := range lines {
		fields := strings.Split(line, "/")
		all = append(all, &Node{A: util.ParseInt(fields[0]), B: util.ParseInt(fields[1])})
	}
	for _, a := range all {
		for _, b := range all {
			if a != b {
				if a.Matches(b) {
					a.Attach(b)
					b.Attach(a)
				}
			}
		}
	}
	return all[0]
}

type Node struct {
	A, B int
	a, b []*Node
}

func (this *Node) Attach(that *Node) {
	if that.A == this.A || that.B == this.A {
		this.a = append(this.a, that)
	}
	if that.A == this.B || that.B == this.B {
		this.b = append(this.b, that)
	}
}

func (this *Node) Matches(that *Node) bool {
	return (this.a == nil && that.a == nil && this.A == that.A) ||
		(this.b == nil && that.b == nil && this.B == that.B) ||
		(this.a == nil && that.a == nil && this.B == that.B) ||
		(this.b == nil && that.b == nil && this.A == that.A)
}
