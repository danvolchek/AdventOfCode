package lib_test

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"strconv"
	"testing"
)

type aStarTestNode struct {
	id        int
	neighbors []*aStarTestNode
}

func (t *aStarTestNode) String() string {
	return strconv.Itoa(t.id)
}

func (t *aStarTestNode) Adjacent() []*aStarTestNode {
	return t.neighbors
}

func TestAStar_NoPriority(t *testing.T) {
	type testCase struct {
		name     string
		start    *aStarTestNode
		target   int
		expected []int
	}

	testCases := []testCase{
		{
			name:     "start is target",
			start:    &aStarTestNode{id: 0},
			target:   0,
			expected: []int{0},
		},
		{
			name: "one hop",
			start: &aStarTestNode{
				id: 0,
				neighbors: []*aStarTestNode{
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
			start: &aStarTestNode{
				id: 0,
				neighbors: []*aStarTestNode{
					{
						id: 1,
						neighbors: []*aStarTestNode{
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
			start: &aStarTestNode{
				id: 0,
				neighbors: []*aStarTestNode{
					{
						id: 1,
						neighbors: []*aStarTestNode{
							{
								id: 2,
								neighbors: []*aStarTestNode{
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
			actual, ok := lib.AStar(tc.start, func(n *aStarTestNode) bool {
				return n.id == tc.target
			}, func(n *aStarTestNode) int {
				return 0
			})
			if !ok {
				t.Fatalf("path not found")
			}

			ids := lib.Map(actual, func(n *aStarTestNode) int { return n.id })

			if !slices.Equal(tc.expected, ids) {
				t.Errorf("got %+v, want %+v", actual, tc.expected)
			}
		})
	}
}

func TestAStar(t *testing.T) {
	start := &aStarTestNode{
		id: 0,
	}

	totalNodes := 10
	target := totalNodes - 1
	var expected []int
	for i := 0; i < totalNodes; i++ {
		expected = append(expected, i)
	}

	curr := start
	for i := 1; i < totalNodes; i++ {
		newNode := &aStarTestNode{id: -i}
		curr.neighbors = append(curr.neighbors, newNode)
		newNode.neighbors = append(newNode.neighbors, curr)
		curr = newNode
	}

	curr = start
	for i := 1; i < totalNodes; i++ {
		newNode := &aStarTestNode{id: i}
		curr.neighbors = append(curr.neighbors, newNode)
		newNode.neighbors = append(newNode.neighbors, curr)
		curr = newNode
	}

	totalNodesSeen := 0
	actual, ok := lib.AStar(start, func(n *aStarTestNode) bool {
		totalNodesSeen += 1
		return n.id == target
	}, func(n *aStarTestNode) int {
		return target - n.id
	})
	if !ok {
		t.Fatalf("path not found")
	}

	ids := lib.Map(actual, func(n *aStarTestNode) int { return n.id })

	if !slices.Equal(expected, ids) {
		t.Errorf("got %+v, want %+v", actual, expected)
	}

	if totalNodesSeen != totalNodes {
		t.Errorf("heuristic didn't work, should have seen %v nodes but saw %v", totalNodes, totalNodesSeen)
	}
}
