package breadthfirstsearch

import (
	"errors"
	"grogos/shared"
)

var ERR_NOT_FOUND = errors.New("BFS didn't find what you are searching for :(")

// Ch 6. BFS
func BreadthFirstSearch[T comparable, K any](
	graph K,
	getInitialItems func(graph K) []T,
	getNeighbours func(graph K, item T) []T,
	checker func(item T) bool,
) (T, error) {
	var res T
	visitedItems := make(map[T]int)
	q := shared.NewQueue(getInitialItems(graph)...)
	for q.Size() > 0 {
		err, item := q.Deque()
		shared.Must(err)

		if _, isVisited := visitedItems[item]; isVisited {
			continue
		}

		visitedItems[item] = 1

		if checker(item) {
			return item, nil
		}

		q.EnqueMany(getNeighbours(graph, item)...)
	}
	return res, ERR_NOT_FOUND
}
