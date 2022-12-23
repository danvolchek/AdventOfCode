package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type NodeType int

const (
	Start NodeType = iota
	End
	Neither
)

type Node struct {
	height int

	nodeType NodeType

	adjacent []*Node
}

func (n *Node) String() string {
	switch n.nodeType {
	case Start:
		return "S"
	case End:
		return "E"
	default:
		return string(rune('a' + n.height))
	}
}

func (n *Node) Adjacent() []*Node {
	return n.adjacent
}

func parse(char byte) *Node {
	var height int
	var nodeType NodeType

	switch char {
	case 'S':
		nodeType = Start
		height = 0
	case 'E':
		nodeType = End
		height = 26
	default:
		nodeType = Neither
		height = int(char - 'a')
	}

	return &Node{
		height:   height,
		nodeType: nodeType,
	}
}

func solve(grid [][]*Node) int {
	var start *Node

	for y, line := range grid {
		for x, node := range line {
			if node.nodeType == Start {
				start = node
			}

			adjacent := lib.Adjacent(false, y, x, grid)
			reachable := lib.Filter(adjacent, func(n *Node) bool {
				return node.height >= n.height-1
			})

			node.adjacent = reachable
		}
	}

	path, ok := lib.BFS(start, func(n *Node) bool {
		return n.nodeType == End
	})

	if !ok {
		panic("not found")
	}

	return len(path) - 1
}

func main() {
	solver := lib.Solver[[][]*Node, int]{
		ParseF: lib.ParseGrid(parse),
		SolveF: solve,
	}

	solver.Expect("Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi", 31)
	solver.Verify(330)
}
