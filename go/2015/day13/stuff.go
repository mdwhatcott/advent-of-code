package main

type Person struct {
	Name      string
	Relations map[*Person]int
}

func NewPerson(name string) *Person {
	return &Person{Name: name, Relations: map[*Person]int{}}
}

func (this *Person) String() string {
	return this.Name
}

func ComputeHappiestArrangement(people ...*Person) (max int) {
	combinations := &Combinations{}
	combinations.Permute(people, 0, len(people)-1)

	for _, combo := range combinations.combos {
		if score := computeHappinessOfCurrentArrangement(combo...); score > max {
			max = score
		}
	}
	return max
}

type Combinations struct {
	combos [][]*Person
}

// Permute is a translation of this helpful resource:
// http://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
func (this *Combinations) Permute(a []*Person, l, r int) {
	if l == r {
		this.combos = append(this.combos, a)
	} else {
		for i := l; i < r+1; i++ {
			c := clone(a)
			c[l], c[i] = c[i], c[l]
			this.Permute(c, l+1, r)
		}
	}
}

func clone(people []*Person) []*Person {
	var clones []*Person
	for _, p := range people {
		clones = append(clones, p)
	}
	return clones
}

func computeHappinessOfCurrentArrangement(people ...*Person) (sum int) {
	for x := 0; x < len(people)-1; x++ {
		sum += people[x].Relations[people[x+1]]
		sum += people[x+1].Relations[people[x]]
	}
	sum += people[len(people)-1].Relations[people[0]]
	sum += people[0].Relations[people[len(people)-1]]
	return sum
}
