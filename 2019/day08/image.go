package advent

import (
	"fmt"
	"strings"
)

type Image struct {
	input  string
	width  int
	height int
}

func ParseImage(input string, width int, height int) Image {
	return Image{input: input, width: width, height: height}
}

func (this Image) PixelCount() int {
	return this.width * this.height
}

func (this Image) LayerCount() int {
	return len(this.input) / this.PixelCount()
}

func (this Image) LayerSegment(layer int) string {
	pixelsPerLayer := this.PixelCount()
	start := pixelsPerLayer * layer
	stop := start + pixelsPerLayer
	return this.input[start:stop]
}

func (this Image) RenderPixel(x int, y int) string {
	for l := 0; l < this.LayerCount(); l++ {
		segment := this.LayerSegment(l)
		start := y*this.width + x
		point := segment[start : start+1]
		if point == "0" {
			return "0"
		} else if point == "1" {
			return "1"
		}
	}
	panic(fmt.Sprintf("Could not render pixel (%d,%d)", x, y))
}

func (this Image) RenderFull() string {
	var builder strings.Builder
	for r := 0; r < this.height; r++ {
		for c := 0; c < this.width; c++ {
			builder.WriteString(pixel[this.RenderPixel(c, r)])
		}
		builder.WriteString("\n")
	}
	return strings.TrimSpace(builder.String())
}

var pixel = map[string]string{
	"0": " ",
	"1": "*",
}
