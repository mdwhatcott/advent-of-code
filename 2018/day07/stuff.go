package day07

import (
	"strings"

	"advent/lib/util"
)

type SortReference struct {
	forward  map[string][]string
	backward map[string][]string

	workers      []*Worker
	taskSeconds  int
	totalSeconds int
}

type TopologicalSort struct {
	SortReference
	stack *SortedStack
	order string
}

func NewTopologicalSort(input string) *TopologicalSort {
	this := new(TopologicalSort)
	this.taskSeconds = 0
	this.workers = append(this.workers, NewWorker(0))
	this.forward, this.backward = parseRelations(input)
	this.stack = NewSortedStack()
	return this
}

func NewConcurrentTopologicalSort(input string, workers int, taskSeconds int) *TopologicalSort {
	this := NewTopologicalSort(input)
	for ; workers > 1; workers-- {
		this.workers = append(this.workers, NewWorker(taskSeconds))
	}
	return this
}

func (this *TopologicalSort) Sort() string {
	this.stack.Push(findFirsts(this.forward)...)

	for len(this.order) < len(this.forward)+1 {
		for _, worker := range this.workers {
			finished := worker.DoWork()
			this.order += finished

			for _, next := range this.forward[finished] {
				if this.isReady(next) {
					this.stack.Push(next)
				}
			}

			if worker.IsIdle() && this.stack.Size() > 0 {
				worker.Accept(this.stack.Pop())
			}
		}
		this.totalSeconds++
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

func (this *TopologicalSort) DurationSeconds() int {
	return this.totalSeconds
}

func (this *TopologicalSort) anyWorkersBusy() bool {
	for _, worker := range this.workers {
		if !worker.IsIdle() {
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
