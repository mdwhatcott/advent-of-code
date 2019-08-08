package main

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestInterpreterFixture(t *testing.T) {
	gunit.Run(new(InterpreterFixture), t)
}

type InterpreterFixture struct {
	*gunit.Fixture
	a, b int
}

func (this *InterpreterFixture) execute(a, b int, program string) {
	interpreter := NewInterpreter(this.prepare(program))

	interpreter.Set("a", a)
	interpreter.Set("b", b)

	interpreter.ExecuteProgram()

	this.a = interpreter.Get("a")
	this.b = interpreter.Get("b")
}
func (this *InterpreterFixture) prepare(program string) []string {
	program = strings.TrimSpace(program)
	program = strings.Replace(program, "\t", "", -1)
	return strings.Split(program, "\n")
}

func (this *InterpreterFixture) TestSample() {
	this.execute(0, 0, `
		cpy 41 a
		inc a
		inc a
		dec a
		jnz a 2
		dec a
	`)
	this.So(this.a, should.Equal, 42)
}

func (this *InterpreterFixture) TestSampleWithJumpThatEvaluates() {
	this.execute(0, 0, `
		cpy 41 a
		inc a
		jnz a 2
		inc a
		dec a
		dec a
	`)
	this.So(this.a, should.Equal, 40)
}

func (this *InterpreterFixture) Test_Part1_Input() {
	this.execute(0, 0, `
		cpy 1 a
		cpy 1 b
		cpy 26 d
		jnz c 2
		jnz 1 5
		cpy 7 c
		inc d
		dec c
		jnz c -2
		cpy a c
		inc a
		dec b
		jnz b -2
		cpy c b
		dec d
		jnz d -6
		cpy 19 c
		cpy 14 d
		inc a
		dec d
		jnz d -2
		dec c
		jnz c -5
	`)
	this.So(this.a, should.Equal, 318077)
}
