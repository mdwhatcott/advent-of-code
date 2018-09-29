package day07

import (
	"strings"

	"advent/lib/util"
)

type Program struct {
	Name     string
	weight   int
	Children []*Program
}

func (this *Program) Weight() (sum int) {
	for _, child := range this.Children {
		sum += child.Weight()
	}
	return sum + this.weight
}

func (this *Program) FindUnbalance() (diff, value int) {
	weights := make(map[int]int)
	for _, child := range this.Children {
		weights[child.Weight()]++
	}

	var expected int
	for weight, count := range weights {
		if count > 1 {
			expected = weight
		}
	}

	for weight, count := range weights {
		if count == 1 {
			for _, child := range this.Children {
				if child.Weight() == weight {
					diff = expected - weight

					sub, w := child.FindUnbalance()
					if sub == 0 {
						return diff, child.weight
					} else {
						return sub, w
					}
				}
			}
		}
	}

	return 0, this.Weight()
}

func (this *Program) HasUniformChildren() bool {
	sums := make(map[int]struct{})
	for _, child := range this.Children {
		sums[child.Weight()] = struct{}{}
	}
	return len(sums) == 1
}

type Tower struct {
	listing map[string]*Program
	bottom  string
}

func NewTower() *Tower {
	return &Tower{listing: make(map[string]*Program)}
}

func (this *Tower) AddProgram(description string) {
	description = strings.Replace(description, "(", "", 1)
	description = strings.Replace(description, ")", "", 1)
	description = strings.Replace(description, ",", "", -1)
	fields := strings.Fields(description)
	if len(fields) < 2 {
		return
	}
	program := this.getProgram(fields[0])
	program.weight = util.ParseInt(strings.Trim(fields[1], "()"))

	for x := 3; x < len(fields); x++ {
		program.Children = append(program.Children, this.getProgram(fields[x]))
	}
}

func (this *Tower) getProgram(name string) *Program {
	program := this.listing[name]
	if program == nil {
		program = &Program{Name: name}
		this.listing[name] = program
	}
	return program
}

func (this *Tower) FindBottom() string {
	counts := make(map[string]int)
	for _, program := range this.listing {
		counts[program.Name]++
		for _, child := range program.Children {
			counts[child.Name]++
		}
	}
	for name, count := range counts {
		if count == 1 {
			this.bottom = name
			return name
		}
	}
	panic("NOT FOUND")
}
