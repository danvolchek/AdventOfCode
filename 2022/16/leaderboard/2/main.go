package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
	"strings"
)

type Valve struct {
	name      string
	flowRate  int
	neighbors []string
}

var reg = regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func parse(parts []string) Valve {
	rate := lib.Atoi(parts[1])
	connections := strings.Split(parts[2], ", ")

	return Valve{
		name:      parts[0],
		flowRate:  rate,
		neighbors: connections,
	}
}

type Node struct {
	name      string
	flowRate  int
	neighbors []*Node
}

func (n *Node) Adjacent() []*Node {
	return n.neighbors
}

func buildGraph(valves []Valve) *Node {
	nodeMap := make(map[string]*Node)
	allNodes := make(map[int]*Node)

	for i, valve := range valves {
		nodeMap[valve.name] = &Node{
			name:     valve.name,
			flowRate: valve.flowRate,
		}

		allNodes[i] = nodeMap[valve.name]
	}

	for _, valve := range valves {
		for _, neighbor := range valve.neighbors {
			nodeMap[valve.name].neighbors = append(nodeMap[valve.name].neighbors, nodeMap[neighbor])
		}
	}

	return nodeMap["AA"]
}

func solve(valves []Valve) int {
	//graph := buildGraph(valves)

	return 0
}

func main() {
	solver := lib.Solver[[]Valve, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II", 1707)
	solver.Incorrect(3389, 2275)
	solver.Solve()
}
