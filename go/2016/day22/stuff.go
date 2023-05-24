package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

type Disk []Node

func (this Disk) String() string {
	buffer := new(bytes.Buffer)
	for _, node := range this {
		if node.Y == 0 {
			buffer.WriteString("\n")
		}
		fmt.Fprintf(buffer, "%d/%d  ", node.Used, node.Size)
	}
	return buffer.String()
}

func scanDisk() Disk {
	lines := util.InputLines()
	nodes := Disk{}
	for _, line := range lines[2:] {
		line = strings.Replace(line, "-x", " ", -1)
		line = strings.Replace(line, "-y", " ", -1)
		line = strings.Replace(line, "T", "", -1)
		line = strings.Replace(line, "/dev/grid/node", "", -1)
		fields := strings.Fields(line)
		nodes = append(nodes, Node{
			X:     parse.Int(fields[0]),
			Y:     parse.Int(fields[1]),
			Size:  parse.Int(fields[2]),
			Used:  parse.Int(fields[3]),
			Avail: parse.Int(fields[4]),
		})
	}
	return nodes
}

type Node struct {
	X     int
	Y     int
	Size  int
	Used  int
	Avail int
}
