package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"testing"
)

func TestHeap_Min(t *testing.T) {
	var m lib.Heap[string]

	m.Add("!", 5)
	m.Add("hello", 1)
	m.Add("world", 3)

	expected := []string{"hello", "world", "!"}

	var items []string
	for !m.Empty() {
		items = append(items, m.Pop())
	}

	if !slices.Equal(items, expected) {
		t.Fatalf("incorrect output: expected %v, got %v", expected, items)
	}
}

func TestHeap_Max(t *testing.T) {
	var m lib.Heap[string]
	m.Max = true

	m.Add("!", 1)
	m.Add("hello", 5)
	m.Add("world", 3)

	expected := []string{"hello", "world", "!"}

	var items []string
	for !m.Empty() {
		items = append(items, m.Pop())
	}

	if !slices.Equal(items, expected) {
		t.Fatalf("incorrect output: expected %v, got %v", expected, items)
	}
}
