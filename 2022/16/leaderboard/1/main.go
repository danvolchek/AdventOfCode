package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
	"sort"
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
	name      int
	flowRate  int
	neighbors []*Node
}

func buildGraph(valves []Valve) *Node {
	nodeMap := make(map[string]*Node)

	for i, valve := range valves {
		nodeMap[valve.name] = &Node{
			name:     i,
			flowRate: valve.flowRate,
		}
	}

	for _, valve := range valves {
		for _, neighbor := range valve.neighbors {
			nodeMap[valve.name].neighbors = append(nodeMap[valve.name].neighbors, nodeMap[neighbor])
		}
	}

	return nodeMap["AA"]
}

func copyMap[T comparable, V any](m map[T]V) map[T]V {
	r := make(map[T]V, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

type args struct {
	you  int
	time int
	on   uint64
}

func hash(ints []int) uint64 {
	var result uint64

	for _, v := range ints {
		result += 1 << v
	}

	return result
}

var cache map[args]int
var cacheSeenAtAll map[args]bool
var total int
var maxTime int

func onToStr(a map[string]bool) string {
	v := lib.Keys(a)
	sort.Strings(v)
	return strings.Join(v, "")
}

func set2(index1 int, index2 int, on uint64) uint64 {
	return on + 1<<index1 + 1<<index2
}

func set(index int, on uint64) uint64 {
	return on + 1<<index
}

func isSet(index int, on uint64) bool {
	return (on>>index)&1 == 1
}

func size(on uint64) int {
	sum := 0
	for on > 0 {
		sum += int(on & 1)

		on = on >> 1
	}

	return sum
}

func exScore(valve *Node, time int, on uint64, onSize int) int {
	// nothing more we can do if there's no time left
	if time <= 0 {
		return 0
	}

	// if all the valves are on there's no point moving around more
	if onSize == total {
		return 0
	}

	aaargs := args{valve.name, time, on}

	if val, ok := cache[aaargs]; ok {
		return val
	}

	// if all the valves are open at a previous time, this path must be worse
	for i := time + 1; i < maxTime-onSize; i++ {
		if _, ok := cacheSeenAtAll[args{valve.name, i, on}]; ok {
			return 0
		}
	}

	cacheSeenAtAll[aaargs] = true

	var max int

	var scoreFromThisValve int
	if !isSet(valve.name, on) {
		// we can try turning the valve on
		scoreFromThisValve = valve.flowRate * (time - 1)

		turnedOn := set(valve.name, on)
		turnOnAndMove := scoreFromThisValve + exScore(valve, time-1, turnedOn, onSize+1)
		max = lib.Max(max, turnOnAndMove)
	}

	for _, neighbor := range valve.neighbors {
		justMove := exScore(neighbor, time-1, on, onSize)
		max = lib.Max(max, justMove)
	}

	cache[aaargs] = max
	return max
}

func solve(valves []Valve) int {
	graph := buildGraph(valves)

	var on uint64
	for i, valve := range valves {
		if valve.flowRate == 0 {
			on = set(i, on)
		}
	}

	total = len(valves)
	cache = make(map[args]int)
	cacheSeenAtAll = make(map[args]bool)
	maxTime = 30

	return exScore(graph, 30, on, size(on))
}

func main() {
	solver := lib.Solver[[]Valve, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II", 1651)
	solver.Verify(1724)
}
