package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"strconv"
	"testing"
)

type testNode struct {
	id        int
	neighbors []*testNode
}

func (t *testNode) String() string {
	return strconv.Itoa(t.id)
}

func (t *testNode) Adjacent() []*testNode {
	return t.neighbors
}

func TestBFS(t *testing.T) {

	type testCase struct {
		name     string
		start    *testNode
		target   int
		expected []int
	}

	testCases := []testCase{
		{
			name:     "start is target",
			start:    &testNode{id: 0},
			target:   0,
			expected: []int{0},
		},
		{
			name: "one hop",
			start: &testNode{
				id: 0,
				neighbors: []*testNode{
					{
						id: 2,
					},
				},
			},
			target:   2,
			expected: []int{0, 2},
		},
		{
			name: "longer path",
			start: &testNode{
				id: 0,
				neighbors: []*testNode{
					{
						id: 1,
						neighbors: []*testNode{
							{
								id: 2,
							},
						},
					},
					{
						id: 2,
					},
				},
			},
			target:   2,
			expected: []int{0, 2},
		},
		{
			name: "cycle path",
			start: &testNode{
				id: 0,
				neighbors: []*testNode{
					{
						id: 1,
						neighbors: []*testNode{
							{
								id: 2,
								neighbors: []*testNode{
									{
										id: 0,
									},
									{
										id: 3,
									},
								},
							},
						},
					},
				},
			},
			target:   3,
			expected: []int{0, 1, 2, 3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := lib.BFS(tc.start, func(n *testNode) bool { return n.id == tc.target })
			if !ok {
				t.Fatalf("path not found")
			}

			ids := lib.Map(actual, func(n *testNode) int { return n.id })

			if !slices.Equal(tc.expected, ids) {
				t.Errorf("got %+v, want %+v", actual, tc.expected)
			}
		})
	}
}
