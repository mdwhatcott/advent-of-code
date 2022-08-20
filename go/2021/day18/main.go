package day18

import (
	"encoding/json"
	"fmt"
	"math"
)

type Node struct {
	Depth  int
	IsLeaf bool
	Value  int
	Left   *Node
	Right  *Node
	Prev   *Node
	Next   *Node
}

func NewNode(v interface{}, depth int) *Node {
	n := &Node{Depth: depth}
	pair, ok := v.([]interface{})
	n.IsLeaf = !ok
	if n.IsLeaf {
		n.Value = int(v.(float64))
	} else {
		n.Left = NewNode(pair[0], depth+1)
		n.Right = NewNode(pair[1], depth+1)
	}
	return n
}

func (this *Node) Visit() (result []*Node) {
	if this == nil {
		return result
	}
	result = append(result, this)
	if !this.IsLeaf {
		result = append(result, this.Left.Visit()...)
		result = append(result, this.Right.Visit()...)
	}
	return result
}

func (this *Node) Magnitude() int {
	if this.IsLeaf {
		return this.Value
	}
	return this.Left.Magnitude()*3 + this.Right.Magnitude()*2
}

func (this *Node) String() string {
	if this.IsLeaf {
		return fmt.Sprint(this.Value)
	}
	return fmt.Sprintf("[%s,%s]", this.Left.String(), this.Right.String())
}

func ParseTree(raw string) *Node {
	var j []interface{}
	err := json.Unmarshal([]byte(raw), &j)
	if err != nil {
		panic(err)
	}
	root := NewNode(j, 0)
	var leaves []*Node
	for _, node := range root.Visit() {
		if node.IsLeaf {
			leaves = append(leaves, node)
		}
	}
	for l := 1; l < len(leaves); l++ {
		leaves[l-1].Next = leaves[l]
		leaves[l].Prev = leaves[l-1]
	}
	return root
}

func Add(a *Node, b *Node) *Node {
	return ParseTree(fmt.Sprintf("[%s,%s]", a.String(), b.String()))
}

func Process(tree *Node) bool {
	for _, node := range tree.Visit() {
		if node.Depth == 4 && !node.IsLeaf {
			explode(node)
			return false
		}
	}
	for _, node := range tree.Visit() {
		if node.IsLeaf && node.Value >= 10 {
			split(node)
			return false
		}
	}
	return true
}
func explode(node *Node) {
	if node.Left.Prev != nil {
		node.Left.Prev.Value += node.Left.Value
	}
	if node.Right.Next != nil {
		node.Right.Next.Value += node.Right.Value
	}
	node.IsLeaf = true
	node.Value = 0
}
func split(node *Node) {
	v := float64(node.Value) / 2.0
	node.Left = NewNode(math.Floor(v), node.Depth+1)
	node.Right = NewNode(math.Ceil(v), node.Depth+1)
	node.IsLeaf = false
}

func Reduce(node *Node) *Node {
	for {
		node = ParseTree(node.String())
		if Process(node) {
			return node
		}
	}
}
func Sum(lines []string) *Node {
	node := ParseTree(lines[0])
	for _, line := range lines[1:] {
		node = Reduce(Add(node, ParseTree(line)))
	}
	return node
}
func MaxSumPair(lines []string) (max int) {
	for a, A := range lines {
		for b, B := range lines {
			if a == b {
				continue
			}
			sum := Reduce(Add(ParseTree(A), ParseTree(B))).Magnitude()
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
