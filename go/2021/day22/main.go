package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent/lib/util"
)

type cuboid struct{ x, y, z int }

func inRange(c cuboid) bool {
	return c.x >= -50 && c.x <= 50 &&
		c.y >= -50 && c.y <= 50 &&
		c.z >= -50 && c.z <= 50
}

func parseLine(s string) (ons, offs []cuboid) {
	var all []cuboid
	m := rangesPattern.FindAllStringSubmatch(s, -1)[0]
	x1, x2, y1, y2, z1, z2 := N(m[1]), N(m[2]), N(m[3]), N(m[4]), N(m[5]), N(m[6])
	if x1 > 50 || y1 > 50 || z1 > 50 || x2 < -50 || y2 < -50 || z2 < -50 {
		return nil, nil
	}
	x1 = util.Max(-50, x1)
	y1 = util.Max(-50, y1)
	z1 = util.Max(-50, z1)
	x2 = util.Min(50, x2)
	y2 = util.Min(50, y2)
	z2 = util.Min(50, z2)
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for z := z1; z <= z2; z++ {
				all = append(all, cuboid{x, y, z})
			}
		}
	}
	if strings.HasPrefix(s, "on") {
		return all, nil
	}
	return nil, all
}

func N(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

const rangePattern = `(-?\d+)\.\.(-?\d+)`

var rangesPattern = regexp.MustCompile(fmt.Sprintf("x=%s,y=%s,z=%s", rangePattern, rangePattern, rangePattern))
