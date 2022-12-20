package lib

// TODO: write tests.

// LinkedList represents a doubly-linked list of items. It provides O(1) removal/addition of nodes at any
// position in the list, but does not provide random access to list items.
// The zero value is a list of 1 item that has the zero value of T.
type LinkedList[T any] struct {
	Item T

	prev *LinkedList[T]
	next *LinkedList[T]
}

// Next returns the node after l.
func (l *LinkedList[T]) Next() *LinkedList[T] {
	return l.next
}

// Prev returns the node before l.
func (l *LinkedList[T]) Prev() *LinkedList[T] {
	return l.prev
}

// Remove removes l from the linked list, making its previous/next nodes
// point to each other and l point to nothing.
func (l *LinkedList[T]) Remove() {
	if l.prev != nil {
		l.prev.next = l.next
	}

	if l.next != nil {
		l.next.prev = l.prev
	}

	l.prev = nil
	l.next = nil
}

// Append sets m to be the node after l.
// If l is the end of the list, the result is l <-> m.
// If the current layout is l <-> n, the result is l <-> m <-> n.
func (l *LinkedList[T]) Append(m *LinkedList[T]) {
	n := l.next

	if n == nil {
		l.next = m
		m.prev = l
		return
	}

	n.prev = m
	l.next = m

	m.next = n
	m.prev = l
}

// Prepend sets m to be the node before l.
// If l is the start of the list, the result is m <-> l.
// If the current list is n <-> l, the result is n <-> m <-> l.
func (l *LinkedList[T]) Prepend(m *LinkedList[T]) {
	n := l.prev

	if n == nil {
		l.prev = m
		m.next = l
		return
	}

	n.next = m
	l.prev = m

	m.prev = n
	m.next = l
}
