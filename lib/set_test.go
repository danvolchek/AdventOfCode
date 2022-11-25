package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"testing"
)

func TestSet_Size(t *testing.T) {
	var s lib.Set[int]

	if s.Size() != 0 {
		t.Fatal("Size of empty set is not zero")
	}

	s.Add(1)

	if s.Size() != 1 {
		t.Fatal("Size of set with 1 item is not 1")
	}

	s.Remove(1)

	if s.Size() != 0 {
		t.Fatal("Size of set with 0 items is not 0")
	}
}

func TestSet_Contains(t *testing.T) {
	var s lib.Set[int]

	if s.Contains(1) {
		t.Fatalf("Set should not contain 0")
	}

	s.Add(1)
	if !s.Contains(1) {
		t.Fatalf("Set should contain 1")
	}

	s.Remove(1)
	if s.Contains(1) {
		t.Fatalf("Set should not contain 1")
	}
}

func TestSet(t *testing.T) {
	var q lib.Set[int]

	for i := 0; i < 5; i++ {
		added := q.Add(i)
		if !added {
			t.Fatalf("expected to be able to add item")
		}
	}
	for i := 0; i < 5; i++ {
		added := q.Add(i)
		if added {
			t.Fatalf("expected not to be able to add item")
		}
	}

	actual := q.Items()
	slices.Sort(actual)

	expected := []int{0, 1, 2, 3, 4}
	if !slices.Equal(actual, expected) {
		t.Fatalf("got %+v, want %+v", actual, expected)
	}

	for i := 0; i < 5; i++ {
		removed := q.Remove(i)
		if !removed {
			t.Fatalf("expected to be able to remove item")
		}
	}

	for i := 0; i < 5; i++ {
		removed := q.Remove(i)
		if removed {
			t.Fatalf("expected not to be able to remove item")
		}
	}

}
