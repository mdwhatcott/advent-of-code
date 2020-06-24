package advent

import "advent/lib/util"

func Part1() interface{} {
	return assembleOrbitalSystem(util.InputLines()).OrbitalChecksum()
}

func Part2() interface{} {
	return nil
}
