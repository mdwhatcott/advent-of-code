package day07

import (
	"sort"
	"strings"

	"advent/lib/util"
)

type SortReference struct {
	forward  map[string][]string
	backward map[string][]string
	last     string
}

type TopologicalSort struct {
	SortReference
	stack *Stack
	order string
}

func NewTopologicalSort(input string) *TopologicalSort {
	this := new(TopologicalSort)
	this.forward, this.backward = parseRelations(input)
	this.last = findLast(this.forward)
	this.stack = NewStack()
	return this
}

func (this *TopologicalSort) Sort() string {
	this.stack.Push(reverse(findFirsts(this.forward))...)

	for this.stack.Size() > 0 {
		pop := this.stack.Pop()
		this.order += pop
		for _, next := range reverse(this.forward[pop]) {
			if this.isReady(next) {
				this.stack.Push(next)
			}
		}
	}
	return this.order
}

func (this *TopologicalSort) isReady(next string) bool {
	return !this.hasUnsatisfiedPrerequisite(next)
}

func (this *TopologicalSort) hasUnsatisfiedPrerequisite(next string) bool {
	for _, prerequisite := range this.backward[next] {
		if !strings.Contains(this.order, prerequisite) {
			return true
		}
	}
	return false
}

func findFirsts(steps map[string][]string) (firsts []string) {
	for key := range steps {
		found := false
		for _, values := range steps {
			if strings.Contains(strings.Join(values, ""), key) {
				found = true
			}
		}
		if !found {
			firsts = append(firsts, key)
		}
	}
	return firsts
}

func findLast(steps map[string][]string) (last string) {
	for _, values := range steps {
		for _, value := range values {
			_, found := steps[value]
			if !found {
				return value
			}
		}
	}
	panic("Couldn't find ending node.")
}

func parseRelations(input string) (forward, backward map[string][]string) {
	forward = make(map[string][]string)
	backward = make(map[string][]string)
	reader := strings.NewReader(input)
	scanner := util.NewScanner(reader)
	for scanner.Scan() {
		fields := scanner.Fields()
		before, after := fields[1], fields[7]
		forward[before] = append(forward[before], after)
		backward[after] = append(backward[after], before)
	}
	return forward, backward
}

func reverse(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}
