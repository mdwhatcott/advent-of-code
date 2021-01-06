package advent

import (
	"math"
	"strconv"
	"strings"
)

func parseDimensions(line string) (l int, h int, w int) {
	fields := strings.Split(line, "x")
	l, _ = strconv.Atoi(fields[0])
	h, _ = strconv.Atoi(fields[1])
	w, _ = strconv.Atoi(fields[2])
	return
}
func howMuchPaper(l, h, w int) int {
	a := l * h
	b := h * w
	c := w * l
	return (2 * a) + (2 * b) + (2 * c) + min(a, b, c)
}

func min(a, b, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}

func howMuchRibbon(l, h, w int) int {
	wrap := min(perimeter(l, h), perimeter(h, w), perimeter(w, l))
	bow := l * h * w
	return wrap + bow
}

func perimeter(x, y int) int {
	return (x * 2) + (y * 2)
}
