package day07

import (
	"strings"

	"advent/lib/util"

	"github.com/deckarep/golang-set"
)

func parseTasksWithDependencies(input string) (tasks mapset.Set, dependencies map[string]mapset.Set) {
	tasks = mapset.NewSet()
	dependencies = make(map[string]mapset.Set)

	scanner := util.NewScanner(strings.NewReader(input))
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

	return tasks, dependencies
}
