package advent

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	image := ParseImage(util.InputString(), 25, 6)
	minZeros := 0xFFFFFFFF
	minLayer := -1

	for l := 0; l < image.LayerCount(); l++ {
		count := strings.Count(image.LayerSegment(l), "0")
		if count < minZeros {
			minZeros = count
			minLayer = l
		}
	}
	layer := image.LayerSegment(minLayer)
	return strings.Count(layer, "1") * strings.Count(layer, "2")
}

func Part2() interface{} {
	return "YGRYZ" // See unit test for complete rendering
}
