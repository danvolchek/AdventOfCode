package vm

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

type queue struct {
	head *node
	tail *node
}

func newQueue(values []int) *queue {
	q := &queue{}
	for _, v := range values {
		q.push(v)
	}

	return q
}

func (q *queue) push(v int) {
	newTail := &node{
		value: v,
	}

	if q.tail == nil {
		q.head = newTail
		q.tail = newTail
	} else {
		q.tail.next = newTail
		q.tail = newTail
	}

}

func (q *queue) pop() int {
	if q.head == nil {
		panic("empty queue!")
	}

	value := q.head.value

	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return value
}

func (q *queue) String() string {
	if q.head == nil {
		return "[]"
	}

	var ret strings.Builder

	ret.WriteRune('[')
	ret.WriteString(fmt.Sprintf("%d", q.head.value))
	v := q.head
	for v != nil {
		ret.WriteString(fmt.Sprintf(" %d", v.value))
		v = v.next
	}
	ret.WriteRune(']')

	return ret.String()
}

