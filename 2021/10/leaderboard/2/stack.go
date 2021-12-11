package main

import (
	"fmt"
)

type node struct {
	value byte
	next  *node
}

type stack struct {
	top *node
}

func newStack(values []byte) *stack {
	s := &stack{}
	for _, v := range values {
		s.push(v)
	}

	return s
}

func (s *stack) push(v byte) {
	newTop := &node{
		value: v,
	}

	newTop.next = s.top
	s.top = newTop

}

func (s *stack) pop() byte {
	if s.top == nil {
		panic("empty stack!")
	}

	value := s.top.value

	s.top = s.top.next

	return value
}

func (s *stack) Empty() bool {
	return s.top == nil
}

func (s *stack) String() string {
	if s.top == nil {
		return ""
	}

	ret := ""

	v := s.top
	for v != nil {
		ret = fmt.Sprintf("%c", v.value) + ret

		v = v.next
	}

	return ret
}
