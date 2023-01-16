package day11

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test(t *testing.T) {
	v1 := util.InputString()
	v2 := Increment(v1)
	v3 := Increment(v2)

	should.So(t, v2, should.Equal, "cqjxxyzz")
	should.So(t, v3, should.Equal, "cqkaabcc")
}
