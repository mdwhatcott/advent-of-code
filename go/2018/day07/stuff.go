package day07

import (
	"sort"

	"advent/lib/set"
)

type Pool []*Worker

func (this Pool) AllIdle() bool {
	for _, worker := range this {
		if !worker.IsIdle() {
			return false
		}
	}
	return true
}

type Worker struct {
	count int
	work  string
}

func (this *Worker) IsFinishing() bool {
	return this.count == 1
}

func (this *Worker) Advance() {
	if this.count > 0 {
		this.count--
	}
}

func (this *Worker) IsIdle() bool {
	return this.count == 0
}

func (this *Worker) Enqueue(task string, delay int) {
	this.work = task
	this.count = int(task[0]-'A') + delay + 1
}

type TopologicalSort struct {
	pool         Pool
	tasks        set.Set[string]
	dependencies map[string]set.Set[string]
	done         set.Set[string]
	order        string
	seconds      int
	delay        int
}

func NewTopologicalSort(input string, workers, delay int) *TopologicalSort {
	tasks, dependencies := parseTasksWithDependencies(input)
	return &TopologicalSort{
		tasks:        tasks,
		dependencies: dependencies,
		done:         set.New[string](0),
		pool:         prepareWorkers(workers),
		delay:        delay,
	}
}

func (this *TopologicalSort) Sort() (order string, seconds int) {
	for {
		this.advanceWorkers()
		this.loadWorkers()
		if this.pool.AllIdle() {
			break
		}
		this.seconds++
	}

	return this.order, this.seconds
}

func (this *TopologicalSort) advanceWorkers() {
	for _, worker := range this.pool {
		if worker.IsFinishing() {
			this.done.Add(worker.work)
			this.order += worker.work
		}
		worker.Advance()
	}
}

func (this *TopologicalSort) loadWorkers() {
	for _, worker := range this.pool {
		ready := this.findReadyTasks()
		if len(ready) == 0 {
			break
		}
		sort.Strings(ready)
		for _, task := range ready {
			if worker.IsIdle() {
				this.tasks.Remove(task)
				worker.Enqueue(task, this.delay)
				break
			}
		}
	}
}

func (this *TopologicalSort) findReadyTasks() (ready []string) {
	for _, task := range this.tasks.Slice() {
		deps := this.dependencies[task]
		if deps == nil || deps.IsSubset(this.done) {
			ready = append(ready, task)
		}
	}
	return ready
}

func prepareWorkers(workers int) Pool {
	var pool Pool
	for w := 0; w < workers; w++ {
		pool = append(pool, new(Worker))
	}
	return pool
}
