package day16

import (
	"strings"

	"advent/lib/util"
)

func spin(s string, i int) string {
	divide := len(s) - i
	first := s[:divide]
	last := s[divide:]
	return last + first
}

func exchange(s string, a, b int) string {
	ss := []byte(s)
	ss[a], ss[b] = ss[b], ss[a]
	return string(ss)
}

func partner(s string, a, b string) string {
	ia := strings.Index(s, a)
	ib := strings.Index(s, b)
	return exchange(s, ia, ib)
}

type Interpreter struct {
	state string
}

func NewInterpreter(start string) *Interpreter {
	return &Interpreter{state: start}
}

func (this *Interpreter) Dance(move string) string {
	switch move[0] {
	case 's':
		this.spin(util.ParseInt(move[1:]))
	case 'x':
		elements := strings.Split(move[1:], "/")
		this.exchange(util.ParseInt(elements[0]), util.ParseInt(elements[1]))
	case 'p':
		elements := strings.Split(move[1:], "/")
		this.partner(elements[0], elements[1])
	}
	return this.state
}

func (this *Interpreter) spin(n int) {
	this.state = spin(this.state, n)
}

func (this *Interpreter) exchange(a, b int) {
	this.state = exchange(this.state, a, b)
}

func (this *Interpreter) partner(a, b string) {
	this.state = partner(this.state, a, b)
}

func (interpreter *Interpreter) doCompleteDance() {
	for _, move := range danceSteps {
		interpreter.Dance(move)
	}
}
