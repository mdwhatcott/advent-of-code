package day20

import (
	"sort"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() int {
	var pixels []*Pixel
	for id, line := range util.InputLines() {
		pixels = append(pixels, ParsePixel(id, line))
	}
	for x := 0; ; x++ {
		for _, pixel := range pixels {
			pixel.Update()
		}
		sort.Slice(pixels, func(i, j int) bool {
			return pixels[i].DistanceFromOrigin() < pixels[j].DistanceFromOrigin()
		})
		if x > 10000 && pixels[0].DistanceFromOrigin() < pixels[1].DistanceFromOrigin() {
			break // we should have a clear winner after 10000 'ticks'
		}
	}
	return pixels[0].ID
}

func Part2() int {
	var pixels []*Pixel
	for id, line := range util.InputLines() {
		pixels = append(pixels, ParsePixel(id, line))
	}
	for x := 0; x < 100; x++ { // 100 ticks turns out to be enough to detect all collisions for my input.
		for _, pixel := range pixels {
			pixel.Update()
		}
		collisions := detectCollisions(pixels)
		for collision := range collisions {
			for _, id := range collisions[collision] {
				index := 0
				for i, pixel := range pixels {
					if pixel.ID == id {
						index = i
					}
				}
				// remove colliding pixel: (see https://github.com/golang/go/wiki/SliceTricks)
				pixels[index] = pixels[len(pixels)-1]
				pixels[len(pixels)-1] = nil
				pixels = pixels[:len(pixels)-1]
			}
		}
		sort.Slice(pixels, func(i, j int) bool {
			return pixels[i].DistanceFromOrigin() < pixels[j].DistanceFromOrigin()
		})
	}
	return len(pixels)
}
