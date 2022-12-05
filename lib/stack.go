package lib

// Stack is a simple first in, last out data structure implementation. It is not safe to use concurrently.
// The zero value is an empty stack.
type Stack[T any] struct {
	items []T
}

// Empty returns whether the stack has any items in it.
func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes an item from the stack.
func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("empty stack!")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek is like Pop but the item isn't removed.
func (s *Stack[T]) Peek() T {
	if s.Empty() {
		panic("can't pop empty stack")
	}

	return s.items[len(s.items)-1]
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Reverse reverses the order of the stack.
func (s *Stack[T]) Reverse() {
	for i := 0; i < len(s.items)/2; i++ {
		s.items[i], s.items[len(s.items)-i-1] = s.items[len(s.items)-i-1], s.items[i]
	}
}
