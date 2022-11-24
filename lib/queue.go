package lib

// Queue is a simple first in, first out data structure implementation. It is not safe to use concurrently.
// The zero value is an empty queue.
type Queue[T any] struct {
	items []T
}

// Empty returns whether the queue has any items in it.
func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

// Push adds an item to the queue.
func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Pop removes an item from the queue.
func (q *Queue[T]) Pop() T {
	if len(q.items) == 0 {
		panic("empty queue!")
	}

	result := q.items[0]
	q.items = q.items[1:]

	return result
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return len(q.items)
}
