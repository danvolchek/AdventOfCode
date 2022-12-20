package lib

type AStarNode[T any] interface {
	comparable
	Adjacent() []T
}

func AStar[T AStarNode[T]](start T, target func(T) bool, h func(T) int) ([]T, bool) {
	hStart := h(start)

	var openSet Heap[T]
	openSet.Add(start, hStart)

	parent := make(map[T]T)

	gScore := map[T]int{
		start: 0,
	}
	fScore := map[T]int{
		start: hStart,
	}

	getScore := func(scoreMap map[T]int, n T) int {
		score, ok := scoreMap[n]
		if !ok {
			return 1<<63 - 5
		}

		return score
	}

	for !openSet.Empty() {
		curr := openSet.Pop()

		if target(curr) {
			return reconstructPath(start, curr, parent), true
		}

		for _, neighbor := range curr.Adjacent() {
			tentativeGScore := getScore(gScore, curr) + 1
			if neighborGScore, ok := gScore[neighbor]; !ok || tentativeGScore < neighborGScore {
				parent[neighbor] = curr
				gScore[neighbor] = tentativeGScore

				newFScore := tentativeGScore + h(neighbor)
				fScore[neighbor] = newFScore

				if !openSet.Contains(neighbor) {
					openSet.Add(neighbor, newFScore)
				}
			}
		}
	}

	return nil, false
}
