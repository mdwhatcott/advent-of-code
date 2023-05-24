package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Test9Example(t *testing.T) {
	shortest, longest := CalculateShortestAndLongestDistances(input9_Example)
	should.So(t, shortest, should.Equal, 605)
	should.So(t, longest, should.Equal, 982)
}

func Test9Real(t *testing.T) {
	shortest, longest := CalculateShortestAndLongestDistances(util.InputLines())
	should.So(t, shortest, should.Equal, 141)
	should.So(t, longest, should.Equal, 736)
}

var input9_Example = strings.Split(`London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`, "\n")
