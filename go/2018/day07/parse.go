package day07

import (
	"strings"

	"github.com/mdwhatcott/go-collections/set"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func parseTasksWithDependencies(input string) (tasks set.Set[string], dependencies map[string]set.Set[string]) {
	tasks = set.From[string]()
	dependencies = make(map[string]set.Set[string])

	scanner := util.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fields := scanner.Fields()
		a, b := fields[1], fields[7]
		tasks.Add(a)
		tasks.Add(b)
		deps := dependencies[b]
		if deps == nil {
			deps = set.From[string]()
			dependencies[b] = deps
		}
		deps.Add(a)
	}

	return tasks, dependencies
}
