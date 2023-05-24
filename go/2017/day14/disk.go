package day14

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
	"github.com/mdwhatcott/advent-of-code-knot"
)

type Disk struct {
	disk   [128][128]bool
	visual *bytes.Buffer
}

func makeDisk(key string) *Disk {
	disk := &Disk{visual: new(bytes.Buffer)}
	for row := 0; row < 128; row++ {
		digest := knot.HashString(key + "-" + strconv.Itoa(row))
		hash, _ := hex.DecodeString(digest)
		for s, sector := range hash {
			for d, digit := range encodeBinary(sector) {
				column := s*8 + d
				if digit == '1' {
					fmt.Fprint(disk.visual, "•")
				} else {
					fmt.Fprint(disk.visual, " ")
				}
				disk.disk[row][column] = digit == '1'
			}
		}
		fmt.Fprintln(disk.visual)
	}
	return disk
}

func (disk *Disk) CountUsedRegions() (regions int) {
	for row := 0; row < 128; row++ {
		for column := 0; column < 128; column++ {
			if !disk.disk[row][column] {
				continue
			}
			disk.resetRegion(row, column)
			regions++
		}
	}
	return regions
}

func (disk *Disk) CountUsedSectors() (used int) {
	for row := 0; row < 128; row++ {
		for column := 0; column < 128; column++ {
			if disk.disk[row][column] {
				used++
			}
		}
	}
	return used
}

func (disk *Disk) resetRegion(row, column int) {
	disk.disk[row][column] = false
	point := grid.NewPoint(float64(row), float64(column))
	for _, neighbor := range point.Neighbors4() {
		if x := int(neighbor.X()); x < 0 || x > 127 {
			continue
		} else if y := int(neighbor.Y()); y < 0 || y > 127 {
			continue
		} else if disk.disk[x][y] {
			disk.resetRegion(x, y)
		}
	}
}

func encodeBinary(value byte) string {
	return fmt.Sprintf("%08b", value)
}
