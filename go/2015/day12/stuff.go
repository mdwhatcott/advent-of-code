package main

import (
	"bufio"
	"strings"

	"advent/lib/util"
)

// prepare removes superfluous tokens and adds padding
// so that bufio.ScanWords can be applied to iterate tokens.
func prepare(input string) *bufio.Scanner {
	input = strings.Replace(input, ":", " ", -1)
	input = strings.Replace(input, ",", " ", -1)
	input = strings.Replace(input, "'", " ", -1)
	input = strings.Replace(input, `"`, " ", -1)
	input = strings.Replace(input, "[", " [ ", -1)
	input = strings.Replace(input, "]", " ] ", -1)
	input = strings.Replace(input, "{", " { ", -1)
	input = strings.Replace(input, "}", " } ", -1)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	return scanner
}

/////////////////////////////////////////////////////

type Scope struct {
	sum  int
	red  bool
	safe bool

	parent   *Scope
	children []*Scope
}

func (this *Scope) Receive(token string) *Scope {
	switch token {
	case "red":
		return this.markRed()
	case "{":
		return this.descend()
	case "[":
		return this.descendImmune()
	case "}", "]":
		return this.ascend()
	default:
		return this.add(token)
	}
}
func (this *Scope) markRed() *Scope {
	this.red = !this.safe
	return this
}
func (this *Scope) descend() *Scope {
	child := &Scope{parent: this}
	this.children = append(this.children, child)
	return child
}
func (this *Scope) descendImmune() *Scope {
	child := this.descend()
	child.safe = true
	return child
}
func (this *Scope) ascend() *Scope {
	return this.parent
}
func (this *Scope) add(token string) *Scope {
	this.sum += util.ParseInt(token)
	return this
}

func (this *Scope) Sum() (total int) {
	if this.red {
		return 0
	}

	for _, child := range this.children {
		total += child.Sum()
	}

	return this.sum + total
}
