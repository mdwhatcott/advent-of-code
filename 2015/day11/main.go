package main

import (
	"fmt"

	"advent/lib/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	v1 := util.InputString()
	v2 := Increment(v1)
	v3 := Increment(v2)

	fmt.Println("If the first password is:", v1, " then the next 2 would be:")
	fmt.Println(assert.So(v2, should.Equal, "cqjxxyzz"))
	fmt.Println(assert.So(v3, should.Equal, "cqkaabcc"))
}
