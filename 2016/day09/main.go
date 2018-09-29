package main

import (
	"fmt"
	"strings"

	"advent/lib/util"
)

func main() {
	fmt.Println("V1:", GetUncompressedSize(strings.TrimSpace(util.InputString()), false))
	fmt.Println("V2:", GetUncompressedSize(strings.TrimSpace(util.InputString()), true))
}
