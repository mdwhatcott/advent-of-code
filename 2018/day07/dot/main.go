package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("digraph {")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		from, to := fields[1], fields[7]
		fmt.Printf("  %s -> %s\n", from, to)
	}
	fmt.Println("}")
}
