package advent

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
