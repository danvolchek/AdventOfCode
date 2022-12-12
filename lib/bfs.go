package lib

type Node[T any] interface {
	Adjacent() []T
}

type BFSConstraint[T any] interface {
	comparable
	Node[T]
}

// BFS runs breadth-first-search from start until target returns true.
// T must be a type that 1. has a method to return adjacent nodes (see Node[T]) and 2. is comparable.
func BFS[T BFSConstraint[T]](start T, target func(T) bool) ([]T, bool) {
	var q Queue[T]

	explored := map[T]bool{}
	parent := map[T]T{}

	q.Push(start)

	for !q.Empty() {
		v := q.Pop()

		if target(v) {
			return reconstructPath(start, v, parent), true
		}

		for _, neighbor := range v.Adjacent() {
			if !explored[neighbor] {
				explored[neighbor] = true
				parent[neighbor] = v
				q.Push(neighbor)
			}
		}
	}

	return nil, false
}

// reconstructPath reconstructs the path from start to end using a parent map.
func reconstructPath[T comparable](start, end T, parent map[T]T) []T {
	result := []T{end}

	if start == end {
		return result
	}

	curr := end
	for {
		curr = parent[curr]

		result = append(result, curr)

		if curr == start {
			break
		}
	}

	return Reverse(result)
}
