package lib

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type OrderedSet[T constraints.Ordered] struct {
	Set[T]
}

// Items returns the items in the Set in sorted order.
func (s *OrderedSet[T]) Items() []T {
	result := s.Set.Items()

	slices.Sort(result)

	return result
}
