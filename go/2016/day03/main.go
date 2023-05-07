package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	part1 := 0
	part2 := 0
	rows := [][]string{}

	for _, line := range util.InputLines() {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		fields := strings.Fields(line)
		if isValidTriangle(fields[0], fields[1], fields[2]) {
			part1++
		}

		rows = append(rows, []string{fields[0], fields[1], fields[2]})
		if len(rows) < 3 {
			continue
		}
		a1, a2, a3 := rows[0][0], rows[1][0], rows[2][0]
		if isValidTriangle(a1, a2, a3) {
			part2++
		}
		b1, b2, b3 := rows[0][1], rows[1][1], rows[2][1]
		if isValidTriangle(b1, b2, b3) {
			part2++
		}
		c1, c2, c3 := rows[0][2], rows[1][2], rows[2][2]
		if isValidTriangle(c1, c2, c3) {
			part2++
		}
		rows = [][]string{}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func isValidTriangle(A, B, C string) bool {
	a, _ := strconv.Atoi(A)
	b, _ := strconv.Atoi(B)
	c, _ := strconv.Atoi(C)

	if a+b <= c {
		return false
	}
	if b+c <= a {
		return false
	}
	if c+a <= b {
		return false
	}
	return true
}
