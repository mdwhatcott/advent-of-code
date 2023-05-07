package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test8_CodeLength(t *testing.T) {
	should.So(t, CodeLength(`""`), should.Equal, 2)
	should.So(t, CodeLength(`"abc"`), should.Equal, 5)
	should.So(t, CodeLength(`"aaa\"aaa"`), should.Equal, 10)
	should.So(t, CodeLength(`"\x27"`), should.Equal, 6)

	should.So(t, CodeLength(`"ffsfyxbyuhqkpwatkjgudo"`), should.Equal, 24)
	should.So(t, CodeLength(`"byc\x9dyxuafof\\\xa6uf\\axfozomj\\olh\x6a"`), should.Equal, 43)
}

func Test8_MemoryLength(t *testing.T) {
	should.So(t, MemoryLength(`""`), should.Equal, 0)
	should.So(t, MemoryLength(`"abc"`), should.Equal, 3)
	should.So(t, MemoryLength(`"aaa\"aaa"`), should.Equal, 7)
	should.So(t, MemoryLength(`"\x27"`), should.Equal, 1)

	should.So(t, MemoryLength(`"ffsfyxbyuhqkpwatkjgudo"`), should.Equal, 22)
	should.So(t, MemoryLength(`"byc\x9dyxuafof\\\xa6uf\\axfozomj\\olh\x6a"`), should.Equal, 29)
}

func Test8_EscapeLength(t *testing.T) {
	should.So(t, EscapedLength(`""`), should.Equal, 6)
	should.So(t, EscapedLength(`"abc"`), should.Equal, 9)
	should.So(t, EscapedLength(`"aaa\"aaa"`), should.Equal, 16)
	should.So(t, EscapedLength(`"\x27"`), should.Equal, 11)
}

func Test8(t *testing.T) {
	code := 0
	memory := 0
	escaped := 0

	for _, line := range strings.Split(util.InputString(), "\n") {
		code += CodeLength(line)
		memory += MemoryLength(line)
		escaped += EscapedLength(line)
	}

	should.So(t, code-memory, should.Equal, 1350)
	should.So(t, escaped-code, should.Equal, 2085)
	//t.Log("Code - Memory:", code-memory)
	//t.Log("Escaped - Code:", escaped - code)
}
