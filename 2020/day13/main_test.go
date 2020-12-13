package advent

import (
	"fmt"
	"strings"
	"testing"

	"advent/lib/util"
)

//lines := []string{"939", "7,13,x,x,59,x,31,19"}
// 29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,653,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19

func TestA(t *testing.T) {
	s := 1
	busses := util.InputLines()[1]
	t.Log(busses)
	for x, num := range strings.Split(busses, ",") {
		if num == "x" {
			continue
		}
		s *= util.ParseInt(num)
		fmt.Println(x, num, s)
	}
}

 //100000000000000
//2283338533368659