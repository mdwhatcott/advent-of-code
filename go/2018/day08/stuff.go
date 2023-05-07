package day08

import (
	"github.com/mdwhatcott/advent-of-code/go/lib/maths"
)

func RootValue(node Node) (sum int) {
	if len(node.Children) == 0 {
		return maths.Sum[int](node.Metadata...)
	}
	for _, meta := range node.Metadata {
		if 1 <= meta && meta <= len(node.Children) {
			sum += RootValue(node.Children[meta-1])
		}
	}
	return sum
}

func SumMetadata(node Node) (sum int) {
	sum += maths.Sum[int](node.Metadata...)
	for _, child := range node.Children {
		sum += SumMetadata(child)
	}
	return sum
}

func LoadInputs(input []int) chan int {
	c := make(chan int)
	go func() {
		for _, i := range input {
			c <- i
		}
	}()
	return c
}

type Node struct {
	Metadata []int
	Children []Node
}

// See: https://www.michaelfogleman.com/aoc18/#8
func ParseTree(source chan int) (node Node) {
	numChildren := <-source
	numMetadata := <-source

	for x := 0; x < numChildren; x++ {
		node.Children = append(node.Children, ParseTree(source))
	}

	for x := 0; x < numMetadata; x++ {
		node.Metadata = append(node.Metadata, <-source)
	}

	return node
}
