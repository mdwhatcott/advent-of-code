package assembunny

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestInterpreterFixture(t *testing.T) {
	should.Run(&InterpreterFixture{T: should.New(t)}, should.Options.UnitTests())
}

type InterpreterFixture struct {
	*should.T
	a, b, c, d int
}

func (this *InterpreterFixture) execute(a, b, c, d int, program string) {
	interpreter := NewInterpreter(this.prepare(program))

	interpreter.Set("a", a)
	interpreter.Set("b", b)
	interpreter.Set("c", c)
	interpreter.Set("d", d)

	interpreter.ExecuteProgram()

	this.a = interpreter.Get("a")
	this.b = interpreter.Get("b")
	this.c = interpreter.Get("c")
	this.d = interpreter.Get("d")
}
func (this *InterpreterFixture) prepare(program string) []string {
	program = strings.TrimSpace(program)
	program = strings.Replace(program, "\t", "", -1)
	return strings.Split(program, "\n")
}

func (this *InterpreterFixture) TestSample() {
	this.execute(0, 0, 0, 0, `
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
	this.execute(0, 0, 0, 0, `
		cpy 41 a
		inc a
		jnz a 2
		inc a
		dec a
		dec a
	`)
	this.So(this.a, should.Equal, 40)
}

func (this *InterpreterFixture) Test_Day12_Part1_Input() {
	this.execute(0, 0, 0, 0, `
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

func (this *InterpreterFixture) Test_tgl_BeyondEndOfProgramIsNop() {
	this.execute(0, 0, 0, 0, `
		cpy 1 a
		tgl 2
		inc a
	`)
	this.So(this.a, should.Equal, 2)
}

func (this *InterpreterFixture) Test_tgl_BeyondBeginningOfProgramIsNop() {
	this.execute(0, 0, 0, 0, `
		cpy 1 a
		tgl -100
		inc a
	`)
	this.So(this.a, should.Equal, 2)
}

func (this *InterpreterFixture) Test_tgl_turns_inc_into_dec() {
	this.execute(0, 0, 0, 0, `
		cpy 1 a
		tgl 1
		inc a
	`)
	this.So(this.a, should.Equal, 0)
}

func (this *InterpreterFixture) Test_tgl_turns_dec_into_inc() {
	this.execute(0, 0, 0, 0, `
		cpy 1 a
		tgl 1
		dec a
	`)
	this.So(this.a, should.Equal, 2)
}

func (this *InterpreterFixture) Test_tgl_turns_tgl_into_inc() {
	this.execute(0, 0, 0, 0, `
		cpy 1 a
		tgl 1
		tgl a
	`)
	this.So(this.a, should.Equal, 2)
}

func (this *InterpreterFixture) Test_tgl_turns_jnz_into_cpy() {
	this.execute(0, 0, 0, 0, `
		cpy 1 a
		tgl 1
		jnz 2 a
	`)
	this.So(this.a, should.Equal, 2)
}

func (this *InterpreterFixture) Test_tgl_turns_cpy_into_jnz() {
	this.execute(0, 0, 0, 0, `
		cpy 2 a
		tgl 1
		cpy 1 a
		inc a
	`)
	this.So(this.a, should.Equal, 2)
}

func (this *InterpreterFixture) Test_Day23_Sample1() {
	this.execute(0, 0, 0, 0, `
		cpy 2 a
		tgl a
		tgl a
		tgl a
		cpy 1 a
		dec a
		dec a
	`)
	this.So(this.a, should.Equal, 3)
}

func (this *InterpreterFixture) Test_Day23_Part1() {
	this.execute(7, 0, 0, 0, `
		cpy a b
		dec b
		cpy a d
		cpy 0 a
		cpy b c
		inc a
		dec c
		jnz c -2
		dec d
		jnz d -5
		dec b
		cpy b c
		cpy c d
		dec d
		inc c
		jnz d -2
		tgl c
		cpy -16 c
		jnz 1 c
		cpy 78 c
		jnz 99 d
		inc a
		inc d
		jnz d -2
		inc c
		jnz c -5
	`)
	this.So(this.a, should.Equal, 12762)
}
