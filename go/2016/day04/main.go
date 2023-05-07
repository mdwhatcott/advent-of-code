package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	sum := 0
	for _, line := range util.InputLines() {
		room := ParseEncryptedRoom(line)
		if room.Valid {
			sum += room.SectorID
			if strings.Contains(room.Decrypt(), "northpole") {
				fmt.Println("The North Pole objects are in sector:", room.SectorID)
			}
		}
	}
	fmt.Println("Sum of all Sector IDs:", sum)
}
