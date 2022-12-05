package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"testing"
)

func TestStack_Empty(t *testing.T) {
	var s lib.Stack[int]

	if !s.Empty() {
		t.Fatal("Zero value queue is not empty")
	}

	s.Push(0)
	if s.Empty() {
		t.Fatal("Queue with one item in it is empty")
	}

	s.Pop()
	if !s.Empty() {
		t.Fatal("Queue with no items in it is not empty")
	}
}

func TestStack_Reverse(t *testing.T) {
	var s lib.Stack[int]

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	s.Reverse()

	var actual []int
	for !s.Empty() {
		actual = append(actual, s.Pop())
	}

	expected := []int{0, 1, 2, 3, 4}
	if !slices.Equal(actual, expected) {
		t.Errorf("got %+v, want %+v", actual, expected)
	}
}

func TestStack(t *testing.T) {
	var s lib.Stack[int]

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	var actual []int
	for !s.Empty() {
		actual = append(actual, s.Pop())
	}

	expected := []int{4, 3, 2, 1, 0}
	if !slices.Equal(actual, expected) {
		t.Errorf("got %+v, want %+v", actual, expected)
	}
}
