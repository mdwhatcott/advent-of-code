package day24

import (
	"fmt"
	"strings"

	"advent/lib/util"
)

func MaxBridgeStrength(paths chan string) int {
	var strengths []int
	for path := range paths {
		fmt.Println(calculateStrength(path), path)
		strengths = append(strengths, calculateStrength(path))
	}
	fmt.Println("Total Paths:", len(strengths))
	return util.Max(strengths...)
}

func buildGraph(lines []string) *Node {
	lookup := make(map[int][]*Node)
	for _, line := range lines {
		node := NewNode(line)
		lookup[node.a] = append(lookup[node.a], node)
		if node.a == node.b {
			continue
		}
		lookup[node.b] = append(lookup[node.b], node)
	}
	root := NewNode("-1/0")
	root.Orient(0, lookup)
	return root
}

type Node struct {
	a, b     int
	original string
	children map[*Node]bool
	parents  map[*Node]bool
}

func NewNode(original string) *Node {
	fields := strings.Split(original, "/")
	a, b := fields[0], fields[1]
	A, B := util.ParseInt(a), util.ParseInt(b)
	if original == "-1/0" {
		original = "0"
	}
	return &Node{
		a: A,
		b: B,

		original: original,
		children: make(map[*Node]bool),
		parents:  make(map[*Node]bool),
	}
}

func (this *Node) Orient(to int, lookup map[int][]*Node) {
	for _, node := range lookup[to] {
		if node == this || this.parents[node] {
			continue
		}
		node.parents[this] = true
		for parent := range this.parents {
			node.parents[parent] = true
		}
		if this.children[node] {
			continue
		}
		this.children[node] = true
		node.Orient(node.other(to), lookup)
	}
}

func (this *Node) other(to int) int {
	if this.a == to {
		return this.b
	}
	return this.a
}

func (this *Node) Traverse() (out chan string) {
	out = make(chan string)
	go this.traverse(out, "")
	return out
}

func (this *Node) traverse(out chan string, path string) {
	newPath := strings.TrimLeft(fmt.Sprintf("%s--%s", path, this.original), "--")
	out <- newPath
	for child := range this.children {
		child.traverse(out, newPath)
	}
	if len(path) == 0 {
		close(out)
	}
}

func calculateStrength(path string) (result int) {
	for _, value := range strings.Split(strings.ReplaceAll(path, "--", "/"), "/") {
		result += util.ParseInt(value)
	}
	return result
}