package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func Test9Example(t *testing.T) {
	shortest, longest := CalculateShortestAndLongestDistances(input9_Example)
	a := assert.Error(t)
	a.So(shortest, should.Equal, 605)
	a.So(longest, should.Equal, 982)
}

func Test9Real(t *testing.T) {
	shortest, longest := CalculateShortestAndLongestDistances(input9)
	a := assert.Error(t)
	a.So(shortest, should.Equal, 141)
	a.So(longest, should.Equal, 736)
}

var input9_Example = strings.Split(`London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`, "\n")

var input9 = strings.Split(`AlphaCentauri to Snowdin = 66
AlphaCentauri to Tambi = 28
AlphaCentauri to Faerun = 60
AlphaCentauri to Norrath = 34
AlphaCentauri to Straylight = 34
AlphaCentauri to Tristram = 3
AlphaCentauri to Arbre = 108
Snowdin to Tambi = 22
Snowdin to Faerun = 12
Snowdin to Norrath = 91
Snowdin to Straylight = 121
Snowdin to Tristram = 111
Snowdin to Arbre = 71
Tambi to Faerun = 39
Tambi to Norrath = 113
Tambi to Straylight = 130
Tambi to Tristram = 35
Tambi to Arbre = 40
Faerun to Norrath = 63
Faerun to Straylight = 21
Faerun to Tristram = 57
Faerun to Arbre = 83
Norrath to Straylight = 9
Norrath to Tristram = 50
Norrath to Arbre = 60
Straylight to Tristram = 27
Straylight to Arbre = 81
Tristram to Arbre = 90`, "\n")
