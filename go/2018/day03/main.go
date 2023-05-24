package day03

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() interface{} {
	fabric := plotClaimsOnFabric(parseClaims(util.InputString()))

	conflict := 0
	for point := range fabric {
		if len(fabric[point].claims) > 1 {
			conflict++
		}
	}
	return conflict
}

func Part2() interface{} {
	claims := parseClaims(util.InputString())
	fabric := plotClaimsOnFabric(claims)
	for id, claim := range claims {
		if fabric.IsUndisputed(claim) {
			return id + 1
		}
	}
	panic("Undisputed claim not found.")
}
