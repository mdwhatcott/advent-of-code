package main

import "fmt"

func main() {
	disk := scanDisk()
	fmt.Println("Part 1 - Viable pairs:", disk.countViableSectorPairs())
	fmt.Println()
	fmt.Println(disk)
}
