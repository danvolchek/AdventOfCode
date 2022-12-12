package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
)

type Node struct {
	height int

	isStart, isEnd bool

	adjacent []*Node
}

func (n *Node) String() string {
	if n.isStart {
		return "S"
	}
	if n.isEnd {
		return "E"
	}

	return string(rune('a' + n.height))
}

func (n *Node) Adjacent() []*Node {
	return n.adjacent
}

func parse(char string) *Node {
	node := &Node{
		height:  int(char[0] - 'a'),
		isStart: char == "S",
		isEnd:   char == "E",
	}

	if node.isStart {
		node.height = 0
	}

	if node.isEnd {
		node.height = int('z' - 'a')
	}

	return node
}

func solve(grid [][]*Node) int {
	var start *Node

	for y, line := range grid {
		for x, node := range line {
			if node.isStart {
				start = node
			}

			adjacent := lib.Adjacent(false, y, x, grid)
			reachable := lib.Filter(adjacent, func(n *Node) bool {
				return node.height >= n.height-1
			})

			node.adjacent = reachable
		}
	}

	result, ok := lib.BFS(start, func(n *Node) bool {
		return n.isEnd
	})

	if !ok {
		panic("not found")
	}

	fmt.Println(result)
	return len(result) - 1
}

func main() {
	solver := lib.Solver[[][]*Node, int]{
		ParseF: lib.ParseGrid(parse),
		SolveF: solve,
	}

	solver.Expect("Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi", 31)
	solver.Verify(330)
}
