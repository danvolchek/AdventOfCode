package lib

// Set is a simple set implementation that only holds unique values. It is not safe to use concurrently.
// The zero value is an empty queue.
type Set[T comparable] struct {
	items map[T]struct{}
}

// Contains returns whether the Set contains the item.
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Add adds an item to the Set. It returns whether the item was added.
func (s *Set[T]) Add(item T) bool {
	if s.items == nil {
		s.items = make(map[T]struct{})
	}

	_, ok := s.items[item]
	if ok {
		return false
	}

	s.items[item] = struct{}{}
	return true
}

// Remove removes an item from the Set. It returns whether the item was removed.
func (s *Set[T]) Remove(item T) bool {
	_, ok := s.items[item]
	if !ok {
		return false
	}

	delete(s.items, item)
	return true
}

// Items returns the items in the Set.
func (s *Set[T]) Items() []T {
	result := make([]T, len(s.items))

	i := 0
	for item := range s.items {
		result[i] = item
		i += 1
	}

	return result
}

// Size returns the size of the set.
func (s *Set[T]) Size() int {
	return len(s.items)
}
