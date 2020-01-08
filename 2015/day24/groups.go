package main

import (
	"sort"
)

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
	A []int
	B []int
	C []int
}

func (this *Sleigh) IsBalanced() bool {
	return areEqual(this.A, this.B) && areEqual(this.B, this.C)
}

func (this *Sleigh) IsComfortable() bool {
	return len(this.A) <= len(this.B) && len(this.A) <= len(this.C)
}

func (this *Sleigh) QuantumEntanglement() int {
	return QuantumEntanglement(this.A...)
}

func EnumerateSleighConfigurations(weights ...int) (all []Sleigh) {
	working := new(Set)
	for _, weight := range weights {
		working.Add(weight)
	}

	return all
}

func SumsOfTheThirdPart(weights ...int) (sums [][]int) {
	//third := sum(weights...) / 3
	//fmt.Println(third)
	sums = append(sums, []int{1, 2, 3, 4, 10})
	return sums
}

type Total struct {
	Target    int
	Working   *Set
	Available *Set
}

func (this *Total) Sum() int {
	return this.Working.Sum()
}

////////////////////////////////////////////////

type Set struct {
	contents []int
}

func (this *Set) Add(x int) {
	this.contents = append(this.contents, x)
}

func (this *Set) At(i int) int {
	if i < 0 || i >= len(this.contents) {
		return -1
	}
	return this.contents[i]
}

func (this *Set) Pop(i int) int {
	at := this.At(i)
	this.contents = append(this.contents[:i], this.contents[i+1:]...)
	return at
}

func (this *Set) Sum() int {
	return sum(this.contents...)
}
