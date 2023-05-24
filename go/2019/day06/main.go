package advent

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() interface{} {
	return assembleOrbitalSystem(util.InputLines()).OrbitalChecksum()
}

func Part2() interface{} {
	return assembleOrbitalSystem(util.InputLines()).OrbitalDistance("YOU", "SAN")
}
