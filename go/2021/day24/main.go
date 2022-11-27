package day24

import (
	"fmt"
	"io"
	"log"
	"strings"
	"unicode"

	"advent/lib/util"
)

func init() { log.SetFlags(0) }

type Interpreter struct {
	writer     io.Writer
	inputs     chan int
	program    []string
	w, x, y, z int
}

func NewInterpreter(writer io.Writer, inputs chan int, program []string) *Interpreter {
	return &Interpreter{writer: writer, inputs: inputs, program: program}
}
func (i *Interpreter) logf(format string, args ...any) { _, _ = fmt.Fprintf(i.writer, format, args...) }
func (i *Interpreter) log(args ...any)                 { _, _ = fmt.Fprint(i.writer, args...) }
func (i *Interpreter) Run() (w, x, y, z int) {
	defer i.reset()
	for pc := 0; pc < len(i.program); pc++ {
		s := i.program[pc]
		i.logf("%16d %16d %16d %16d    %s", i.w, i.x, i.y, i.z, s)
		w := strings.Fields(s)
		a := w[1]
		switch w[0] {
		case "inp":
			i.log(strings.Repeat("-", 80))
			i.set(a, <-i.inputs)
		case "add":
			i.set(a, i.get(a)+i.resolve(w[2]))
		case "mul":
			i.set(a, i.get(a)*i.resolve(w[2]))
		case "div":
			i.set(a, i.get(a)/i.resolve(w[2]))
		case "mod":
			i.set(a, i.get(a)%i.resolve(w[2]))
		case "eql":
			if i.get(a) == i.resolve(w[2]) {
				i.set(a, 1)
			} else {
				i.set(a, 0)
			}
		}
	}
	i.logf("%16d %16d %16d %16d    %s", i.w, i.x, i.y, i.z, "")
	return i.w, i.x, i.y, i.z
}
func (i *Interpreter) set(v string, n int) {
	switch v {
	case "w":
		i.w = n
	case "x":
		i.x = n
	case "y":
		i.y = n
	case "z":
		i.z = n
	default:
		panic("bad key: " + v)
	}
}
func (i *Interpreter) get(v string) int {
	switch v {
	case "w":
		return i.w
	case "x":
		return i.x
	case "y":
		return i.y
	case "z":
		return i.z
	}
	panic("bad key: " + v)
}
func (i *Interpreter) resolve(s string) int {
	if len(s) == 1 && unicode.IsLetter(rune(s[0])) {
		return i.get(s)
	}
	return util.ParseInt(s)
}

func (i *Interpreter) reset() {
	i.w, i.x, i.y, i.z = 0, 0, 0, 0
}
