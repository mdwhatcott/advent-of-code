package day24

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func TestALUSuite(t *testing.T) {
	should.Run(&ALUSuite{T: should.New(t)}, should.Options.UnitTests())
}

type ALUSuite struct {
	*should.T
	program []string
}

func (this *ALUSuite) Setup() {
	this.program = util.InputLines()
}

func (this *ALUSuite) TestSample() {
	this.So(fmt.Sprint(NewInterpreter(this, goLoadInt(0b1101), sample1BinaryNumber).Run()), should.Equal, "1 1 0 1")
}

func (this *ALUSuite) TestOneCycle() {
	//NewInterpreter(this, goLoadModelNumber(13579246899999), this.program).Run()
	NewInterpreter(this, goLoadModelNumber(12321232199999), this.program).Run()
}
func (this *ALUSuite) SkipTestScan() {
	c := make(chan int, 14)
	i := NewInterpreter(io.Discard, c, this.program)
	for x := 99999999999999; x >= 11111111111111; x-- {
		loadModelNumber(c, x)
		_, _, _, z := i.Run()
		if z == 0 {
			this.Log("success:", x, z)
		}
		if x%1000 == 0 {
			this.Log(x)
		}
	}
}

func goLoadInt(i int) (c chan int) {
	c = make(chan int, 1)
	c <- i
	return c
}
func goLoadModelNumber(i int) (c chan int) {
	c = make(chan int)
	go loadModelNumber(c, i)
	return c
}
func loadModelNumber(c chan int, i int) {
	for _, d := range fmt.Sprint(i) {
		c <- util.ParseInt(string(d))
	}
}

var sample1BinaryNumber = strings.Split(
	`inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`, "\n")
