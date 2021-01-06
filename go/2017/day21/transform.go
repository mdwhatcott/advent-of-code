package day21

import "fmt"

func Transformations(pattern string) []string {
	size := sizeOf(pattern)
	if instructions, found := transformations[size]; found {
		return transformPattern(pattern, instructions)
	} else {
		panic(fmt.Sprintln("Grid size something other than 2 or 3:", size))
	}
}

func transformPattern(pattern string, transformations [][]int) (transformed []string) {
	transformed = append(transformed, pattern)
	for _, transformation := range transformations {
		var altered string
		for _, c := range transformation {
			altered += string(pattern[c])
		}
		if !contains(transformed, altered) {
			transformed = append(transformed, altered)
		}
	}
	return transformed
}

var transformations = map[int][][]int{
	2: {
		{3, 0, 2, 4, 1}, // rotate 90
		{4, 3, 2, 1, 0}, // rotate 180
		{1, 4, 2, 0, 3}, // rotate 270
		{3, 4, 2, 0, 1}, // flip horizontal
		{0, 3, 2, 1, 4}, // flip horizontal, rotate 90
		{1, 0, 2, 4, 3}, // flip horizontal, rotate 180
		{4, 1, 2, 3, 0}, // flip horizontal, rotate 270
		{1, 0, 2, 4, 3}, // flip vertical (is same as flip horizontal, rotate 180)
		{4, 1, 2, 3, 0}, // flip vertical, rotate 90 (is the same as flip horizontal, rotate 270)
		{3, 4, 2, 0, 1}, // flip vertical, rotate 180 (is the same as flip horizontal)
		{0, 3, 2, 1, 4}, // flip vertical, rotate 270 (is the same as flip horizontal, rotate 90)
	},
	3: {
		{8, 4, 0, 3, 9, 5, 1, 7, 10, 6, 2}, // rotate 90
		{10, 9, 8, 3, 6, 5, 4, 7, 2, 1, 0}, // rotate 180
		{2, 6, 10, 3, 1, 5, 9, 7, 0, 4, 8}, // rotate 270
		{8, 9, 10, 3, 4, 5, 6, 7, 0, 1, 2}, // flip horizontal
		{0, 4, 8, 3, 1, 5, 9, 7, 2, 6, 10}, // flip horizontal, rotate 90
		{2, 1, 0, 3, 6, 5, 4, 7, 10, 9, 8}, // flip horizontal, rotate 180
		{10, 6, 2, 3, 9, 5, 1, 7, 8, 4, 0}, // flip horizontal, rotate 270
		{2, 1, 0, 3, 6, 5, 4, 7, 10, 9, 8}, // flip vertical (same as flip horizontal, rotate 180)
		{10, 6, 2, 3, 9, 5, 1, 7, 8, 4, 0}, // flip vertical, rotate 90 (same as flip horizontal, rotate 270)
		{8, 9, 10, 3, 4, 5, 6, 7, 0, 1, 2}, // flip vertical, rotate 180 (same as flip horizontal)
		{10, 6, 2, 3, 9, 5, 1, 7, 8, 4, 0}, // flip vertical, rotate 270 (same as flip horizontal, rotate 90)
	},
}

/*
3x3 Transformations:

start:
	0123
	4567
	89x

rotate 90:
	8403
	9517
	x62

rotate 180:
	x983
	6547
	210

rotate 270:
	26x3
	1597
	048

flip horizontal:
	89x3
	4567
	012

flip horizontal, rotate 90:
	0483
	1597
	26x

flip horizontal, rotate 180:
	2103
	6547
	x98

flip horizontal, rotate 270:
	x623
	9517
	840

flip vertical:
	2103
	6547
	x98

flip vertical, rotate 90:
	x623
	9517
	840

flip vertical, rotate 180:
	89x3
	4567
	012

flip vertical, rotate 270:
	0483
	1597
	26x
*/

/*
2x2 Transformations:

start:
	012
	34

90:
	302
	41

180:
	432
	10

270:
	142
	03

flip horizontal:
	342
	01

flip horizontal (90):
	032
	14

flip horizontal (180): (same as flip vertical)
	102
	43

flip horizontal (270):
	412
	30

flip vertical:
	102
	43

flip vertical (90): (same as flip horizontal, 270)
	412
	30

flip vertical (180): (same as flip horizontal)
	342
	01

flip vertical (270): (same as flip horizontal, rotate 90)
	032
	14
*/
