package main

import (
	"fmt"
)

type node struct {
	value character
	next  *node
}

type stack struct {
	top *node
}

func (s *stack) push(v character) {
	newTop := &node{
		value: v,
	}

	newTop.next = s.top
	s.top = newTop

}

func (s *stack) pop() character {
	if s.top == nil {
		panic("empty stack!")
	}

	value := s.top.value

	s.top = s.top.next

	return value
}

func (s *stack) empty() bool {
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
