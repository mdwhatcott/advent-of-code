package advent

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test7A(t *testing.T) {
	circuit := NewRecursiveCircuit(input7_ExampleA)
	should.So(t, circuit.SolveFor("x"), should.Equal, 123)
	should.So(t, circuit.SolveFor("y"), should.Equal, 456)
	should.So(t, circuit.SolveFor("d"), should.Equal, 72)
	should.So(t, circuit.SolveFor("e"), should.Equal, 507)
	should.So(t, circuit.SolveFor("f"), should.Equal, 492)
	should.So(t, circuit.SolveFor("g"), should.Equal, 114)
	should.So(t, circuit.SolveFor("h"), should.Equal, 65412)
	should.So(t, circuit.SolveFor("i"), should.Equal, 65079)
}

func Test7(t *testing.T) {
	circuit := NewRecursiveCircuit(util.InputLines())
	should.So(t, circuit.SolveFor("a"), should.Equal, 3176)
	//t.Log("Signal on the 'a' wire:", circuit.SolveFor("a"))
}

func Test7_Part2(t *testing.T) {
	a := NewRecursiveCircuit(util.InputLines()).SolveFor("a")
	override := strconv.FormatUint(uint64(a), 10)

	circuit := NewRecursiveCircuit(util.InputLines())
	circuit.instructions["b"] = []string{override}

	should.So(t, circuit.SolveFor("a"), should.Equal, 14710)
	//t.Log("Updated signal on the 'a' wire:", circuit.SolveFor("a"))
}

var input7_ExampleA = strings.Split(`123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`, "\n")
