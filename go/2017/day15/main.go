package day15

import "fmt"

func Example() {
	answer := generateMatches(40000000, NewGenerator(1, 16807, 65), NewGenerator(1, 48271, 8921))
	fmt.Println(answer)
}

func Part1() int {
	return generateMatches(40000000, NewGenerator(1, 16807, 277), NewGenerator(1, 48271, 349))
}

func Part2() int {
	return generateMatches(5000000, NewGenerator(4, 16807, 277), NewGenerator(8, 48271, 349))
}

func generateMatches(iterations int, a, b *Generator) (count int) {
	for x := 0; x < iterations; x++ {
		a.Next()
		b.Next()
		if a.Mask() == b.Mask() {
			count++
		}
	}
	return count
}

type Generator struct {
	divisor int
	factor  int
	Current int
}

func NewGenerator(divisor, factor, start int) *Generator {
	return &Generator{divisor: divisor, factor: factor, Current: start}
}

func (this *Generator) Next() {
	this.generate()
	for this.Current%this.divisor != 0 {
		this.generate()
	}
}

func (this *Generator) generate() {
	this.Current *= this.factor
	this.Current %= 2147483647
}

func (this *Generator) Mask() int {
	return this.Current &^ 0xFFFF0000
}
