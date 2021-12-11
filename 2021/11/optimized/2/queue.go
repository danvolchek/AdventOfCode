package main

import (
	"fmt"
	"strings"
)

type node struct {
	value pos
	next  *node
}

type queue struct {
	head *node
	tail *node
}

func (q *queue) push(v pos) {
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

func (q *queue) pop() pos {
	if q.empty() {
		panic("empty queue!")
	}

	value := q.head.value

	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return value
}

func (q *queue) empty() bool {
	return q.head == nil
}

func (q *queue) String() string {
	if q.empty() {
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
