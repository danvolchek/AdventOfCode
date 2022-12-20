package lib

// Heap implements a binary heap: https://en.wikipedia.org/wiki/Binary_heap.
// Items are inserted with a priority and can later be efficiently retrieved.
// The zero value is an empty min heap.
// By default, items are retrieved in increasing priority order. Set Max to true to act as a max heap.
type Heap[T comparable] struct {
	items []heapItem[T]

	Max bool
}

type heapItem[T any] struct {
	item     T
	priority int
}

func (m *Heap[T]) Contains(item T) bool {
	for _, i := range m.items {
		if item == i.item {
			return true
		}
	}

	return false
}

// Empty returns whether the heap has no items in it.
func (m *Heap[T]) Empty() bool {
	return len(m.items) == 0
}

// Add adds a new item to the heap with the associated priority.
func (m *Heap[T]) Add(item T, priority int) {
	m.items = append(m.items, heapItem[T]{
		item:     item,
		priority: priority,
	})

	m.heapifyUp(len(m.items) - 1)
}

// Pop returns the item with the lowest priority (if !max, highest otherwise), removing it from the heap.
func (m *Heap[T]) Pop() T {
	if len(m.items) == 0 {
		panic("empty heap")
	}

	min := m.items[0].item

	m.items[0] = m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]

	m.heapifyDown(0)

	return min
}

func (m *Heap[T]) heapifyDown(index int) {
	if len(m.items) <= 1 {
		return
	}

	comp := func(a, b int) bool {
		if m.Max {
			return a > b
		}

		return a < b
	}

	for {
		left, right := 2*index+1, 2*index-1
		smallest := index

		if left >= 0 && left < len(m.items) && comp(m.items[left].priority, m.items[index].priority) {
			smallest = left
		}

		if right >= 0 && right < len(m.items) && comp(m.items[right].priority, m.items[smallest].priority) {
			smallest = right
		}

		if smallest == index {
			break
		}

		m.items[index], m.items[smallest] = m.items[smallest], m.items[index]

		smallest = index
	}
}

func (m *Heap[T]) heapifyUp(index int) {
	comp := func(a, b int) bool {
		if m.Max {
			return a <= b
		}

		return a >= b
	}

	for index > 0 {
		parent := (index - 1) / 2

		if comp(m.items[index].priority, m.items[parent].priority) {
			break
		}

		m.items[index], m.items[parent] = m.items[parent], m.items[index]

		index = parent
	}
}
