package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test3(t *testing.T) {
	visits := make(map[string]int)
	turtle := NewSantaTurtle(visits)
	for _, c := range util.InputString() {
		turtle.Move(c)
	}
	should.So(t, len(visits), should.Equal, 2572)
	//t.Log("How many houses receive at least one present?", len(visits))
}

func Test3_RoboSanta(t *testing.T) {
	visits := make(map[string]int)
	santa := NewSantaTurtle(visits)
	robot := NewSantaTurtle(visits)
	for i, c := range util.InputString() {
		if i%2 == 1 {
			santa.Move(c)
		} else {
			robot.Move(c)
		}
	}
	should.So(t, len(visits), should.Equal, 2631)
	//t.Log("How many houses receive at least one present?", len(visits))
}
