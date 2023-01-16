package day19

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test(t *testing.T) {
	should.So(t, part1(), should.Equal, 509)
	should.So(t, part2(), should.Equal, 195)
}

const (
	molecule   = `CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr`
	part1Start = molecule
)

func part1() int {
	machine := NewMoleculeMachine()
	for scanner := util.InputScanner(); scanner.Scan(); {
		line := scanner.Text()
		if line == "" {
			break
		}
		machine.RegisterReplacement(line)
	}

	molecules := machine.Calibrate(part1Start)

	return len(molecules)
}

func part2() (transformations int) {
	var replacements []Pair
	for scanner := util.InputScanner(); scanner.Scan(); {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.Fields(line)
		replacements = append(replacements, Pair{fields[0], fields[2]})
	}

	shuffles := 0
	working := molecule
	for len(working) > 1 {
		checkpoint := working
		for _, pair := range replacements {
			from, to := pair.Key, pair.Value
			for strings.Contains(working, to) {
				transformations += strings.Count(working, to)
				working = strings.Replace(working, to, from, -1)
			}
		}
		if working == checkpoint {
			Shuffle(replacements)
			working = molecule
			transformations = 0
			shuffles++
		}
	}
	return transformations
}

// Credit: https://www.calhoun.io/how-to-shuffle-arrays-and-slices-in-go/
func Shuffle(slice []Pair) {
	generator := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := generator.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

type Pair struct {
	Key, Value string
}

/*
Credit: https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/cy4kpms/

# In Python 2, part two only shown. I work from the molecule removing chunks of the string according to the rules until I get down to a molecule of length 1.
# But I apply the transformations randomly. When I apply a transformation, I apply it as many times as a I can before moving on to the next. If the molecule is not smaller after all transformations, I shuffle all the possible transformations and start again.
# The number of times I have to the transformations over is much much fewer than I thought it would be. I find the solution after 1 to 10 restarts. And it's apparently the same number of transformations each time. This is surprising. I'm guessing it's the nature of the data and not my "awesome" random algorithm.
# It'd be great to get a proof that this algorithm quickly finds the solution (as it appears to in practice), but I'm not sure there is a proof...

The code:

count = shuffles = 0
mol = molecule
while len(mol) > 1:
    start = mol
    for frm, to in transforms:
        while to in mol:
            count += mol.count(to)
            mol = mol.replace(to, frm)

    if start == mol:  # no progress start again
        shuffle(transforms)
        mol = molecule
        count = 0
        shuffles += 1

print('{} transforms after {} shuffles'.format(count, shuffles))
*/
