package main

import "fmt"

func main() {
	fmt.Println("Sequential Password:", NewSequentialPassword("uqwqemis").String())
	fmt.Println("Positional Password:", NewPositionalPassword("uqwqemis").String())
}
