package day07

import (
	"advent/lib/util"

	"github.com/deckarep/golang-set"
)

func Do(scanner *util.Scanner, workers int) (order string, seconds int) {
	tasks := mapset.NewSet()
	dependencies := make(map[string]mapset.Set)

	for scanner.Scan() {
		fields := scanner.Fields()
		a, b := fields[1], fields[7]
		tasks.Add(a)
		tasks.Add(b)
		deps := dependencies[b]
		if deps == nil {
			deps = mapset.NewSet()
			dependencies[b] = deps
		}
		deps.Add(a)
	}

	done := mapset.NewSet()
	counts := make(util.Ints, workers)
	work := make(util.Strings, workers)
	for {
		for i, count := range counts {
			if count == 1 {
				done.Add(work[i])
				order += work[i]
			}
			counts[i] = util.Max(0, count-1)
		}

		for counts.Contains(0) {
			idle := counts.Index(0)
			var ready util.Strings
			for _, value := range tasks.ToSlice() {
				set := dependencies[value.(string)]
				if set == nil || set.IsSubset(done) {
					ready = append(ready, value.(string))
				}
			}
			if len(ready) == 0 {
				break
			}

			task := ready.Min()
			tasks.Remove(task)

			counts[idle] = int(task[0] - 'A' + 61)
			work[idle] = task
		}

		if counts.Sum() == 0 {
			break
		}
		seconds++
	}
	return order, seconds
}
