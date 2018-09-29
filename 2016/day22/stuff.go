package main

import (
	"bytes"
	"fmt"
	"strings"

	"advent/lib/util"
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
			X:     util.ParseInt(fields[0]),
			Y:     util.ParseInt(fields[1]),
			Size:  util.ParseInt(fields[2]),
			Used:  util.ParseInt(fields[3]),
			Avail: util.ParseInt(fields[4]),
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
