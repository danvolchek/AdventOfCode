package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"testing"
)

func TestQueue_Empty(t *testing.T) {
	var q lib.Queue[int]

	if !q.Empty() {
		t.Fatal("Zero value queue is not empty")
	}

	q.Push(0)
	if q.Empty() {
		t.Fatal("Queue with one item in it is empty")
	}

	q.Pop()
	if !q.Empty() {
		t.Fatal("Queue with no items in it is not empty")
	}
}

func TestQueue(t *testing.T) {
	var q lib.Queue[int]

	for i := 0; i < 5; i++ {
		q.Push(i)
	}

	var actual []int
	for !q.Empty() {
		actual = append(actual, q.Pop())
	}

	expected := []int{0, 1, 2, 3, 4}
	if !slices.Equal(actual, expected) {
		t.Errorf("got %+v, want %+v", actual, expected)
	}
}
