package lib

// Node is a node used to do BFS.
type Node[I any, T any] interface {
	// Id is the id of this node.
	Id() I

	// Adjacent is the list of nodes reachable from this node.
	Adjacent() []T
}

type NodeConstraint[I, T any] interface {
	comparable
	Node[I, T]
}

// BFS returns the shortest path from start to target.
// The type of Start must be a Node.
func BFS[I comparable, T NodeConstraint[I, T]](start T, target I) []T {
	var q Queue[T]
	explored := map[I]bool{}
	parent := map[I]T{}

	q.Push(start)

	for !q.Empty() {
		v := q.Pop()
		if v.Id() == target {
			var result []T
			curr := v
			for {
				result = append(result, curr)

				var ok bool
				curr, ok = parent[curr.Id()]
				if !ok {
					break
				}

				if curr.Id() == start.Id() {
					result = append(result, curr)
					break
				}
			}
			rev := make([]T, len(result))
			for i, item := range result {
				rev[len(result)-i-1] = item
			}
			return rev
		}

		for _, neighbor := range v.Adjacent() {
			if !explored[neighbor.Id()] {
				explored[neighbor.Id()] = true
				parent[neighbor.Id()] = v
				q.Push(neighbor)
			}
		}
	}

	panic("target not found")
}
