package starter

import (
	"container/list"

	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
)

func findShortestPaths(world set.Set[grid.Point[int]], start, end grid.Point[int]) (results [][]grid.Point[int]) {
	queue := list.New()
	queue.PushBack([]grid.Point[int]{start})
	visited := set.Of(start)
	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]grid.Point[int])
		current := path[len(path)-1]

		if grid.CityBlockDistance(current, end) == 1 && len(path) > 1 {
			results = append(results, path[1:])
			continue
		}
		for _, next := range current.Neighbors4() {
			if world.Contains(next) && !visited.Contains(next) {
				visited.Add(next)
				newPath := append([]grid.Point[int]{}, path...)
				newPath = append(newPath, next)
				queue.PushBack(newPath)
			}
		}
	}
	return results
}
