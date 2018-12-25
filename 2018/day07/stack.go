package day07

import "sort"

type SortedStack struct {
	stack []string
}

func NewSortedStack() *SortedStack {
	return &SortedStack{}
}

func (this *SortedStack) Size() int {
	return len(this.stack)
}

func (this *SortedStack) Push(items ...string) {
	for _, item := range items {
		this.stack = append(this.stack, item)
	}
	reverse(this.stack)
}

func (this *SortedStack) Pop() string {
	item := this.stack[this.Size()-1]
	this.stack = this.stack[:this.Size()-1]
	return item
}

func (this *SortedStack) String() string {
	s := ""
	for _, x := range this.stack {
		s += x
	}
	return s
}

func reverse(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}
