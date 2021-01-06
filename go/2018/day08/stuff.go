package day08

import "advent/lib/util"

func RootValue(node Node) (sum int) {
	if len(node.Children) == 0 {
		return util.Ints(node.Metadata).Sum()
	}
	for _, meta := range node.Metadata {
		if 1 <= meta && meta <= len(node.Children) {
			sum += RootValue(node.Children[meta-1])
		}
	}
	return sum
}

func SumMetadata(node Node) (sum int) {
	sum += util.Ints(node.Metadata).Sum()
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
	num_children := <-source
	num_metadata := <-source

	for x := 0; x < num_children; x++ {
		node.Children = append(node.Children, ParseTree(source))
	}

	for x := 0; x < num_metadata; x++ {
		node.Metadata = append(node.Metadata, <-source)
	}

	return node
}
