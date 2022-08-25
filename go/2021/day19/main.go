package day19

import (
	"bufio"
	"strings"

	"advent/lib/util"
)

func ParseScannerReports(reports string) (results [][]Point) {
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(reports) + "\n"))
	var beacon []Point
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			results = append(results, beacon)
			beacon = nil
			continue
		}
		if strings.HasPrefix(line, "---") {
			continue
		}
		fields := strings.Split(line, ",")
		x := util.ParseInt(fields[0])
		y := util.ParseInt(fields[1])
		z := util.ParseInt(fields[2])
		beacon = append(beacon, NewPoint(x, y, z))
	}
	return append(results, beacon)
}
