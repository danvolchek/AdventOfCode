package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	var q lib.OrderedSet[int]

	for i := 0; i < 5; i++ {
		q.Add(i)
	}

	for i := 0; i < 100; i++ {
		actual := q.Items()

		expected := []int{0, 1, 2, 3, 4}
		if !slices.Equal(actual, expected) {
			t.Fatalf("got %+v, want %+v", actual, expected)
		}
	}
}
