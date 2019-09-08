package main

import "sort"

func sum(values ...int) (sum int) {
	for _, value := range values {
		sum += value
	}
	return sum
}

func areEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for x := 0; x < len(a); x++ {
		if a[x] != b[x] {
			return false
		}
	}
	return true
}

type Sleigh struct {
	Passenger []int
	Left      []int
	Right     []int
}

func (this *Sleigh) IsComfortable() bool {
	return true
}

func (this *Sleigh) IsBalanced() bool {
	return areEqual(this.Passenger, this.Left) && areEqual(this.Left, this.Right)
}

func (this *Sleigh) QuantumEntanglement() int {
	return QuantumEntanglement(this.Passenger...)
}
